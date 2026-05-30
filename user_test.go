package discogs

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetIdentity(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/oauth/identity" {
			t.Errorf("expected path /oauth/identity, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"id": 1, "username": "testuser"}`)
	}

	server, client := setupTestServer(handler)
	defer server.Close()

	identity, err := client.User.GetIdentity(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if identity.Username != "testuser" {
		t.Errorf("expected username testuser, got %s", identity.Username)
	}
}

func TestGetProfile(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/testuser" {
			t.Errorf("expected path /users/testuser, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"id": 1, "username": "testuser", "name": "Test User"}`)
	}

	server, client := setupTestServer(handler)
	defer server.Close()

	profile, err := client.User.GetProfile(context.Background(), "testuser")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if profile.Name != "Test User" {
		t.Errorf("expected name Test User, got %s", profile.Name)
	}
}
