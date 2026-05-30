package discogs_test

import (
	"context"
	"fmt"
	"log"

	"github.com/mrcl29/go-discogs-client"
)

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

func ExampleDatabaseService_GetRelease() {
	client := discogs.NewClient(discogs.WithUserAgent("MyApp/1.0"))
	ctx := context.Background()

	// Fetching Nevermind by Nirvana
	release, err := client.Database.GetRelease(ctx, 249504)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Release: %s (%d)\n", release.Title, release.Year)
}

func ExampleWithBaseURL() {
	// Use a mock server for testing
	client := discogs.NewClient(
		discogs.WithUserAgent("TestApp/1.0"),
		discogs.WithBaseURL("http://localhost:8080"),
	)

	fmt.Println(client.BaseURL())
	// Output: http://localhost:8080
}

func ExampleDatabaseService_Search() {
	client := discogs.NewClient(discogs.WithUserAgent("MyApp/1.0"))
	ctx := context.Background()

	// Searching for Nirvana releases
	opts := &discogs.PageOptions{PerPage: 3}
	filters := map[string]string{"artist": "Nirvana", "type": "release"}

	resp, err := client.Database.Search(ctx, "Nevermind", opts, filters)
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range resp.Results {
		fmt.Printf("Result: %s (ID: %d)\n", result.Title, result.ID)
	}
}

func ExampleMarketplaceService_GetListing() {
	client := discogs.NewClient(discogs.WithUserAgent("MyApp/1.0"))
	ctx := context.Background()

	// Fetching a listing
	listing, err := client.Marketplace.GetListing(ctx, 123456, "USD")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listing Status: %s, Price: %.2f %s\n", listing.Status, listing.Price.Value, listing.Price.Currency)
}

func ExampleCollectionService_ListFolders() {
	auth := &discogs.PersonalTokenAuth{Token: "YOUR_TOKEN"}
	client := discogs.NewClient(
		discogs.WithUserAgent("MyApp/1.0"),
		discogs.WithAuth(auth),
	)
	ctx := context.Background()

	// Listing folders for user "rodneyfool"
	resp, err := client.Collection.ListFolders(ctx, "rodneyfool")
	if err != nil {
		log.Fatal(err)
	}

	for _, folder := range resp.Folders {
		fmt.Printf("Folder: %s (%d items)\n", folder.Name, folder.Count)
	}
}
