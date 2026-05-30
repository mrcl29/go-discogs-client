package discogs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// MarketplaceService provides access to Discogs' marketplace features, including
// inventory management, order processing, and price metadata.
//
// More info: https://www.discogs.com/developers/#page:marketplace
type MarketplaceService interface {
	// GetInventory returns a paginated list of listings in a user's inventory.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-inventory
	GetInventory(ctx context.Context, username string, opts *PageOptions, status string) (*InventoryResponse, error)

	// GetListing views the data associated with a single marketplace listing.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-listing
	GetListing(ctx context.Context, listingID int, currAbbr string) (*Listing, error)

	// CreateListing creates a new marketplace listing. Requires OAuth.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-listing-post
	CreateListing(ctx context.Context, data map[string]interface{}) (*ListingIDResponse, error)

	// UpdateListing edits the data associated with an existing listing. Requires OAuth.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-listing-post-edit
	UpdateListing(ctx context.Context, listingID int, data map[string]interface{}) error

	// DeleteListing permanently removes a listing from the marketplace. Requires OAuth.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-listing-delete
	DeleteListing(ctx context.Context, listingID int) error

	// GetOrder views the data associated with a marketplace order. Requires OAuth.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-order
	GetOrder(ctx context.Context, orderID string) (*Order, error)

	// UpdateOrder edits the status or data of an existing order. Requires OAuth.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-order-post
	UpdateOrder(ctx context.Context, orderID string, data map[string]interface{}) (*Order, error)

	// ListOrders returns a paginated list of the authenticated user's orders. Requires OAuth.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-list-orders
	ListOrders(ctx context.Context, opts *PageOptions, status string) (*OrdersResponse, error)

	// ListOrderMessages returns a paginated list of messages for a specific order. Requires OAuth.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-order-messages
	ListOrderMessages(ctx context.Context, orderID string, opts *PageOptions) (*OrderMessagesResponse, error)

	// CreateOrderMessage adds a new message to an order's communication log. Requires OAuth.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-order-messages-post
	CreateOrderMessage(ctx context.Context, orderID string, message string, status string) (*OrderMessage, error)

	// GetFee calculates the Discogs fee for selling an item at a specific price.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-fee
	GetFee(ctx context.Context, price float64, currency string) (*FeeResponse, error)

	// GetPriceSuggestions retrieves recommended prices for a release based on its condition.
	// Requires that the user has selling permissions and a verified seller account.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-price-suggestions
	GetPriceSuggestions(ctx context.Context, releaseID int) (map[string]Price, error)

	// GetReleaseMarketplaceStats retrieves high-level marketplace statistics for a release.
	// https://www.discogs.com/developers/#page:marketplace,header:marketplace-release-stats
	GetReleaseMarketplaceStats(ctx context.Context, releaseID int, currAbbr string) (*MarketplaceStats, error)
}

// ListingIDResponse contains the ID of a newly created listing.
type ListingIDResponse struct {
	// ListingID is the unique identifier for the new listing.
	ListingID int `json:"listing_id"`
	// ResourceURL is the API endpoint for the new listing.
	ResourceURL string `json:"resource_url"`
}

// OrderMessagesResponse represents a paginated list of order messages.
type OrderMessagesResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Messages is the list of messages in the order log.
	Messages []OrderMessage `json:"messages"`
}

// OrderMessage represents a single message or status event in an order's log.
type OrderMessage struct {
	// ID is the unique identifier for the message.
	ID string `json:"id"`
	// Timestamp is the date and time the message was sent.
	Timestamp string `json:"timestamp"`
	// Message is the text content of the message.
	Message string `json:"message"`
	// Type is the message type (e.g., "message", "status").
	Type string `json:"type"`
	// Subject is the message subject line.
	Subject string `json:"subject"`
	// From is a reference to the user who sent the message.
	From UserRef `json:"from"`
	// Order is a reference to the associated order.
	Order OrderRef `json:"order"`
}

// OrderRef provides a simple reference to a marketplace order.
type OrderRef struct {
	// ID is the unique identifier for the order.
	ID string `json:"id"`
	// ResourceURL is the API endpoint for the order.
	ResourceURL string `json:"resource_url"`
}

// FeeResponse contains the results of a marketplace fee calculation.
type FeeResponse struct {
	// Value is the numeric fee amount.
	Value float64 `json:"value"`
	// Currency is the currency code for the fee.
	Currency string `json:"currency"`
}

// MarketplaceStats contains marketplace availability and pricing statistics for a release.
type MarketplaceStats struct {
	// LowestPrice is the current minimum price for the release.
	LowestPrice Price `json:"lowest_price"`
	// NumForSale is the total number of copies available.
	NumForSale int `json:"num_for_sale"`
	// BlockedFromSale indicates if the release is restricted from marketplace sales.
	BlockedFromSale bool `json:"blocked_from_sale"`
}

