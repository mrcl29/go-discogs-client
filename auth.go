package discogs

import (
	"fmt"
	"net/http"
	"time"
)

// Authenticator defines the interface for different Discogs API authentication strategies.
// Implementation must handle the injection of required headers or query parameters into requests.
type Authenticator interface {
	// Apply adds the required authentication headers to the provided HTTP request.
	Apply(req *http.Request)
	// IsAuthenticated returns true if this strategy identifies the client (enabling higher rate limits).
	IsAuthenticated() bool
}

// AnonymousAuth implements the unauthenticated request strategy.
// It provides no identity to Discogs, resulting in a lower rate limit (25 req/min)
// and restricted access to user data and images.
type AnonymousAuth struct{}

// Apply does nothing for anonymous requests.
func (a *AnonymousAuth) Apply(req *http.Request) {}

// IsAuthenticated always returns false for AnonymousAuth.
func (a *AnonymousAuth) IsAuthenticated() bool { return false }

// PersonalTokenAuth uses a single generated User Token for authentication.
// This is the simplest way to access your own account data and images
// while enjoying a higher rate limit (60 req/min).
type PersonalTokenAuth struct {
	// Token is the secret Discogs personal access token.
	Token string
}

// Apply adds the Token to the Authorization header using the Discogs schema.
func (a *PersonalTokenAuth) Apply(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Discogs token=%s", a.Token))
}

// IsAuthenticated always returns true for PersonalTokenAuth.
func (a *PersonalTokenAuth) IsAuthenticated() bool { return true }

// ConsumerKeyAuth uses the application's key and secret for authentication.
// This identifies the application but not a specific user. It grants higher
// rate limits and access to images, but not private user-specific data.
type ConsumerKeyAuth struct {
	// Key is the Discogs consumer key.
	Key string
	// Secret is the Discogs consumer secret.
	Secret string
}

// Apply adds the Key and Secret to the Authorization header.
func (a *ConsumerKeyAuth) Apply(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Discogs key=%s, secret=%s", a.Key, a.Secret))
}

// IsAuthenticated always returns true for ConsumerKeyAuth.
func (a *ConsumerKeyAuth) IsAuthenticated() bool { return true }

// OAuth1Auth implements the full Discogs OAuth 1.0a flow using the PLAINTEXT signature method.
// This strategy is required for performing actions on behalf of any Discogs user,
// such as adding items to their collection or wantlist.
type OAuth1Auth struct {
	// ConsumerKey is the Discogs consumer key.
	ConsumerKey string
	// ConsumerSecret is the Discogs consumer secret.
	ConsumerSecret string
	// AccessToken is the user's OAuth access token.
	AccessToken string
	// AccessSecret is the user's OAuth access token secret.
	AccessSecret string
}

// Apply adds the complex OAuth 1.0a Authorization header to the request.
func (a *OAuth1Auth) Apply(req *http.Request) {
	timestamp := time.Now().Unix()
	// Discogs PLAINTEXT signature is built by concatenating the secrets.
	signature := fmt.Sprintf("%s&%s", a.ConsumerSecret, a.AccessSecret)

	authHeader := fmt.Sprintf(
		`OAuth oauth_consumer_key="%s", oauth_nonce="%d", oauth_token="%s", oauth_signature="%s", oauth_signature_method="PLAINTEXT", oauth_timestamp="%d"`,
		a.ConsumerKey, timestamp, a.AccessToken, signature, timestamp,
	)
	req.Header.Set("Authorization", authHeader)
}

// IsAuthenticated always returns true for OAuth1Auth.
func (a *OAuth1Auth) IsAuthenticated() bool { return true }
