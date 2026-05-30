package discogs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CollectionService provides access to user collection management.
// It allows organizing releases into folders, rating items, and valuing your collection.
//
// More info: https://www.discogs.com/developers/#page:user-collection
type CollectionService interface {
	// ListFolders retrieves a list of folders in a user's collection.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-collection-folders
	ListFolders(ctx context.Context, username string) (*CollectionFoldersResponse, error)

	// CreateFolder creates a new folder in a user's collection. Requires OAuth.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-collection-folders-post
	CreateFolder(ctx context.Context, username string, name string) (*Folder, error)

	// GetFolder retrieves metadata about a specific folder in a user's collection.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-collection-folder
	GetFolder(ctx context.Context, username string, folderID int) (*Folder, error)

	// UpdateFolder edits a folder's name. Folders 0 ("All") and 1 ("Uncategorized") cannot be renamed. Requires OAuth.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-collection-folder-post
	UpdateFolder(ctx context.Context, username string, folderID int, name string) (*Folder, error)

	// DeleteFolder deletes an empty folder from a user's collection. Requires OAuth.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-collection-folder-delete
	DeleteFolder(ctx context.Context, username string, folderID int) error

	// GetCollectionItemsByRelease views the user's collection folders which contain a specified release.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-collection-items-by-release
	GetCollectionItemsByRelease(ctx context.Context, username string, releaseID int) (*CollectionItemsResponse, error)

	// GetCollectionItemsByFolder returns a paginated list of items in a specific collection folder.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-collection-items-by-folder
	GetCollectionItemsByFolder(ctx context.Context, username string, folderID int, opts *PageOptions) (*CollectionItemsResponse, error)

	// AddToCollectionFolder adds a release to a folder in a user's collection. Requires OAuth.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-add-to-collection-folder
	AddToCollectionFolder(ctx context.Context, username string, folderID int, releaseID int) (*CollectionItem, error)

	// UpdateInstance changes the rating of a release instance or moves it to another folder. Requires OAuth.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-change-rating-of-release
	UpdateInstance(ctx context.Context, username string, folderID int, releaseID int, instanceID int, data map[string]interface{}) error

	// DeleteInstance removes an instance of a release from a user's collection folder. Requires OAuth.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-delete-instance-from-folder
	DeleteInstance(ctx context.Context, username string, folderID int, releaseID int, instanceID int) error

	// ListCustomFields retrieves a list of user-defined collection notes fields.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-list-custom-fields
	ListCustomFields(ctx context.Context, username string) ([]NoteField, error)

	// EditInstanceField changes the value of a custom notes field on a particular release instance. Requires OAuth.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-edit-fields-instance
	EditInstanceField(ctx context.Context, username string, folderID int, releaseID int, instanceID int, fieldID int, value string) error

	// GetCollectionValue returns the minimum, median, and maximum value of a user's collection. Requires OAuth.
	// https://www.discogs.com/developers/#page:user-collection,header:user-collection-collection-value
	GetCollectionValue(ctx context.Context, username string) (*CollectionValue, error)
}

// NoteField represents a user-defined custom notes field.
type NoteField struct {
	// ID is the unique identifier for the custom field.
	ID int `json:"id"`
	// Name is the name of the custom field (e.g., "Shelf Location").
	Name string `json:"name"`
	// Type is the field input type: e.g., "dropdown" or "textarea".
	Type string `json:"type"`
	// Public indicates if the field is visible to other users.
	Public bool `json:"public"`
	// Position is the display order of the field.
	Position int `json:"position"`
	// Options is a list of valid values for "dropdown" type fields.
	Options []string `json:"options,omitempty"`
	// Lines is the number of lines to display for "textarea" type fields.
	Lines int `json:"lines,omitempty"`
}

// CollectionValue represents the estimated monetary value of a collection.
type CollectionValue struct {
	// Minimum is the estimated minimum value of the collection.
	Minimum string `json:"minimum"`
	// Median is the estimated median value of the collection.
	Median string `json:"median"`
	// Maximum is the estimated maximum value of the collection.
	Maximum string `json:"maximum"`
}

type collectionService struct {
	client *Client
}

