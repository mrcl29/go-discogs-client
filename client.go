package discogs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

var (
	// ErrUnauthorized is returned when the API request lacks valid authentication
	// or the credentials provided are invalid.
	ErrUnauthorized = errors.New("discogs: unauthorized access - check your credentials")

	// ErrNotFound is returned when the requested resource (Release, Artist, etc.)
	// does not exist in the Discogs database.
	ErrNotFound = errors.New("discogs: resource not found")

	// ErrRateLimited is returned when the client exceeds Discogs rate limits.
	// The client handles this locally, so this error usually indicates a logic issue
	// or that the limiter was bypassed.
	ErrRateLimited = errors.New("discogs: rate limit exceeded - slow down")
)

const (
	// DefaultBaseURL is the standard API endpoint for Discogs v2.0.
	DefaultBaseURL = "https://api.discogs.com"

	// DefaultUserAgent is the default identifier sent in the User-Agent header.
	// Using a custom User-Agent via WithUserAgent is highly recommended.
	DefaultUserAgent = "GoDiscogsClient/1.0"
)

// Client is the main orchestrator for the Discogs API SDK.
// It manages the underlying HTTP client, authentication strategies, and rate limiting.
// All API interactions are performed through the service fields attached to this struct.
//
// Use NewClient() to create a thread-safe instance.
type Client struct {
	httpClient  *http.Client
	baseURL     string
	userAgent   string
	auth        Authenticator
	rateLimiter *rate.Limiter

	// Database provides methods to browse the Discogs music database.
	// See DatabaseService for more details.
	Database DatabaseService

	// Marketplace provides methods to manage listings, orders, and pricing.
	// See MarketplaceService for more details.
	Marketplace MarketplaceService

	// Inventory provides methods for batch CSV inventory operations.
	// See InventoryService for more details.
	Inventory InventoryService

	// User provides methods to access user profiles and account identity.
	// See UserService for more details.
	User UserService

	// Collection provides methods to manage and value a user's personal record collection.
	// See CollectionService for more details.
	Collection CollectionService

	// Wantlist provides methods to manage a user's release wantlist.
	// See WantlistService for more details.
	Wantlist WantlistService

	// Lists provides methods to access user-created release lists.
	// See ListsService for more details.
	Lists ListsService
}

// NewClient initializes a new Discogs API client with the provided functional options.
// It defaults to Anonymous authentication and standard Discogs rate limits.
func NewClient(opts ...Option) *Client {
	c := &Client{
		httpClient: &http.Client{Timeout: 15 * time.Second},
		baseURL:    DefaultBaseURL,
		userAgent:  DefaultUserAgent,
		auth:       &AnonymousAuth{},
	}

	for _, opt := range opts {
		opt(c)
	}

	// Apply Rate Limiter based on Auth status as per Discogs policy:
	// - Authenticated: 60 requests per minute.
	// - Anonymous: 25 requests per minute.
	if c.auth.IsAuthenticated() {
		c.rateLimiter = rate.NewLimiter(rate.Limit(60.0/60.0), 1)
	} else {
		c.rateLimiter = rate.NewLimiter(rate.Limit(25.0/60.0), 1)
	}

	// Initialize services with a reference back to the client.
	c.Database = &databaseService{client: c}
	c.Marketplace = &marketplaceService{client: c}
	c.Inventory = &inventoryService{client: c}
	c.User = &userService{client: c}
	c.Collection = &collectionService{client: c}
	c.Wantlist = &wantlistService{client: c}
	c.Lists = &listsService{client: c}

	return c
}

// BaseURL returns the configured API base URL.
func (c *Client) BaseURL() string {
	return c.baseURL
}

// UserAgent returns the configured User-Agent string.
func (c *Client) UserAgent() string {
	return c.userAgent
}

// Auth returns the currently active authentication strategy.
func (c *Client) Auth() Authenticator {
	return c.auth
}

// HTTPClient returns the underlying net/http client used by the SDK.
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

// do centralizes all HTTP request execution. It enforces rate limiting,
// injects required headers, applies authentication, and handles error responses.
func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) error {
	// Respect the local rate limiter.
	if err := c.rateLimiter.Wait(ctx); err != nil {
		return fmt.Errorf("rate limit wait error: %w", err)
	}

	// Set required headers.
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "application/json")

	// Dynamically apply authentication headers.
	c.auth.Apply(req)

	// Execute the request with context support.
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	// Handle standard Discogs error status codes.
	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusTooManyRequests:
		return ErrRateLimited
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("discogs API error: %d %s", resp.StatusCode, resp.Status)
	}

	// Decode response body if a destination interface is provided.
	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
