package discogs_test

import (
	"context"
	"fmt"
	"log"

	"github.com/mrcl29/go-discogs-client"
)

// ExampleNewClient demonstrates how to initialize a new Discogs client
// with custom options such as User-Agent and Authentication.
func ExampleNewClient() {
	// Initialize a client with a custom User-Agent and Personal Access Token
	auth := &discogs.PersonalTokenAuth{Token: "YOUR_TOKEN"}
	client := discogs.NewClient(
		discogs.WithUserAgent("MyApp/1.0"),
		discogs.WithAuth(auth),
	)

	fmt.Println(client.UserAgent())
	// Output: MyApp/1.0
}

// ExampleWithBaseURL shows how to point the client to a different URL,
// which is particularly useful for integration testing with mock servers.
func ExampleWithBaseURL() {
	// Use a mock server for testing
	client := discogs.NewClient(
		discogs.WithUserAgent("TestApp/1.0"),
		discogs.WithBaseURL("http://localhost:8080"),
	)

	fmt.Println(client.BaseURL())
	// Output: http://localhost:8080
}

// ExampleDatabaseService_GetRelease demonstrates fetching a specific release by its ID.
func ExampleDatabaseService_GetRelease() {
	client := discogs.NewClient(discogs.WithUserAgent("MyApp/1.0"))
	ctx := context.Background()

	// Fetching Nevermind by Nirvana (ID: 249504)
	release, err := client.Database.GetRelease(ctx, 249504)
	if err != nil {
		// In a real application, handle the error appropriately
		return
	}

	fmt.Printf("Release: %s (%d)\n", release.Title, release.Year)
}

// ExampleDatabaseService_Search demonstrates performing an authenticated search
// with filters and pagination.
func ExampleDatabaseService_Search() {
	// Authentication is highly recommended for search to get better results
	auth := &discogs.PersonalTokenAuth{Token: "YOUR_TOKEN"}
	client := discogs.NewClient(
		discogs.WithUserAgent("MyApp/1.0"),
		discogs.WithAuth(auth),
	)
	ctx := context.Background()

	// Searching for Nirvana releases on Vinyl
	opts := &discogs.PageOptions{Page: 1, PerPage: 5}
	filters := map[string]string{
		"artist": "Nirvana",
		"type":   "release",
		"format": "Vinyl",
	}

	resp, err := client.Database.Search(ctx, "Nevermind", opts, filters)
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range resp.Results {
		fmt.Printf("Found: %s (ID: %d)\n", result.Title, result.ID)
	}
}

// ExampleMarketplaceService_GetListing demonstrates retrieving details for a single marketplace listing.
func ExampleMarketplaceService_GetListing() {
	client := discogs.NewClient(discogs.WithUserAgent("MyApp/1.0"))
	ctx := context.Background()

	// Fetching a specific listing by ID and requesting price in USD
	listing, err := client.Marketplace.GetListing(ctx, 123456, "USD")
	if err != nil {
		return
	}

	fmt.Printf("Listing Status: %s, Price: %.2f %s\n",
		listing.Status,
		listing.Price.Value,
		listing.Price.Currency)
}

// ExampleCollectionService_ListFolders demonstrates listing all folders in a user's collection.
func ExampleCollectionService_ListFolders() {
	auth := &discogs.PersonalTokenAuth{Token: "YOUR_TOKEN"}
	client := discogs.NewClient(
		discogs.WithUserAgent("MyApp/1.0"),
		discogs.WithAuth(auth),
	)
	ctx := context.Background()

	// Listing folders for a specific user
	resp, err := client.Collection.ListFolders(ctx, "rodneyfool")
	if err != nil {
		log.Fatal(err)
	}

	for _, folder := range resp.Folders {
		fmt.Printf("Folder: %s (%d items)\n", folder.Name, folder.Count)
	}
}

// ExampleCollectionService_GetCollectionValue demonstrates calculating the estimated
// monetary value of a user's collection.
func ExampleCollectionService_GetCollectionValue() {
	auth := &discogs.PersonalTokenAuth{Token: "YOUR_TOKEN"}
	client := discogs.NewClient(
		discogs.WithUserAgent("MyApp/1.0"),
		discogs.WithAuth(auth),
	)
	ctx := context.Background()

	// Get collection valuation for a user
	value, err := client.Collection.GetCollectionValue(ctx, "rodneyfool")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Collection Value (Median): %s\n", value.Median)
}

// ExampleUserService_GetIdentity demonstrates verifying current authentication
// and retrieving the authenticated user's basic information.
func ExampleUserService_GetIdentity() {
	// OAuth or Token authentication is required for this endpoint
	auth := &discogs.PersonalTokenAuth{Token: "YOUR_TOKEN"}
	client := discogs.NewClient(
		discogs.WithUserAgent("MyApp/1.0"),
		discogs.WithAuth(auth),
	)
	ctx := context.Background()

	identity, err := client.User.GetIdentity(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Authenticated as: %s\n", identity.Username)
}
