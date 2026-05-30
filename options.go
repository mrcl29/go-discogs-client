package discogs

import (
	"net/http"
)

// Option defines a functional configuration pattern for the Discogs Client.
type Option func(*Client)

// WithUserAgent overrides the default User-Agent header.
// Discogs mandates a descriptive User-Agent for all requests (e.g., "MyMusicApp/1.0").
func WithUserAgent(agent string) Option {
	return func(c *Client) {
		c.userAgent = agent
	}
}

// WithHTTPClient allows using a custom net/http Client.
// This is useful for injecting custom timeouts, transport configurations, or mocks during testing.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithAuth sets the authentication strategy for the client.
// Choosing an authenticated strategy (Token, Consumer, or OAuth) increases your rate limit to 60 req/min.
func WithAuth(auth Authenticator) Option {
	return func(c *Client) {
		c.auth = auth
	}
}

// WithBaseURL overrides the default Discogs API base URL.
// Primary use case is pointing the client to a mock server or proxy during development.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}
