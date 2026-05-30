# Contributing to go-discogs-client

First off, thank you for considering contributing to `go-discogs-client`! It's people like you that make the open-source community such an amazing place to learn, inspire, and create.

## 📜 Code of Conduct

By participating in this project, you agree to abide by the same standards of professional and respectful communication expected in any professional software environment.

## 🚀 How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues to see if the problem has already been reported. When creating a bug report, please include as many details as possible:

* **Use a clear and descriptive title.**
* **Describe the exact steps which reproduce the problem.**
* **Explain which behavior you expected to see and why.**
* **Include code snippets or a small reproduction script.**

### Suggesting Enhancements

If you have an idea for a new feature or an improvement to an existing one, please open an issue first to discuss it with the maintainers. This ensures that the proposed change aligns with the project's goals and architecture.

### Pull Requests

1. **Fork the repository** and create your branch from `main`.
2. **Ensure your code follows Go standards**:
    * Run `go fmt ./...`
    * Run `go mod tidy`
3. **Documentation is mandatory**:
    * All exported types, methods, and fields must have GoDoc-compliant comments.
    * If you add a new feature, include an executable example in `example_test.go`.
4. **Testing is mandatory**:
    * Include unit tests for any new logic.
    * Ensure all tests pass by running `go test -v -race ./...`.
5. **Passing CI**: Your PR must pass all GitHub Actions checks (Lint, Test, and Security).
6. **Commit Messages**: We follow [Conventional Commits](https://www.conventionalcommits.org/) (e.g., `feat: add order management`, `fix: handle nil response in search`).

## 🛠️ Development Setup

To set up your local development environment:

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/go-discogs-client.git
cd go-discogs-client

# Install dependencies
go mod download

# Run the full validation suite (same as CI)
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
go run golang.org/x/vuln/cmd/govulncheck@latest ./...
```

## ⚖️ License

By contributing, you agree that your contributions will be licensed under the project's [MIT License](LICENSE).
