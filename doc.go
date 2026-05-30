// Package discogs provides a professional, feature-rich, and idiomatic Go client for the Discogs API v2.0.
//
// The client supports various authentication strategies, automatic rate limiting,
// and gives you access to the rich Discogs database, marketplace, and user-specific
// data such as collections and wantlists. It is designed to be thread-safe,
// context-aware, and follows standard Go conventions for clean and scalable integration.
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
// # Rate Limiting
//
// The client automatically handles rate limiting based on your authentication status
// (60 req/min for authenticated, 25 req/min for anonymous). It uses a local token
// bucket to pace requests, ensuring your application remains resilient and avoids
// 429 Too Many Requests errors from Discogs servers.
//
// # Services
//
// The client is organized into logical services, each representing a domain-specific
// area of the Discogs API:
//
//   - Database: Search and browse Releases, Masters, Artists, and Labels.
//   - Marketplace: Manage listings, orders, and price suggestions.
//   - Inventory: Batch operations via CSV export and upload.
//   - User: Identity verification and profile management.
//   - Collection: Organize and value your personal record collection.
//   - Wantlist: Manage items you are looking to acquire.
//   - Lists: Access user-curated lists of releases.
//
// # Pagination
//
// Many Discogs endpoints return paginated results. This library provides a PageOptions
// struct to easily specify page numbers, items per page (up to 100), and sorting fields/order.
// All paginated responses include a Pagination object with metadata for easy navigation.
//
// # User-Agent Requirement
//
// Discogs mandates a unique User-Agent header for all requests. The client defaults
// to "GoDiscogsClient/1.0", but it is highly recommended to set your own using the
// WithUserAgent(string) option to identify your application clearly.
//
// # Examples
//
// For complete, runnable examples, see the examples/ directory in the source repository.
//
// Basic anonymous client:
//
//	client := discogs.NewClient(discogs.WithUserAgent("MyMusicApp/1.0"))
//
// Authenticated client with a Personal Token:
//
//	auth := &discogs.PersonalTokenAuth{Token: "your_secret_token"}
//	client := discogs.NewClient(
//		discogs.WithUserAgent("MyMusicApp/1.0"),
//		discogs.WithAuth(auth),
//	)
package discogs
