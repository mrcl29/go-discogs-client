package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mrcl29/go-discogs-client"
)

func main() {
	// This example shows how to configure a client with OAuth1.0a credentials.
	// Note: You must obtain the AccessToken and AccessSecret through the 3-legged OAuth flow.

	oauth := &discogs.OAuth1Auth{
		ConsumerKey:    "YOUR_CONSUMER_KEY",
		ConsumerSecret: "YOUR_CONSUMER_SECRET",
		AccessToken:    "USER_ACCESS_TOKEN",
		AccessSecret:   "USER_ACCESS_SECRET",
	}

	client := discogs.NewClient(
		discogs.WithUserAgent("MyAwesomeApp/1.0"),
		discogs.WithAuth(oauth),
	)

	ctx := context.Background()

	// Verify identity
	identity, err := client.User.GetIdentity(ctx)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	fmt.Printf("Authenticated as: %s\n", identity.Username)
	fmt.Printf("Resource URL: %s\n", identity.ResourceURL)
}
