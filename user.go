package discogs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// UserService provides access to user identity, profiles, and account-specific activity.
//
// More info: https://www.discogs.com/developers/#page:user-identity
type UserService interface {
	// GetIdentity retrieves basic information about the currently authenticated user.
	// This endpoint is the primary way to verify if your OAuth credentials are valid.
	// https://www.discogs.com/developers/#page:user-identity,header:user-identity-identity
	GetIdentity(ctx context.Context) (*UserRef, error)

	// GetProfile retrieves a user's public profile data by their username.
	// If authenticated as the requested user, additional private fields like email will be visible.
	// https://www.discogs.com/developers/#page:user-identity,header:user-identity-profile
	GetProfile(ctx context.Context, username string) (*Profile, error)

	// UpdateProfile edits a user's profile data. Requires authentication as the user.
	// The data map can contain fields like "name", "home_page", "location", "profile", or "curr_abbr".
	// https://www.discogs.com/developers/#page:user-identity,header:user-identity-profile-post
	UpdateProfile(ctx context.Context, username string, data map[string]interface{}) (*Profile, error)

	// GetUserSubmissions retrieves a paginated list of all edits made by a user to the database.
	// https://www.discogs.com/developers/#page:user-identity,header:user-identity-user-submissions
	GetUserSubmissions(ctx context.Context, username string, opts *PageOptions) (*SubmissionsResponse, error)

	// GetUserContributions retrieves a paginated list of releases, labels, and artists submitted by a user.
	// https://www.discogs.com/developers/#page:user-identity,header:user-identity-user-contributions
	GetUserContributions(ctx context.Context, username string, opts *PageOptions) (*ContributionsResponse, error)
}

type userService struct {
	client *Client
}

func (s *userService) GetIdentity(ctx context.Context) (*UserRef, error) {
	url := fmt.Sprintf("%s/oauth/identity", s.client.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var identity UserRef
	if err := s.client.do(ctx, req, &identity); err != nil {
		return nil, err
	}
	return &identity, nil
}

func (s *userService) GetProfile(ctx context.Context, username string) (*Profile, error) {
	url := fmt.Sprintf("%s/users/%s", s.client.baseURL, username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var profile Profile
	if err := s.client.do(ctx, req, &profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *userService) UpdateProfile(ctx context.Context, username string, data map[string]interface{}) (*Profile, error) {
	url := fmt.Sprintf("%s/users/%s", s.client.baseURL, username)
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var profile Profile
	if err := s.client.do(ctx, req, &profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *userService) GetUserSubmissions(ctx context.Context, username string, opts *PageOptions) (*SubmissionsResponse, error) {
	baseURL := fmt.Sprintf("%s/users/%s/submissions", s.client.baseURL, username)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp SubmissionsResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *userService) GetUserContributions(ctx context.Context, username string, opts *PageOptions) (*ContributionsResponse, error) {
	baseURL := fmt.Sprintf("%s/users/%s/contributions", s.client.baseURL, username)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp ContributionsResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
