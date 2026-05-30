package discogs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// WantlistService provides access to user wantlist management.
// It allows users to track releases they want to buy and add personal notes to them.
//
// More info: https://www.discogs.com/developers/#page:user-wantlist
type WantlistService interface {
	// GetWantlist returns a paginated list of releases in a user's wantlist.
	// https://www.discogs.com/developers/#page:user-wantlist,header:user-wantlist-wantlist
	GetWantlist(ctx context.Context, username string, opts *PageOptions) (*WantlistResponse, error)

	// AddToWantlist adds a release to a user's wantlist. Requires OAuth.
	// https://www.discogs.com/developers/#page:user-wantlist,header:user-wantlist-add-to-wantlist
	AddToWantlist(ctx context.Context, username string, releaseID int, notes string, rating int) (*Want, error)

	// UpdateWantlist edits the notes or rating of a release in a user's wantlist. Requires OAuth.
	// https://www.discogs.com/developers/#page:user-wantlist,header:user-wantlist-add-to-wantlist-post
	UpdateWantlist(ctx context.Context, username string, releaseID int, notes string, rating int) (*Want, error)

	// DeleteFromWantlist removes a release from a user's wantlist. Requires OAuth.
	// https://www.discogs.com/developers/#page:user-wantlist,header:user-wantlist-add-to-wantlist-delete
	DeleteFromWantlist(ctx context.Context, username string, releaseID int) error
}

type wantlistService struct {
	client *Client
}

func (s *wantlistService) GetWantlist(ctx context.Context, username string, opts *PageOptions) (*WantlistResponse, error) {
	baseURL := fmt.Sprintf("%s/users/%s/wants", s.client.baseURL, username)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp WantlistResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *wantlistService) AddToWantlist(ctx context.Context, username string, releaseID int, notes string, rating int) (*Want, error) {
	url := fmt.Sprintf("%s/users/%s/wants/%d", s.client.baseURL, username, releaseID)
	data := map[string]interface{}{}
	if notes != "" {
		data["notes"] = notes
	}
	if rating > 0 {
		data["rating"] = rating
	}
	body, _ := json.Marshal(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var want Want
	if err := s.client.do(ctx, req, &want); err != nil {
		return nil, err
	}
	return &want, nil
}

func (s *wantlistService) UpdateWantlist(ctx context.Context, username string, releaseID int, notes string, rating int) (*Want, error) {
	url := fmt.Sprintf("%s/users/%s/wants/%d", s.client.baseURL, username, releaseID)
	data := map[string]interface{}{}
	if notes != "" {
		data["notes"] = notes
	}
	if rating > 0 {
		data["rating"] = rating
	}
	body, _ := json.Marshal(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var want Want
	if err := s.client.do(ctx, req, &want); err != nil {
		return nil, err
	}
	return &want, nil
}

func (s *wantlistService) DeleteFromWantlist(ctx context.Context, username string, releaseID int) error {
	url := fmt.Sprintf("%s/users/%s/wants/%d", s.client.baseURL, username, releaseID)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	return s.client.do(ctx, req, nil)
}
