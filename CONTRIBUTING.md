# Contributing to zh-fhir-go

Thank you for your interest in contributing to *zh-fhir-go*! Whether you're reporting
an issue, requesting a new feature, or submitting a pull request, we appreciate
your time and effort.

## Getting Started

1. **Fork** the repository and clone your fork locally.
2. Create a new branch for your work (e.g. `feature/my-change` or `bugfix/issue-123`).
3. Install dependencies:
   ```sh
   go mod download   # for Go code
   npm install       # for documentation site (in `docs/`)
   ```
4. Make your changes and add tests where appropriate. Use `go test ./...` to run
the Go test suite.
5. Update documentation in `docs/` or README as needed.
6. Commit with clear messages and push your branch to GitHub.
7. Open a pull request against the `main` branch of this repository.

## Code Style & Guidelines

* Follow Go idioms (`gofmt`/`go vet`).
* Keep API changes backward compatible when possible.
* Add or update unit tests for any new functionality or bug fixes.
* Run the linters included in CI (the `ci.yml` job) before submitting.

## Issue Reporting

* Use the [GitHub Issues](https://github.com/zs-health/zh-fhir-go/issues) page.
* Provide a clear description of the problem or enhancement, steps to reproduce,
  and any relevant logs or screenshots.
* Label your issue appropriately if you can (bug, feature request, etc.).

## Security

See [SECURITY.md](SECURITY.md) for details on responsibly disclosing vulnerabilities.

## Pull Request Checklist

* [ ] Code compiles and tests pass (`go test ./...`).
* [ ] New features are documented.
* [ ] Tests cover new or changed behavior.
* [ ] PR description explains the motivation and impact of your changes.

Thanks again for helping improve `zh-fhir-go`!
