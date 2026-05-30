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
	if token == "" {
		log.Fatal("DISCOGS_TOKEN environment variable is required")
	}

	client := discogs.NewClient(
		discogs.WithUserAgent("DiscogsMarketplaceExample/1.0"),
		discogs.WithAuth(&discogs.PersonalTokenAuth{Token: token}),
	)

	ctx := context.Background()

	// 1. Get Price Suggestions for a Release
	releaseID := 249504 // Nirvana - Nevermind
	fmt.Printf("Fetching price suggestions for Release ID %d...\n", releaseID)

	suggestions, err := client.Marketplace.GetPriceSuggestions(ctx, releaseID)
	if err != nil {
		log.Printf("Failed to get suggestions: %v", err)
	} else {
		for condition, price := range suggestions {
			fmt.Printf("- %s: %.2f %s\n", condition, price.Value, price.Currency)
		}
	}
	fmt.Println()

	// 2. Calculate Fee
	price := 50.0
	currency := "USD"
	fmt.Printf("Calculating fee for a %.2f %s sale...\n", price, currency)

	fee, err := client.Marketplace.GetFee(ctx, price, currency)
	if err != nil {
		log.Printf("Failed to calculate fee: %v", err)
	} else {
		fmt.Printf("Fee: %.2f %s\n", fee.Value, fee.Currency)
	}
	fmt.Println()

	// 3. List recent orders (Requires token/OAuth)
	fmt.Println("Listing recent orders...")
	orders, err := client.Marketplace.ListOrders(ctx, &discogs.PageOptions{PerPage: 3}, "")
	if err != nil {
		log.Printf("Failed to list orders: %v", err)
	} else {
		for _, order := range orders.Orders {
			fmt.Printf("- Order ID: %s | Status: %s | Total: %.2f %s\n",
				order.ID, order.Status, order.Total.Value, order.Total.Currency)
		}
	}
}
