# 🏥 ZARISH HEALTH - FHIR R5 Implementation

**Status**: Production Ready ✅  
**Version**: 0.4.0-alpha  
**Documentation**: [https://zs-health.github.io/zh-fhir-go](https://zs-health.github.io/zh-fhir-go)

## 📋 Table of Contents

1. [Quick Start](#-quick-start)
2. [Features](#-features)
3. [Documentation](#-documentation)
4. [Project Structure](#-project-structure)
5. [Build & Deploy](#-build--deploy)
6. [Resources & Support](#-resources--support)

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/zs-health/zh-fhir-go.git
cd zh-fhir-go

# Build the CLI
go build -o zh-fhir ./cmd/zh-fhir

# Start FHIR server
./zh-fhir --server --port 8080

# Start terminology server only
./zh-fhir --term-server --port 8080
```

## ✨ Features

| Feature | Description |
|---------|-------------|
| **FHIR R5** | Complete implementation of all FHIR R5 resources |
| **Type-Safe** | Custom Go types with validation for primitives |
| **Bangladesh Profiles** | NID, BRN, UHID identifiers |
| **Rohingya Support** | FCN, Progress ID, Camp locations |
| **Terminology Server** | ICD-11 and Bangladesh divisions |
| **REST API** | Full CRUD operations |

## 📖 Documentation

### Interactive Documentation

Visit our **modern documentation site**: [https://zs-health.github.io/zh-fhir-go](https://zs-health.github.io/zh-fhir-go)

### Quick Links

- [Getting Started](/guide/introduction) - Introduction and overview
- [Installation](/guide/installation) - Step-by-step setup
- [Quick Start](/guide/quickstart) - Get running in 5 minutes
- [API Reference](/api/overview) - REST API documentation
- [FHIR Resources](/fhir/overview) - Resource documentation
- [Terminology](/terminology/overview) - ICD-11 and codes

## 📁 Project Structure

```
zh-fhir-go/
├── cmd/zh-fhir/           # CLI application
│   ├── main.go            # Entry point
│   └── terminology.go     # Terminology server
├── fhir/                  # FHIR library
│   ├── r5/                # FHIR R5 resources
│   ├── r4/                # FHIR R4 resources
│   ├── primitives/        # Type-safe primitives
│   └── validation/        # Validation utilities
├── internal/
│   ├── server/            # FHIR REST server
│   └── ig/                # IG loader
├── docs/                  # Documentation (VitePress)
│   ├── guide/             # Getting started
│   ├── api/               # API reference
│   ├── fhir/              # FHIR resources
│   └── terminology/       # Terminology docs
└── .github/workflows/     # CI/CD
```

## 🛠 Build & Deploy

### Local Development

```bash
# Install dependencies
go mod download

# Run tests
go test -v ./...

# Build
go build -v ./...

# Run CLI
./zh-fhir --help
```

### Docker

```bash
# Build Docker image
docker build -t zh-fhir .

# Run FHIR server
docker run -p 8080:8080 zh-fhir --server --port 8080
```

### Documentation

```bash
# Install Node.js dependencies
npm install

# Development server
npm run docs:dev

# Build for production
npm run docs:build
```

## 🏗 FHIR Infrastructure

### 1. Terminology Server
- **Endpoint**: `http://localhost:8080/fhir/ValueSet/$expand`
- **Supports**: ICD-11, Bangladesh divisions
- **Start**: `./zh-fhir --term-server --port 8080`

### 2. FHIR REST Server
- **Full CRUD**: Create, Read, Update, Delete
- **Search**: Basic search capabilities
- **Start**: `./zh-fhir --server --port 8080`

### 3. Bangladesh Profiles
- **BDPatient**: NID, BRN, UHID identifiers
- **BDAddress**: Administrative divisions
- **Rohingya**: FCN, Progress ID, Camp locations

## 📞 Resources & Support

- **Documentation**: [https://zs-health.github.io/zh-fhir-go](https://zs-health.github.io/zh-fhir-go)
- **GitHub Issues**: [https://github.com/zs-health/zh-fhir-go/issues](https://github.com/zs-health/zh-fhir-go/issues)
- **FHIR Spec**: [https://hl7.org/fhir/](https://hl7.org/fhir/)

---

© 2026 ZARISH HEALTH. All rights reserved.
