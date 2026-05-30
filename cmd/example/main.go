package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mrcl29/go-discogs-client"
)

// main demonstrates how to use the discogs client with different authentication methods.
func main() {
	ctx := context.Background()

	// ---------------------------------------------------------
	// Example 1: Anonymous usage (Low Rate Limit, 25 req/min)
	// ---------------------------------------------------------
	anonClient := discogs.NewClient(
		discogs.WithUserAgent("MyMusicApp/1.0"),
	)

	fetchData(ctx, anonClient, "Anonymous Client")

	// ---------------------------------------------------------
	// Example 2: Personal Access Token (High Rate Limit, acts as user)
	// ---------------------------------------------------------
	tokenAuth := &discogs.PersonalTokenAuth{Token: "YOUR_PERSONAL_TOKEN"}
	tokenClient := discogs.NewClient(
		discogs.WithUserAgent("MyMusicApp/1.0"),
		discogs.WithAuth(tokenAuth),
	)

	fetchData(ctx, tokenClient, "Personal Token Client")

	// ---------------------------------------------------------
	// Example 3: Full OAuth 1.0a (Required for user write actions)
	// ---------------------------------------------------------
	oauth := &discogs.OAuth1Auth{
		ConsumerKey:    "YOUR_CONSUMER_KEY",
		ConsumerSecret: "YOUR_CONSUMER_SECRET",
		AccessToken:    "USER_ACCESS_TOKEN",
		AccessSecret:   "USER_ACCESS_SECRET",
	}
	oauthClient := discogs.NewClient(
		discogs.WithUserAgent("MyMusicApp/1.0"),
		discogs.WithAuth(oauth),
	)

	fetchData(ctx, oauthClient, "OAuth Client")

	// ---------------------------------------------------------
	// Example 4: Advanced usage (Collection Management)
	// ---------------------------------------------------------
	// Listing user collection folders
	folders, err := tokenClient.Collection.ListFolders(ctx, "rodneyfool")
	if err != nil {
		log.Printf("[Collection] Error: %v\n", err)
	} else {
		fmt.Println("--- Collection Folders ---")
		for _, folder := range folders.Folders {
			fmt.Printf("- %s (ID: %d, Count: %d)\n", folder.Name, folder.ID, folder.Count)
		}
		fmt.Println()
	}
}

// fetchData is a helper function that attempts to retrieve a release from Discogs using the provided client.
func fetchData(ctx context.Context, client *discogs.Client, label string) {
	fmt.Printf("--- Testing %s ---\n", label)

	// Getting Nirvana's Nevermind release (ID: 249504)
	release, err := client.Database.GetRelease(ctx, 249504)
	if err != nil {
		log.Printf("[%s] Error: %v\n\n", label, err)
		return
	}

	fmt.Printf("[%s] Success! Release: %s (%d)\n\n", label, release.Title, release.Year)
}