func (s *collectionService) ListFolders(ctx context.Context, username string) (*CollectionFoldersResponse, error) {
	url := fmt.Sprintf("%s/users/%s/collection/folders", s.client.baseURL, username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var resp CollectionFoldersResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *collectionService) CreateFolder(ctx context.Context, username string, name string) (*Folder, error) {
	url := fmt.Sprintf("%s/users/%s/collection/folders", s.client.baseURL, username)
	body, _ := json.Marshal(map[string]string{"name": name})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var folder Folder
	if err := s.client.do(ctx, req, &folder); err != nil {
		return nil, err
	}
	return &folder, nil
}

func (s *collectionService) GetFolder(ctx context.Context, username string, folderID int) (*Folder, error) {
	url := fmt.Sprintf("%s/users/%s/collection/folders/%d", s.client.baseURL, username, folderID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var folder Folder
	if err := s.client.do(ctx, req, &folder); err != nil {
		return nil, err
	}
	return &folder, nil
}

func (s *collectionService) UpdateFolder(ctx context.Context, username string, folderID int, name string) (*Folder, error) {
	url := fmt.Sprintf("%s/users/%s/collection/folders/%d", s.client.baseURL, username, folderID)
	body, _ := json.Marshal(map[string]string{"name": name})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var folder Folder
	if err := s.client.do(ctx, req, &folder); err != nil {
		return nil, err
	}
	return &folder, nil
}

func (s *collectionService) DeleteFolder(ctx context.Context, username string, folderID int) error {
	url := fmt.Sprintf("%s/users/%s/collection/folders/%d", s.client.baseURL, username, folderID)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	return s.client.do(ctx, req, nil)
}

func (s *collectionService) GetCollectionItemsByRelease(ctx context.Context, username string, releaseID int) (*CollectionItemsResponse, error) {
	url := fmt.Sprintf("%s/users/%s/collection/releases/%d", s.client.baseURL, username, releaseID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var resp CollectionItemsResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *collectionService) GetCollectionItemsByFolder(ctx context.Context, username string, folderID int, opts *PageOptions) (*CollectionItemsResponse, error) {
	baseURL := fmt.Sprintf("%s/users/%s/collection/folders/%d/releases", s.client.baseURL, username, folderID)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp CollectionItemsResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *collectionService) AddToCollectionFolder(ctx context.Context, username string, folderID int, releaseID int) (*CollectionItem, error) {
	url := fmt.Sprintf("%s/users/%s/collection/folders/%d/releases/%d", s.client.baseURL, username, folderID, releaseID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	var item CollectionItem
	if err := s.client.do(ctx, req, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *collectionService) UpdateInstance(ctx context.Context, username string, folderID int, releaseID int, instanceID int, data map[string]interface{}) error {
	url := fmt.Sprintf("%s/users/%s/collection/folders/%d/releases/%d/instances/%d", s.client.baseURL, username, folderID, releaseID, instanceID)
	body, _ := json.Marshal(data)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return s.client.do(ctx, req, nil)
}

func (s *collectionService) DeleteInstance(ctx context.Context, username string, folderID int, releaseID int, instanceID int) error {
	url := fmt.Sprintf("%s/users/%s/collection/folders/%d/releases/%d/instances/%d", s.client.baseURL, username, folderID, releaseID, instanceID)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	return s.client.do(ctx, req, nil)
}

func (s *collectionService) ListCustomFields(ctx context.Context, username string) ([]NoteField, error) {
	url := fmt.Sprintf("%s/users/%s/collection/fields", s.client.baseURL, username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Fields []NoteField `json:"fields"`
	}
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Fields, nil
}

func (s *collectionService) EditInstanceField(ctx context.Context, username string, folderID int, releaseID int, instanceID int, fieldID int, value string) error {
	url := fmt.Sprintf("%s/users/%s/collection/folders/%d/releases/%d/instances/%d/fields/%d", s.client.baseURL, username, folderID, releaseID, instanceID, fieldID)
	body, _ := json.Marshal(map[string]string{"value": value})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return s.client.do(ctx, req, nil)
}

func (s *collectionService) GetCollectionValue(ctx context.Context, username string) (*CollectionValue, error) {
	url := fmt.Sprintf("%s/users/%s/collection/value", s.client.baseURL, username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var val CollectionValue
	if err := s.client.do(ctx, req, &val); err != nil {
		return nil, err
	}
	return &val, nil
}
