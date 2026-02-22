# 🏥 ZARISH HEALTH - FHIR R5 Implementation

**Status**: Production Ready ✅  
**Version**: 0.4.0-alpha  
**Documentation**: [https://zs-health.github.io/zh-fhir-go](https://zs-health.github.io/zh-fhir-go)

## 🌟 What This Repository Does

This repository contains the core components for a **FHIR (Fast Healthcare Interoperability Resources) R5** implementation tailored for Bangladesh. In simple terms, FHIR is a global standard that helps different healthcare computer systems talk to each other. Think of it as a universal language for health data.

An **Implementation Guide (IG)** is like a detailed instruction manual that explains how to use FHIR for a specific purpose or region. This repository includes the Bangladesh Core FHIR IG, which provides guidelines for exchanging health information within Bangladesh.

## 📋 Table of Contents

1. [What This Repository Does](#-what-this-repository-does)
2. [How to View the Published IG](#-how-to-view-the-published-ig)
3. [How Changes Are Published](#-how-changes-are-published)
4. [For Non-Technical Users](#-for-non-technical-users)
5. [Quick Start](#-quick-start)
6. [Features](#-features)
7. [Documentation](#-documentation)
8. [Project Structure](#-project-structure)
9. [Build & Deploy](#-build--deploy)
10. [Resources & Support](#-resources--support)

## 🌐 How to View the Published IG

The complete Bangladesh Core FHIR Implementation Guide and related documentation are automatically published and available online:

*   **FHIR Implementation Guide**: [https://zs-health.github.io/zh-fhir-go/](https://zs-health.github.io/zh-fhir-go/)
*   **VitePress Documentation**: [https://zs-health.github.io/zh-fhir-go/](https://zs-health.github.io/zh-fhir-go/)

## 🔄 How Changes Are Published

This repository uses **GitHub Actions** to automatically build and publish the FHIR IG and documentation to **GitHub Pages** whenever changes are pushed to the `main` branch. This means that any updates made to the FHIR profiles or documentation will be reflected on the live website without manual intervention.

## 🧑‍💻 For Non-Technical Users

If you're not a developer, here's what you need to know:

*   **Viewing the Published Documentation**: You can access the latest version of the FHIR IG and project documentation at the links above. This is where you'll find all the information about how health data is structured and exchanged in Bangladesh.
*   **Reporting Issues**: If you find any issues with the documentation, the FHIR profiles, or have suggestions for improvement, please report them on our [GitHub Issues page](https://github.com/zs-health/zh-fhir-go/issues). Your feedback is valuable!
*   **What Happens When Code is Pushed**: When a developer makes changes and pushes them to the `main` branch, our automated system takes over. It builds the IG and documentation, and then publishes them to GitHub Pages. This process usually takes a few minutes.

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
