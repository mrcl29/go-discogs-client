# go-discogs-client

[![Go Version](https://img.shields.io/github/go-mod/go-version/mrcl29/go-discogs-client)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/mrcl29/go-discogs-client)
[![Go Reference](https://pkg.go.dev/badge/github.com/mrcl29/go-discogs-client.svg)](https://pkg.go.dev/github.com/mrcl29/go-discogs-client)
[![CI](https://github.com/mrcl29/go-discogs-client/actions/workflows/ci.yml/badge.svg)](https://github.com/mrcl29/go-discogs-client/actions/workflows/ci.yml)

A professional, feature-rich, and idiomatic Go client for the [Discogs API v2.0](https://www.discogs.com/developers).

Built with scalability and idiomatic Go in mind, this library provides a complete interface to Discogs' vast music database and marketplace, handling authentication and rate limiting automatically so you can focus on your application logic.

## ✨ Features

- **Full Service Coverage**: Implementation for `Database`, `Marketplace`, `Inventory`, `User`, `Collection`, `Wantlist`, and `Lists` services.
- **Multiple Authentication Support**: Native support for Anonymous, Personal Access Token, Consumer Key/Secret, and full OAuth 1.0a flows.
- **Automatic Rate Limiting**: Intelligent local rate limiting based on your authentication tier (60 req/min for auth, 25 req/min for unauth) using token buckets.
- **Context Support**: Every API call supports `context.Context` for proper timeout and cancellation management.
- **Functional Options**: Clean client initialization using the functional options pattern.
- **Strongly Typed Models**: Comprehensive and documented Go structs for all API responses, ensuring type safety and easy development.
- **Modern Go**: Leverages Go 1.26.3 features and follows standard project layout conventions.

## 🚀 Installation

```bash
go get github.com/mrcl29/go-discogs-client
```

## 🛠️ Quick Start

### 1. Simple Metadata Retrieval (Anonymous)

```go
package main

import (
    "context"
    "fmt"
    "github.com/mrcl29/go-discogs-client"
)

func main() {
    // Initialize a simple client (limited to 25 req/min)
    client := discogs.NewClient(discogs.WithUserAgent("MyMusicApp/1.0"))
    ctx := context.Background()

    // Fetch a release by ID
    release, err := client.Database.GetRelease(ctx, 249504) // Nirvana - Nevermind
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Release: %s (%d)\n", release.Title, release.Year)
}
```

### 2. Authenticated Search & Pagination

```go
package main

import (
    "context"
    "fmt"
    "github.com/mrcl29/go-discogs-client"
)

func main() {
    auth := &discogs.PersonalTokenAuth{Token: "YOUR_PERSONAL_TOKEN"}
    client := discogs.NewClient(
        discogs.WithUserAgent("MyMusicApp/1.0"),
        discogs.WithAuth(auth),
    )
    ctx := context.Background()

    // Search with filters and pagination
    opts := &discogs.PageOptions{PerPage: 5}
    filters := map[string]string{"format": "Vinyl", "year": "1991"}
    
    res, err := client.Database.Search(ctx, "Nevermind", opts, filters)
    if err != nil {
        panic(err)
    }

    for _, result := range res.Results {
        fmt.Printf("Found: %s (ID: %d)\n", result.Title, result.ID)
    }
}
```

### 3. Collection Management (Requires Token/OAuth)

```go
package main

import (
    "context"
    "github.com/mrcl29/go-discogs-client"
)

func main() {
    // ... client initialization with auth ...
    
    // Add a release to your 'Uncategorized' folder
    item, err := client.Collection.AddToCollectionFolder(ctx, "my_username", 1, 249504)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Added with Instance ID: %d\n", item.InstanceID)
}
```

## 🏗️ Architecture & Design

### Service-Oriented Design

The client is organized into domain-specific services, accessible directly from the main `Client` struct:

- `client.Database`: Search, Releases, Masters, Artists, and Labels.
- `client.Marketplace`: Listings, Orders, Price Suggestions, and Fee calculation.
- `client.Inventory`: Inventory Export (CSV) and Batch Uploads.
- `client.User`: Identity, Profiles, Submissions, and Contributions.
- `client.Collection`: Folder management, item instances, and collection valuation.
- `client.Wantlist`: Manage your personal release wantlist.
- `client.Lists`: Access user-created lists.

### Resilience & Thread Safety

The client is designed to be thread-safe and includes a built-in rate limiter that prevents your application from hitting the API limits by pacing requests locally. All network operations use `http.Client` timeouts and support Go's `context` package for cancellation.

## 📚 Examples

Check the [examples/](examples/) directory for more detailed use cases:

- [Authentication Flows](examples/auth/): Setting up different auth strategies.
- [Database Search](examples/search/): Advanced searching with filters.
- [Marketplace & Pricing](examples/marketplace/): Managing listings and checking prices.
- [Collection Operations](examples/collection/): Organizing your personal database.

## 🧪 Testing

The library is designed with testability as a core principle. All services can be easily tested using the provided mock server patterns.

Run tests:

```bash
go test -v ./...
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
