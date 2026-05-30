package discogs

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListFolders(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/testuser/collection/folders" {
			t.Errorf("expected path /users/testuser/collection/folders, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"folders": [{"id": 0, "name": "All"}, {"id": 1, "name": "Uncategorized"}]}`)
	}

	server, client := setupTestServer(handler)
	defer server.Close()

	resp, err := client.Collection.ListFolders(context.Background(), "testuser")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resp.Folders) != 2 {
		t.Errorf("expected 2 folders, got %d", len(resp.Folders))
	}
}

func TestGetCollectionItemsByFolder(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/testuser/collection/folders/1/releases" {
			t.Errorf("expected path /users/testuser/collection/folders/1/releases, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"releases": [{"id": 123, "instance_id": 1, "basic_information": {"title": "Release 1"}}]}`)
	}

	server, client := setupTestServer(handler)
	defer server.Close()

	resp, err := client.Collection.GetCollectionItemsByFolder(context.Background(), "testuser", 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resp.Releases) != 1 {
		t.Errorf("expected 1 release, got %d", len(resp.Releases))
	}
}
