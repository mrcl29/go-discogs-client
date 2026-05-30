package discogs

import (
	"context"
	"fmt"
	"net/http"
)

// DatabaseService provides access to Discogs' vast music database including
// Releases, Masters, Artists, and Labels, as well as a powerful search engine.
//
// More info: https://www.discogs.com/developers/#page:database
type DatabaseService interface {
	// GetRelease fetches detailed metadata for a specific release by its unique ID.
	// https://www.discogs.com/developers/#page:database,header:database-release
	GetRelease(ctx context.Context, releaseID int) (*Release, error)

	// GetMaster fetches metadata for a master release (a set of similar releases).
	// https://www.discogs.com/developers/#page:database,header:database-master-release
	GetMaster(ctx context.Context, masterID int) (*Master, error)

	// GetMasterVersions retrieves a paginated list of all releases associated with a master release.
	// https://www.discogs.com/developers/#page:database,header:database-master-release-versions
	GetMasterVersions(ctx context.Context, masterID int, opts *PageOptions) (*MasterVersionsResponse, error)

	// GetArtist fetches detailed information about a person or group in the database.
	// https://www.discogs.com/developers/#page:database,header:database-artist
	GetArtist(ctx context.Context, artistID int) (*Artist, error)

	// GetArtistReleases returns a paginated list of releases and masters associated with an artist.
	// https://www.discogs.com/developers/#page:database,header:database-artist-releases
	GetArtistReleases(ctx context.Context, artistID int, opts *PageOptions) (*ArtistReleasesResponse, error)

	// GetLabel fetches information about a specific music label or parent company.
	// https://www.discogs.com/developers/#page:database,header:database-label
	GetLabel(ctx context.Context, labelID int) (*Label, error)

	// GetLabelReleases returns a paginated list of releases associated with a specific label.
	// https://www.discogs.com/developers/#page:database,header:database-all-label-releases
	GetLabelReleases(ctx context.Context, labelID int, opts *PageOptions) (*LabelReleasesResponse, error)

	// Search issues a query to the Discogs database. Authenticated requests are required.
	// Filters can be used to narrow down results (e.g., {"type": "release", "artist": "Nirvana"}).
	// https://www.discogs.com/developers/#page:database,header:database-search
	Search(ctx context.Context, query string, opts *PageOptions, filters map[string]string) (*SearchResponse, error)

	// GetReleaseRating retrieves community rating statistics (average and count) for a release.
	// https://www.discogs.com/developers/#page:database,header:database-community-release-rating
	GetReleaseRating(ctx context.Context, releaseID int) (*Rating, error)

	// GetReleaseStats retrieves the number of users who "have" or "want" a given release.
	// https://www.discogs.com/developers/#page:database,header:database-release-stats
	GetReleaseStats(ctx context.Context, releaseID int) (*Stats, error)
}

type databaseService struct {
	client *Client
}

func (s *databaseService) GetRelease(ctx context.Context, releaseID int) (*Release, error) {
	url := fmt.Sprintf("%s/releases/%d", s.client.baseURL, releaseID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var release Release
	if err := s.client.do(ctx, req, &release); err != nil {
		return nil, err
	}
	return &release, nil
}

func (s *databaseService) GetMaster(ctx context.Context, masterID int) (*Master, error) {
	url := fmt.Sprintf("%s/masters/%d", s.client.baseURL, masterID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var master Master
	if err := s.client.do(ctx, req, &master); err != nil {
		return nil, err
	}
	return &master, nil
}

func (s *databaseService) GetMasterVersions(ctx context.Context, masterID int, opts *PageOptions) (*MasterVersionsResponse, error) {
	baseURL := fmt.Sprintf("%s/masters/%d/versions", s.client.baseURL, masterID)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp MasterVersionsResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *databaseService) GetArtist(ctx context.Context, artistID int) (*Artist, error) {
	url := fmt.Sprintf("%s/artists/%d", s.client.baseURL, artistID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var artist Artist
	if err := s.client.do(ctx, req, &artist); err != nil {
		return nil, err
	}
	return &artist, nil
}

func (s *databaseService) GetArtistReleases(ctx context.Context, artistID int, opts *PageOptions) (*ArtistReleasesResponse, error) {
	baseURL := fmt.Sprintf("%s/artists/%d/releases", s.client.baseURL, artistID)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp ArtistReleasesResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *databaseService) GetLabel(ctx context.Context, labelID int) (*Label, error) {
	url := fmt.Sprintf("%s/labels/%d", s.client.baseURL, labelID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var label Label
	if err := s.client.do(ctx, req, &label); err != nil {
		return nil, err
	}
	return &label, nil
}

func (s *databaseService) GetLabelReleases(ctx context.Context, labelID int, opts *PageOptions) (*LabelReleasesResponse, error) {
	baseURL := fmt.Sprintf("%s/labels/%d/releases", s.client.baseURL, labelID)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	var resp LabelReleasesResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *databaseService) Search(ctx context.Context, query string, opts *PageOptions, filters map[string]string) (*SearchResponse, error) {
	baseURL := fmt.Sprintf("%s/database/search", s.client.baseURL)
	finalURL, err := opts.ApplyToURL(baseURL)
	if err != nil {
		return nil, err
	}

	// Build the search query with parameters.
	u, _ := http.NewRequest(http.MethodGet, finalURL, nil)
	q := u.URL.Query()
	if query != "" {
		q.Set("q", query)
	}
	for k, v := range filters {
		q.Set(k, v)
	}
	u.URL.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.URL.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp SearchResponse
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *databaseService) GetReleaseRating(ctx context.Context, releaseID int) (*Rating, error) {
	url := fmt.Sprintf("%s/releases/%d/rating", s.client.baseURL, releaseID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Rating Rating `json:"rating"`
	}
	if err := s.client.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Rating, nil
}

func (s *databaseService) GetReleaseStats(ctx context.Context, releaseID int) (*Stats, error) {
	url := fmt.Sprintf("%s/releases/%d/stats", s.client.baseURL, releaseID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var stats Stats
	if err := s.client.do(ctx, req, &stats); err != nil {
		return nil, err
	}
	return &stats, nil
}
