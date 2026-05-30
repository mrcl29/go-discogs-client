package discogs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// InventoryService provides access to Discogs' inventory batch operations.
// These endpoints allow exporting your entire inventory to CSV or uploading
// CSV files to add, change, or delete multiple listings at once.
//
// More info: https://www.discogs.com/developers/#page:inventory
type InventoryService interface {
	// RequestExport triggers a new export of your inventory as a CSV file.
	// Returns the URL where the export status can be tracked.
	// https://www.discogs.com/developers/#page:inventory,header:inventory-inventory-export-post
	RequestExport(ctx context.Context) (string, error)

	// GetRecentExports returns a paginated list of all recent inventory exports.
	// https://www.discogs.com/developers/#page:inventory,header:inventory-inventory-export
	GetRecentExports(ctx context.Context, opts *PageOptions) (*InventoryExportsResponse, error)

	// GetExport retrieves details about the status and metadata of a specific inventory export.
	// https://www.discogs.com/developers/#page:inventory,header:inventory-inventory-export-get
	GetExport(ctx context.Context, exportID int) (*InventoryExport, error)

	// DownloadExport downloads the CSV content of a completed inventory export.
	// The caller is responsible for closing the returned ReadCloser.
	// https://www.discogs.com/developers/#page:inventory,header:inventory-inventory-export-download
	DownloadExport(ctx context.Context, exportID int) (io.ReadCloser, error)

	// AddInventory uploads a CSV file to add new listings to your inventory.
	// https://www.discogs.com/developers/#page:inventory,header:inventory-inventory-upload-post
	AddInventory(ctx context.Context, csvData io.Reader) (*InventoryUpload, error)

	// ChangeInventory uploads a CSV file to update existing listings in your inventory.
	// https://www.discogs.com/developers/#page:inventory,header:inventory-inventory-upload-post-edit
	ChangeInventory(ctx context.Context, csvData io.Reader) (*InventoryUpload, error)

	// DeleteInventory uploads a CSV file to remove listings from your inventory.
	// https://www.discogs.com/developers/#page:inventory,header:inventory-inventory-upload-post-delete
	DeleteInventory(ctx context.Context, csvData io.Reader) (*InventoryUpload, error)

	// GetRecentUploads returns a paginated list of all recent inventory upload operations.
	// https://www.discogs.com/developers/#page:inventory,header:inventory-inventory-upload
	GetRecentUploads(ctx context.Context, opts *PageOptions) (*InventoryUploadsResponse, error)

	// GetUpload retrieves details about the status and results of a specific inventory upload.
	// https://www.discogs.com/developers/#page:inventory,header:inventory-inventory-upload-get
	GetUpload(ctx context.Context, uploadID int) (*InventoryUpload, error)
}

// InventoryExport represents the metadata and status of an inventory export task.
type InventoryExport struct {
	ID          int    `json:"id"`
	Status      string `json:"status"`
	URL         string `json:"url"`
	DownloadURL string `json:"download_url"`
	CreatedTS   string `json:"created_ts"`
	FinishedTS  string `json:"finished_ts"`
	Filename    string `json:"filename"`
}

// InventoryExportsResponse represents a paginated list of inventory exports.
type InventoryExportsResponse struct {
	Pagination Pagination        `json:"pagination"`
	Items      []InventoryExport `json:"items"`
}

// InventoryUpload represents the metadata and status of an inventory upload task.
type InventoryUpload struct {
	ID         int    `json:"id"`
	Status     string `json:"status"`
	CreatedTS  string `json:"created_ts"`
	FinishedTS string `json:"finished_ts"`
	Filename   string `json:"filename"`
	Results    string `json:"results"`
	Type       string `json:"type"`
}

// InventoryUploadsResponse represents a paginated list of inventory uploads.
type InventoryUploadsResponse struct {
	Pagination Pagination        `json:"pagination"`
	Items      []InventoryUpload `json:"items"`
}

type inventoryService struct {
	client *Client
}

func (s *inventoryService) RequestExport(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s/inventory/export", s.client.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return "", err
	}

	// The API returns the location of the new export in the Location header
	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("discogs API error: %d %s", resp.StatusCode, resp.Status)
	}

	return resp.Header.Get("Location"), nil
}

func (s *inventoryService) GetRecentExports(ctx context.Context, opts *PageOptions) (*InventoryExportsResponse, error) {
	baseURL := fmt.Sprintf("%s/inventory/export", s.client.baseURL)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp InventoryExportsResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *inventoryService) GetExport(ctx context.Context, exportID int) (*InventoryExport, error) {
	url := fmt.Sprintf("%s/inventory/export/%d", s.client.baseURL, exportID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var export InventoryExport
	if err := s.client.do(ctx, req, &export); err != nil {
		return nil, err
	}
	return &export, nil
}

func (s *inventoryService) DownloadExport(ctx context.Context, exportID int) (io.ReadCloser, error) {
	url := fmt.Sprintf("%s/inventory/export/%d/download", s.client.baseURL, exportID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("discogs API error: %d %s", resp.StatusCode, resp.Status)
	}

	return resp.Body, nil
}

func (s *inventoryService) AddInventory(ctx context.Context, csvData io.Reader) (*InventoryUpload, error) {
	return s.upload(ctx, "add", csvData)
}

func (s *inventoryService) ChangeInventory(ctx context.Context, csvData io.Reader) (*InventoryUpload, error) {
	return s.upload(ctx, "change", csvData)
}

func (s *inventoryService) DeleteInventory(ctx context.Context, csvData io.Reader) (*InventoryUpload, error) {
	return s.upload(ctx, "delete", csvData)
}

func (s *inventoryService) upload(ctx context.Context, action string, csvData io.Reader) (*InventoryUpload, error) {
	url := fmt.Sprintf("%s/inventory/upload/%s", s.client.baseURL, action)

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	part, err := w.CreateFormFile("upload", "inventory.csv")
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, csvData); err != nil {
		return nil, err
	}
	_ = w.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	var upload InventoryUpload
	if err := s.client.do(ctx, req, &upload); err != nil {
		return nil, err
	}
	return &upload, nil
}

func (s *inventoryService) GetRecentUploads(ctx context.Context, opts *PageOptions) (*InventoryUploadsResponse, error) {
	baseURL := fmt.Sprintf("%s/inventory/upload", s.client.baseURL)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp InventoryUploadsResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *inventoryService) GetUpload(ctx context.Context, uploadID int) (*InventoryUpload, error) {
	url := fmt.Sprintf("%s/inventory/upload/%d", s.client.baseURL, uploadID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var upload InventoryUpload
	if err := s.client.do(ctx, req, &upload); err != nil {
		return nil, err
	}
	return &upload, nil
}
