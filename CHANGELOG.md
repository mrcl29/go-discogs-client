# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-05-30

### Added

- Initial professional release of the Go Discogs Client.
- Full service coverage for Discogs API v2.0:
  - **Database**: Search, Releases, Masters, Artists, and Labels.
  - **Marketplace**: Listings, Orders, Price Suggestions, and Fee calculations.
  - **Inventory**: CSV Batch operations (Export/Upload).
  - **User**: Identity, Profiles, Submissions, and Contributions.
  - **Collection**: Folder management and valuation.
  - **Wantlist**: Release tracking and personal notes.
  - **Lists**: Access to user-curated lists.
- Support for multiple authentication strategies:
  - Anonymous (unauthenticated).
  - Personal Access Token.
  - Consumer Key/Secret.
  - OAuth 1.0a (PLAINTEXT).
- Automatic local rate limiting based on authentication tier.
- Comprehensive `doc.go` with architectural overview.
- Detailed GoDoc comments for all exported types, methods, and struct fields.
- Idiomatic executable examples in `example_test.go` for core functionality.
- GitHub Actions CI/CD pipeline (Lint, Test with race detection, Security Scan).
- Dependabot configuration for automated dependency updates.
- `CONTRIBUTING.md` and professional `README.md`.
- Functional Options pattern for flexible client initialization.
- Context (`context.Context`) support for all API network calls.

### Changed

- Refactored internal request execution into a centralized `Client.do` method.
- Updated license badge to a reliable static version.

### Fixed

- Fixed invalid JSON tag for `MediaCondition` in `OrderItem` model.
- Resolved multiple linting issues (`errcheck`, `staticcheck`) identified during CI setup.
