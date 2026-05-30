package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mrcl29/go-discogs-client"
)

func main() {
	// A personal token is often required for search results
	token := os.Getenv("DISCOGS_TOKEN")
	if token == "" {
		log.Fatal("DISCOGS_TOKEN environment variable is required")
	}

	client := discogs.NewClient(
		discogs.WithUserAgent("DiscogsSearchExample/1.0"),
		discogs.WithAuth(&discogs.PersonalTokenAuth{Token: token}),
	)

	ctx := context.Background()

	fmt.Println("Searching for Nirvana releases...")

	opts := &discogs.PageOptions{
		Page:    1,
		PerPage: 5,
	}

	filters := map[string]string{
		"artist": "Nirvana",
		"type":   "release",
		"format": "Vinyl",
	}

	resp, err := client.Database.Search(ctx, "Nevermind", opts, filters)
	if err != nil {
		log.Fatalf("Search failed: %v", err)
	}

	fmt.Printf("Found %d items (Page %d/%d)\n\n", resp.Pagination.Items, resp.Pagination.Page, resp.Pagination.Pages)

	for _, result := range resp.Results {
		fmt.Printf("[%d] %s\n", result.ID, result.Title)
		fmt.Printf("    Format: %v | Year: %s | Country: %s\n", result.Format, result.Year, result.Country)
		fmt.Println()
	}
}
