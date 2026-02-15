# Installation

## Prerequisites

- **Go 1.24 or later**: [Download Go](https://golang.org/dl/)
- **Git**: For cloning the repository

## Install from Source

### 1. Clone the Repository

```bash
git clone https://github.com/zs-health/zh-fhir-go.git
cd zh-fhir-go
```

### 2. Download Dependencies

```bash
go mod download
```

### 3. Build the CLI

```bash
# Build the binary
go build -o zh-fhir ./cmd/zh-fhir

# Or install to $GOPATH/bin
go install ./cmd/zh-fhir
```

### 4. Verify Installation

```bash
./zh-fhir --version
```

## Use as a Library

You can also use the FHIR library directly in your Go projects:

```bash
go get github.com/zs-health/zh-fhir-go/fhir
```

## Docker Installation

### Build Docker Image

```bash
docker build -t zh-fhir .
```

### Run the Server

```bash
# Run FHIR server
docker run -p 8080:8080 zh-fhir --server --port 8080

# Run terminology server
docker run -p 8080:8080 zh-fhir --term-server --port 8080
```

## Development Setup

### Clone with Submodules

This project includes the Bangladesh FHIR Implementation Guide as a submodule:

```bash
git clone --recurse-submodules https://github.com/zs-health/zh-fhir-go.git
```

If you already cloned without submodules:

```bash
git submodule update --init --recursive
```

### Run Tests

```bash
# Run all tests
go test -v ./...

# Run specific package tests
go test -v ./fhir/...
go test -v ./fhir/primitives/...
```

### Code Generation

To regenerate FHIR resources from StructureDefinitions:

```bash
# Download FHIR schemas first
# Then run the code generator
./fhir/bin/fhir-gen -version r5 -input fhir_schemas/r5/profiles-resources.json -output fhir/r5/resources
```

## Next Steps

- [Quick Start Guide](/guide/quickstart) - Get up and running
- [API Reference](/api/overview) - Explore the server endpoints
- [FHIR Resources](/fhir/overview) - Learn about supported resources
