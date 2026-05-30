package discogs

import (
	"context"
	"fmt"
	"net/http"
)

// ListsService provides access to user-created release lists.
//
// More info: https://www.discogs.com/developers/#page:user-lists
type ListsService interface {
	// GetUserLists returns a paginated list of lists created by a user.
	// https://www.discogs.com/developers/#page:user-lists,header:user-lists-user-lists
	GetUserLists(ctx context.Context, username string, opts *PageOptions) (*UserListsResponse, error)

	// GetList returns the detailed items and metadata for a specific list.
	// https://www.discogs.com/developers/#page:user-lists,header:user-lists-list
	GetList(ctx context.Context, listID int) (*List, error)
}

type listsService struct {
	client *Client
}

func (s *listsService) GetUserLists(ctx context.Context, username string, opts *PageOptions) (*UserListsResponse, error) {
	baseURL := fmt.Sprintf("%s/users/%s/lists", s.client.baseURL, username)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp UserListsResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *listsService) GetList(ctx context.Context, listID int) (*List, error) {
	url := fmt.Sprintf("%s/lists/%d", s.client.baseURL, listID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var list List
	if err := s.client.do(ctx, req, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
