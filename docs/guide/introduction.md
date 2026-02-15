# Introduction

Welcome to **zh-fhir-go**, a comprehensive FHIR R5 implementation designed for Bangladesh healthcare systems and humanitarian response.

## What is FHIR?

FHIR (Fast Healthcare Interoperability Resources) is the next generation of HL7 standards that enables easy exchange of healthcare information between different systems. It combines the best features of HL7 v2, v3, and CDA while leveraging modern web technologies like JSON and REST.

## Why zh-fhir-go?

### Key Features

| Feature | Description |
|---------|-------------|
| **FHIR R5 Compliant** | Complete implementation of all FHIR R5 resources and types |
| **Bangladesh Profiles** | Localized profiles for Bangladesh National Health Data Standards |
| **Rohingya Support** | Extensions for refugee camp operations and humanitarian response |
| **Type-Safe** | Custom Go types with validation for all FHIR primitives |
| **Zero Dependencies** | No external dependencies for the core FHIR library |
| **CLI Tools** | Ready-to-use FHIR server and terminology service |

## Project Status

- ✅ FHIR R5 Resources: Complete
- ✅ FHIR R5 Types: Complete
- ✅ Primitive Types: Complete with validation
- ✅ Bangladesh Profiles: Active development
- ✅ Terminology Server: Implemented
- ✅ FHIR REST Server: Implemented

## Use Cases

### 1. Healthcare Interoperability

Connect different healthcare systems within Bangladesh using standardized FHIR resources.

### 2. Refugee Health Records

Support humanitarian organizations working with Rohingya refugees by using appropriate identifiers and location data.

### 3. National Health Dashboard

Aggregate health data from multiple sources using the FHIR server's RESTful API.

### 4. Terminology Services

Validate and expand healthcare codes (ICD-11, local codes) using the built-in terminology server.

## Architecture

```
zh-fhir-go/
├── cmd/zh-fhir/           # CLI application
│   ├── main.go            # Entry point
│   └── terminology.go     # Terminology server
├── fhir/                  # FHIR library
│   ├── r5/                # FHIR R5 resources
│   ├── primitives/        # Type-safe primitive types
│   └── validation/        # Validation utilities
├── internal/
│   ├── server/            # FHIR REST server
│   └── ig/                # Implementation guide loader
└── docs/                  # Documentation (VitePress)
```

## Getting Help

- **GitHub Issues**: https://github.com/zs-health/zh-fhir-go/issues
- **Documentation**: https://zs-health.github.io/zh-fhir-go
- **FHIR Specification**: https://hl7.org/fhir/

## License

This project is licensed under the MIT License - see the LICENSE file for details.
