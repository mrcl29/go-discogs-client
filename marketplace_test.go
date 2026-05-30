package discogs

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetListing(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/marketplace/listings/123" {
			t.Errorf("expected path /marketplace/listings/123, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id": 123, "status": "For Sale", "price": {"currency": "USD", "value": 10.5}}`)
	}

	server, client := setupTestServer(handler)
	defer server.Close()

	listing, err := client.Marketplace.GetListing(context.Background(), 123, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if listing.ID != 123 {
		t.Errorf("expected ID 123, got %d", listing.ID)
	}
	if listing.Price.Value != 10.5 {
		t.Errorf("expected price 10.5, got %f", listing.Price.Value)
	}
}

func TestGetOrder(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/marketplace/orders/1-1" {
			t.Errorf("expected path /marketplace/orders/1-1, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id": "1-1", "status": "New Order"}`)
	}

	server, client := setupTestServer(handler)
	defer server.Close()

	order, err := client.Marketplace.GetOrder(context.Background(), "1-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if order.ID != "1-1" {
		t.Errorf("expected ID 1-1, got %s", order.ID)
	}
}
