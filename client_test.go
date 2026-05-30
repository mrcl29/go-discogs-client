package discogs

import (
	"net/http"
	"net/http/httptest"
)

func setupTestServer(handler http.HandlerFunc) (*httptest.Server, *Client) {
	server := httptest.NewServer(handler)
	client := NewClient(
		WithBaseURL(server.URL),
		WithUserAgent("TestAgent/1.0"),
	)
	return server, client
}
