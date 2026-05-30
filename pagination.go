package discogs

import (
	"fmt"
	"net/url"
	"strconv"
)

// Pagination contains metadata about a paginated response from the Discogs API.
// It includes current page, total pages, total items, and navigation URLs.
type Pagination struct {
	// Page is the current page number.
	Page int `json:"page"`
	// Pages is the total number of available pages.
	Pages int `json:"pages"`
	// Items is the total number of items across all pages.
	Items int `json:"items"`
	// PerPage is the number of items per page.
	PerPage int `json:"per_page"`
	// URLs contains links to related pages (first, last, next, prev).
	URLs map[string]string `json:"urls"`
}

// PageOptions specifies parameters for paginated and sorted API requests.
// Passing a nil PageOptions to a service method uses the API's default values.
type PageOptions struct {
	// Page is the page number to retrieve (starts at 1).
	Page int
	// PerPage is the number of items per page (up to 100).
	PerPage int
	// Sort is the field name to sort by (e.g., "year", "title").
	Sort string
	// SortOrder is the direction of sorting ("asc" or "desc").
	SortOrder string
}

// ApplyToURL appends the pagination and sorting options as query parameters to the provided base URL.
// It handles URL parsing and parameter encoding automatically.
func (o *PageOptions) ApplyToURL(baseAPIURL string) (string, error) {
	if o == nil {
		return baseAPIURL, nil
	}
	u, err := url.Parse(baseAPIURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse url: %w", err)
	}

	q := u.Query()
	if o.Page > 0 {
		q.Set("page", strconv.Itoa(o.Page))
	}
	if o.PerPage > 0 {
		q.Set("per_page", strconv.Itoa(o.PerPage))
	}
	if o.Sort != "" {
		q.Set("sort", o.Sort)
	}
	if o.SortOrder != "" {
		q.Set("sort_order", o.SortOrder)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