type marketplaceService struct {
	client *Client
}

func (s *marketplaceService) GetInventory(ctx context.Context, username string, opts *PageOptions, status string) (*InventoryResponse, error) {
	baseURL := fmt.Sprintf("%s/users/%s/inventory", s.client.baseURL, username)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	if status != "" {
		u, _ := http.NewRequest(http.MethodGet, finalURL, nil)
		q := u.URL.Query()
		q.Set("status", status)
		u.URL.RawQuery = q.Encode()
		finalURL = u.URL.String()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp InventoryResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *marketplaceService) GetListing(ctx context.Context, listingID int, currAbbr string) (*Listing, error) {
	url := fmt.Sprintf("%s/marketplace/listings/%d", s.client.baseURL, listingID)
	if currAbbr != "" {
		url += "?curr_abbr=" + currAbbr
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var listing Listing
	if err := s.client.do(ctx, req, &listing); err != nil {
		return nil, err
	}
	return &listing, nil
}

func (s *marketplaceService) CreateListing(ctx context.Context, data map[string]interface{}) (*ListingIDResponse, error) {
	url := fmt.Sprintf("%s/marketplace/listings", s.client.baseURL)
	body, _ := json.Marshal(data)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var resp ListingIDResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *marketplaceService) UpdateListing(ctx context.Context, listingID int, data map[string]interface{}) error {
	url := fmt.Sprintf("%s/marketplace/listings/%d", s.client.baseURL, listingID)
	body, _ := json.Marshal(data)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return s.client.do(ctx, req, nil)
}

func (s *marketplaceService) DeleteListing(ctx context.Context, listingID int) error {
	url := fmt.Sprintf("%s/marketplace/listings/%d", s.client.baseURL, listingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	return s.client.do(ctx, req, nil)
}

func (s *marketplaceService) GetOrder(ctx context.Context, orderID string) (*Order, error) {
	url := fmt.Sprintf("%s/marketplace/orders/%s", s.client.baseURL, orderID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := s.client.do(ctx, req, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *marketplaceService) UpdateOrder(ctx context.Context, orderID string, data map[string]interface{}) (*Order, error) {
	url := fmt.Sprintf("%s/marketplace/orders/%s", s.client.baseURL, orderID)
	body, _ := json.Marshal(data)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var order Order
	if err := s.client.do(ctx, req, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *marketplaceService) ListOrders(ctx context.Context, opts *PageOptions, status string) (*OrdersResponse, error) {
	baseURL := fmt.Sprintf("%s/marketplace/orders", s.client.baseURL)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	if status != "" {
		u, _ := http.NewRequest(http.MethodGet, finalURL, nil)
		q := u.URL.Query()
		q.Set("status", status)
		u.URL.RawQuery = q.Encode()
		finalURL = u.URL.String()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp OrdersResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *marketplaceService) ListOrderMessages(ctx context.Context, orderID string, opts *PageOptions) (*OrderMessagesResponse, error) {
	baseURL := fmt.Sprintf("%s/marketplace/orders/%s/messages", s.client.baseURL, orderID)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp OrderMessagesResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *marketplaceService) CreateOrderMessage(ctx context.Context, orderID string, message string, status string) (*OrderMessage, error) {
	url := fmt.Sprintf("%s/marketplace/orders/%s/messages", s.client.baseURL, orderID)
	data := map[string]string{}
	if message != "" {
		data["message"] = message
	}
	if status != "" {
		data["status"] = status
	}
	body, _ := json.Marshal(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var msg OrderMessage
	if err := s.client.do(ctx, req, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

func (s *marketplaceService) GetFee(ctx context.Context, price float64, currency string) (*FeeResponse, error) {
	url := fmt.Sprintf("%s/marketplace/fee/%.2f", s.client.baseURL, price)
	if currency != "" {
		url += "/" + currency
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var resp FeeResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *marketplaceService) GetPriceSuggestions(ctx context.Context, releaseID int) (map[string]Price, error) {
	url := fmt.Sprintf("%s/marketplace/price_suggestions/%d", s.client.baseURL, releaseID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var resp map[string]Price
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *marketplaceService) GetReleaseMarketplaceStats(ctx context.Context, releaseID int, currAbbr string) (*MarketplaceStats, error) {
	url := fmt.Sprintf("%s/marketplace/stats/%d", s.client.baseURL, releaseID)
	if currAbbr != "" {
		url += "?curr_abbr=" + currAbbr
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var stats MarketplaceStats
	if err := s.client.do(ctx, req, &stats); err != nil {
		return nil, err
	}
	return &stats, nil
}
