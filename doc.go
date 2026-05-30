// Package discogs provides a professional, feature-rich, and idiomatic Go client for the Discogs API v2.0.
//
// The client supports various authentication strategies, automatic rate limiting,
// and gives you access to the rich Discogs database, marketplace, and user-specific
// data such as collections and wantlists. It is designed to be thread-safe,
// context-aware, and follows standard Go conventions for clean and scalable integration.
//
// # Getting Started
//
// To begin using the SDK, initialize a new client using the NewClient constructor.
// By default, the client is configured for anonymous access, which is subject to
// stricter rate limits and cannot access user-specific data or images.
//
//	client := discogs.NewClient(discogs.WithUserAgent("MyMusicApp/1.0"))
//
// # Authentication
//
// Discogs provides several levels of authentication, each with different rate limits
// and capabilities. The package supports all of them via the Authenticator interface:
//
//  1. Anonymous: No authentication. Limited to 25 requests per minute.
//     Ideal for simple metadata retrieval where user identity isn't required.
//     Images are not accessible via this method.
//  2. Personal Access Token: The simplest way to authenticate for a single user.
//     Grants 60 requests per minute and access to the user's private data and images.
//  3. Consumer Key and Secret: Authenticates the application but not a specific user.
//     Useful for search-intensive apps and image retrieval with a higher rate limit.
//  4. OAuth 1.0a: Full three-legged authentication. Required for actions that
//     modify user data on behalf of any Discogs user (e.g., managing collections).
//
// # Automatic Rate Limiting
//
// The client automatically handles rate limiting based on your authentication status
// (60 req/min for authenticated, 25 req/min for anonymous). It uses a local token
// bucket to pace requests, ensuring your application remains resilient and avoids
// 429 Too Many Requests errors from Discogs servers. Network calls will block
// until a token is available or the context is cancelled.
//
// # Service-Oriented Design
//
// The SDK is organized into domain-specific services, accessible via the main Client struct:
//
//   - Database: Search and browse Releases, Masters, Artists, and Labels.
//   - Marketplace: Manage listings, orders, and price suggestions.
//   - Inventory: Batch operations via CSV export and upload.
//   - User: Identity verification and profile management.
//   - Collection: Organize and value your personal record collection.
//   - Wantlist: Manage items you are looking to acquire.
//   - Lists: Access user-curated lists of releases.
//
// Each service method accepts a context.Context as its first argument for proper
// timeout and cancellation management.
//
// # Concurrency and Thread Safety
//
// The Client and its attached services are designed to be thread-safe and can be
// shared across multiple goroutines. The internal HTTP client and rate limiter
// handle concurrent access safely.
//
// # Functional Options
//
// Client configuration follows the functional options pattern, allowing for clean
// and extensible initialization:
//
//	auth := &discogs.PersonalTokenAuth{Token: "your_secret_token"}
//	client := discogs.NewClient(
//		discogs.WithUserAgent("MyMusicApp/1.0"),
//		discogs.WithAuth(auth),
//		discogs.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
//	)
package discogs
