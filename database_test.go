package discogs

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetRelease(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/releases/123" {
			t.Errorf("expected path /releases/123, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id": 123, "title": "Nevermind", "year": 1991}`)
	}

	server, client := setupTestServer(handler)
	defer server.Close()

	release, err := client.Database.GetRelease(context.Background(), 123)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if release.ID != 123 {
		t.Errorf("expected ID 123, got %d", release.ID)
	}
	if release.Title != "Nevermind" {
		t.Errorf("expected title Nevermind, got %s", release.Title)
	}
}

func TestSearch(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/database/search" {
			t.Errorf("expected path /database/search, got %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("q") != "nirvana" {
			t.Errorf("expected query nirvana, got %s", q.Get("q"))
		}
		if q.Get("type") != "release" {
			t.Errorf("expected type release, got %s", q.Get("type"))
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"results": [{"id": 1, "title": "Nirvana - Nevermind"}]}`)
	}

	server, client := setupTestServer(handler)
	defer server.Close()

	resp, err := client.Database.Search(context.Background(), "nirvana", nil, map[string]string{"type": "release"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resp.Results) != 1 {
		t.Errorf("expected 1 result, got %d", len(resp.Results))
	}
	if resp.Results[0].Title != "Nirvana - Nevermind" {
		t.Errorf("expected title Nirvana - Nevermind, got %s", resp.Results[0].Title)
	}
}
