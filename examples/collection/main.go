package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mrcl29/go-discogs-client"
)

func main() {
	token := os.Getenv("DISCOGS_TOKEN")
	username := os.Getenv("DISCOGS_USERNAME") // Example: rodneyfool
	if token == "" || username == "" {
		log.Fatal("DISCOGS_TOKEN and DISCOGS_USERNAME environment variables are required")
	}

	client := discogs.NewClient(
		discogs.WithUserAgent("DiscogsCollectionExample/1.0"),
		discogs.WithAuth(&discogs.PersonalTokenAuth{Token: token}),
	)

	ctx := context.Background()

	// 1. Get Collection Value
	fmt.Printf("Calculating collection value for user '%s'...\n", username)
	value, err := client.Collection.GetCollectionValue(ctx, username)
	if err != nil {
		log.Printf("Failed to get collection value: %v", err)
	} else {
		fmt.Printf("Collection Value: Min: %s | Median: %s | Max: %s\n",
			value.Minimum, value.Median, value.Maximum)
	}
	fmt.Println()

	// 2. List Folders
	fmt.Println("Listing collection folders...")
	folders, err := client.Collection.ListFolders(ctx, username)
	if err != nil {
		log.Fatalf("Failed to list folders: %v", err)
	}

	for _, folder := range folders.Folders {
		fmt.Printf("- [%d] %s: %d items\n", folder.ID, folder.Name, folder.Count)
	}
	fmt.Println()

	// 3. List items in the 'Uncategorized' folder (ID: 1)
	fmt.Println("Listing items in 'Uncategorized' folder...")
	items, err := client.Collection.GetCollectionItemsByFolder(ctx, username, 1, &discogs.PageOptions{PerPage: 5})
	if err != nil {
		log.Printf("Failed to list folder items: %v", err)
	} else {
		for _, item := range items.Releases {
			fmt.Printf("- %s by %s (Instance ID: %d)\n",
				item.BasicInformation.Title,
				item.BasicInformation.Artists[0].Name,
				item.InstanceID)
		}
	}
}
