package discogs

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetWantlist(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/testuser/wants" {
			t.Errorf("expected path /users/testuser/wants, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"wants": [{"id": 123, "rating": 5, "basic_information": {"title": "Wanted Release"}}]}`)
	}

	server, client := setupTestServer(handler)
	defer server.Close()

	resp, err := client.Wantlist.GetWantlist(context.Background(), "testuser", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resp.Wants) != 1 {
		t.Errorf("expected 1 want, got %d", len(resp.Wants))
	}
	if resp.Wants[0].BasicInformation.Title != "Wanted Release" {
		t.Errorf("expected title Wanted Release, got %s", resp.Wants[0].BasicInformation.Title)
	}
}
