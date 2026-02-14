# 🏥 ZARISH HEALTH - FHIR R5 & HL7 v2 Go Library

**Status**: Integrated with zh-core ✅  
**Current Phase**: Phase 1 - Core Resource Implementation  
**Last Updated**: February 14, 2026  
**Version**: 0.1.0-alpha

## 📋 Table of Contents

1. [Vision & Mission](#-vision--mission)
2. [Project Identity](#-project-identity)
3. [Features](#-features)
4. [Technology Stack](#-technology-stack)
5. [Project Structure](#-project-structure)
6. [Quick Start Guide](#-quick-start-guide)
7. [Development Standards](#-development-standards)
8. [Resources & Support](#-resources--support)

## 🎯 Vision & Mission

### Vision
To provide the foundational data structures and interoperability tools for the **ZARISH HEALTH** ecosystem, ensuring seamless data exchange across humanitarian health settings using global standards.

### Mission
* **Standardize**: Implement full FHIR R5 and HL7 v2 specifications for Go-based microservices.
* **Simplify**: Provide type-safe, easy-to-use Go structs for complex healthcare resources.
* **Integrate**: Enable `zh-core` microservices to communicate using industry-standard protocols.

## 🆔 Project Identity

* **Organization**: [zs-health](https://github.com/zs-health)
* **Project Name**: zh-fhir-go
* **Parent Platform**: [ZARISH HEALTH](https://github.com/zs-health/zh-core)
* **Standards Compliance**: FHIR R5, HL7 v2.x, ICD-11

## ✨ Features

### FHIR R5 Support
* **158 FHIR Resources**: All R5 resources generated from official HL7 StructureDefinitions.
* **Type-Safe API**: Compile-time safety with Go structs matching FHIR specification exactly.
* **Validation**: Built-in validation for cardinality, mandatory elements, and choice types.
* **SMART on FHIR**: Support for OAuth2-based authorization flows.

### HL7 v2 Support
* **Message Parsing**: Parse ADT, ORM, ORU, and other standard v2.x messages.
* **Fluent API**: Programmatically build HL7 messages with ease.
* **MLLP Protocol**: Minimal Lower Layer Protocol for secure message transmission.

## 💻 Technology Stack

* **Language**: Go 1.25+
* **Standards**: HL7 FHIR R5, HL7 v2.x
* **Validation**: [go-playground/validator](https://github.com/go-playground/validator)
* **Testing**: [stretchr/testify](https://github.com/stretchr/testify)

## 📁 Project Structure

```text
.
├── cmd/
│   └── zh-fhir/          # CLI tool for FHIR/HL7 operations
├── fhir/
│   ├── r5/               # FHIR R5 resource definitions
│   │   ├── resources/    # Individual FHIR resources (Patient, etc.)
│   │   └── types/        # Common FHIR data types
│   ├── primitives/       # FHIR primitive types (Date, Instant, etc.)
│   └── validation/       # FHIR resource validation logic
├── hl7/                  # HL7 v2.x implementation
├── go.mod                # Go module definition
└── README.md             # Project documentation
```

## 🚀 Quick Start Guide

### Installation

```bash
go get github.com/zs-health/zh-fhir-go
```

### Usage Example (FHIR R5)

```go
package main

import (
    "fmt"
    "github.com/zs-health/zh-fhir-go/fhir/r5/resources"
)

func main() {
    patient := &resources.Patient{
        Active: boolPtr(true),
        Name: []resources.HumanName{
            {
                Family: stringPtr("Doe"),
                Given:  []string{"John"},
            },
        },
    }
    fmt.Printf("Created patient: %s\n", *patient.Name[0].Family)
}

func stringPtr(s string) *string { return &s }
func boolPtr(b bool) *bool       { return &b }
```

## 📏 Development Standards

This project follows the **ZARISH HEALTH** development standards:
* **Naming**: Use `zh-` prefix for all ecosystem components.
* **Code Style**: Follow standard Go formatting (`gofmt`) and linting.
* **Documentation**: Maintain clear, concise READMEs and inline comments.

## 📞 Resources & Support

* **Organization**: [Zarish Sphere](https://github.com/zs-health)
* **Main Platform**: [zh-core](https://github.com/zs-health/zh-core)
* **Website**: [health.zarishsphere.com](https://health.zarishsphere.com)

---
© 2026 ZARISH HEALTH. All rights reserved.
