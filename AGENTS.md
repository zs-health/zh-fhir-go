# Agent Guidelines for zh-fhir-go

This file provides guidance for AI coding agents operating in this repository.

## Project Overview

zh-fhir-go is a Go-based FHIR (Fast Healthcare Interoperability Resources) implementation with DICOM support. The project includes:
- FHIR resource types (R4, R5)
- CLI tool for FHIR operations
- Server components for terminology services

## Build Commands

### Standard Build
```bash
# Build all packages
go build -v ./...

# Build specific binary
go build -o zh-fhir ./cmd/zh-fhir
```

### Testing
```bash
# Run all tests
go test -v ./...

# Run tests for specific package
go test -v ./fhir/...

# Run a single test
go test -v -run TestFunctionName ./path/to/package

# Run tests with coverage
go test -v -cover ./...

# Run tests in watch mode (requires gotestsum)
gotestsum -- -watch
```

### Linting
```bash
# Run go vet
go vet ./...

# Run golangci-lint (project linting config)
golangci-lint run ./...

# Run golangci-lint on specific file
golangci-lint run path/to/file.go

# Run pre-commit hooks manually
pre-commit run --all-files
```

### Formatting
```bash
# Format all Go files (use before commits)
gofmt -w .

# Format and simplify
gofmt -w -s .

# Run go mod tidy
go mod tidy
```

### Other Useful Commands
```bash
# Download dependencies
go mod download

# Show module graph
go mod graph

# List packages
go list ./...

# Run specific binary with args
./zh-fhir --version
```

## Code Style Guidelines

### General Principles
- Write clear, readable code over clever code
- Follow Go idioms and conventions
- Keep functions short and focused (< 40 lines when possible)
- Use meaningful variable and function names

### Formatting
- Use `gofmt` for all code (tabs, not spaces for indentation)
- Group imports: standard library, external packages, internal packages
- Blank line between import groups
- Maximum line length: 120 characters (soft limit)
- Always run `gofmt -w -s .` before committing

### Naming Conventions
- **Packages**: lowercase, short, no underscores (e.g., `fhir`, `server`)
- **Functions/Variables**: mixedCase or camelCase
- **Constants**: CamelCase for exported, mixedCase for unexported
- **Types**: CamelCase
- **Acronyms**: Use HTTP, not Http; URL, not Url
- **Booleans**: Use is/has/can prefixes: `isValid`, `hasData`, `canExecute`

### Error Handling
- Always handle errors explicitly (no `_` ignored errors unless intentional)
- Return meaningful error messages with context
- Use `fmt.Errorf` with `%w` for wrapped errors
- Check errors immediately after function calls
- Example:
  ```go
  if err != nil {
      return fmt.Errorf("failed to process resource: %w", err)
  }
  ```

### Imports
- Group and order imports:
  1. Standard library (`fmt`, `os`, `time`, etc.)
  2. External packages (`github.com/...`)
  3. Internal packages (`github.com/zs-health/zh-fhir-go/...`)
- Use import aliases only when necessary
- Run `go mod tidy` after adding imports

### Types
- Use concrete types over interfaces unless necessary
- Prefer `struct{}` over `interface{}` for marker types
- Use pointers (`*Type`) when mutation is needed or for nullable values
- Return slices rather than pointers to slices

### Comments
- Comment exported functions and types
- Use sentences with proper punctuation
- Example:
  ```go
  // Validate checks if the resource conforms to FHIR specification.
  func (r *Patient) Validate() error {
  ```

### Testing
- Write tests for all exported functions
- Use table-driven tests when testing multiple cases
- Test file naming: `filename_test.go`
- Use descriptive test names: `TestValidate_ValidPatient_NoError`
- Keep test helper functions in separate files when used across tests

### Logging
- Use structured logging with `github.com/charmbracelet/log`
- Include relevant context in log entries
- Use appropriate log levels: Debug, Info, Warn, Error, Fatal

## Project Structure

```
.
├── cmd/zh-fhir/           # CLI application
│   └── internal/
│       ├── cli/           # CLI commands
│       ├── config/       # Configuration
│       └── build/        # Build info
├── fhir/                 # FHIR resource implementations
│   ├── primitives/      # Base types (Date, Time, etc.)
│   ├── r4/              # FHIR R4 resources
│   └── r5/              # FHIR R5 resources
├── internal/
│   ├── server/           # HTTP server
│   └── ig/              # Implementation guide loader
├── openspec/            # OpenSpec documentation (if used)
└── BD-Core-FHIR-IG/     # Submodule for IG definitions
```

## CI/CD

GitHub Actions workflow (`.github/workflows/ci.yml`):
- Runs on push to main and pull requests
- Uses Go 1.24
- Runs: `go build -v ./...` then `go test -v ./...`

## Dependencies

Key dependencies (see `go.mod`):
- `github.com/alecthomas/kong` - CLI framework
- `github.com/charmbracelet/log` - Logging
- `github.com/charmbracelet/lipgloss` - Terminal styling
- `github.com/go-playground/validator/v10` - Validation
- `github.com/google/uuid` - UUID generation
- `github.com/stretchr/testify` - Testing assertions

## Known Patterns

### FHIR Resource Validation
FHIR resources use `github.com/go-playground/validator` for validation. Add validation tags to struct fields.

### CLI Commands
CLI uses Kong framework. Commands are defined as structs with Kong tags:
```go
type MyCommand struct {
    Input string `arg:"" help:"Input file path"`
    Verbose bool `help:"Enable verbose output"`
}
```

### Configuration
Global config is in `cmd/zh-fhir/internal/config/config.go`. CLI merges defaults with flags and environment variables.

## Common Tasks

### Adding a New FHIR Resource
1. Create struct in appropriate `fhir/r4/resources/` or `fhir/r5/resources/`
2. Add validation tags
3. Add to resource registry if needed
4. Write tests

### Adding a New CLI Command
1. Add struct in appropriate file under `cmd/zh-fhir/internal/cli/`
2. Add command field to CLI struct
3. Implement `Run(*GlobalConfig) error` method

### Updating Dependencies
1. Run `go get -u package@version` or `go get -u` for all
2. Run `go mod tidy`
3. Verify build: `go build ./...`
4. Run tests: `go test ./...`

## Additional Resources

- FHIR Specification: https://hl7.org/fhir/
- Go Style Guide: https://golang.org/doc/style
- Project README: ./README.md
