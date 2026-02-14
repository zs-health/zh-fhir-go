# 🏥 ZARISH HEALTH - FHIR R5 & HL7 v2 Go Library

**Status**: Integrated with zh-core ✅  
**Current Phase**: Phase 2 - Bangladesh Localization & ICD-11 Integration  
**Last Updated**: February 14, 2026  
**Version**: 0.2.0-alpha

## 📋 Table of Contents

1. [Vision & Mission](#-vision--mission)
2. [Project Identity](#-project-identity)
3. [Bangladesh Localization](#-bangladesh-localization)
4. [Features](#-features)
5. [Technology Stack](#-technology-stack)
6. [Project Structure](#-project-structure)
7. [Quick Start Guide](#-quick-start-guide)
8. [Resources & Support](#-resources--support)

## 🎯 Vision & Mission

### Vision
To provide the foundational data structures and interoperability tools for the **ZARISH HEALTH** ecosystem, ensuring seamless data exchange across humanitarian health settings using global standards, specifically tailored for **Bangladesh**.

### Mission
* **Standardize**: Implement full FHIR R5 and HL7 v2 specifications for Go-based microservices.
* **Localize**: Incorporate Bangladesh-specific identifiers (NID, BRN), ValueSets (Divisions, Districts), and ICD-11 terminology.
* **Simplify**: Provide type-safe, easy-to-use Go structs for complex healthcare resources.

## 🆔 Project Identity

* **Organization**: [zs-health](https://github.com/zs-health)
* **Project Name**: zh-fhir-go
* **Parent Platform**: [ZARISH HEALTH](https://github.com/zs-health/zh-core)
* **Standards Compliance**: FHIR R5, HL7 v2.x, ICD-11, LOINC

## 🇧🇩 Bangladesh Localization

This library includes specialized support for the Bangladesh healthcare context:

*   **Identifiers**: Extensions for National ID (NID) and Birth Registration Number (BRN).
*   **Terminology**: Built-in support for **ICD-11** (WHO) coding systems.
*   **ValueSets**: Localized lists for Bangladesh Divisions, Districts, and Upazilas.
*   **Profiles**: Specialized FHIR profiles (e.g., `BDPatient`) to ensure data consistency across local health facilities.

## ✨ Features

### FHIR R5 Support
* **158 FHIR Resources**: All R5 resources generated from official HL7 StructureDefinitions.
* **Type-Safe API**: Compile-time safety with Go structs matching FHIR specification exactly.
* **Validation**: Built-in validation for cardinality, mandatory elements, and choice types.

### HL7 v2 Support
* **Message Parsing**: Parse ADT, ORM, ORU, and other standard v2.x messages.
* **Fluent API**: Programmatically build HL7 messages with ease.

## 📁 Project Structure

```text
.
├── cmd/
│   └── zh-fhir/          # CLI tool for FHIR/HL7 operations
├── fhir/
│   ├── r5/               # FHIR R5 resource definitions
│   │   ├── profiles/bd/  # Bangladesh-specific FHIR profiles
│   │   ├── valuesets/bd/ # Bangladesh ValueSets (Divisions, etc.)
│   │   └── terminology/  # ICD-11 and other terminology helpers
│   ├── primitives/       # FHIR primitive types
│   └── validation/       # FHIR resource validation logic
├── hl7/                  # HL7 v2.x implementation
├── go.mod                # Go module definition
└── README.md             # Project documentation
```

## 🚀 Quick Start Guide

### Usage with Bangladesh Localization

```go
package main

import (
    "github.com/zs-health/zh-fhir-go/fhir/r5/profiles/bd"
    "github.com/zs-health/zh-fhir-go/fhir/r5/terminology/icd11"
)

func main() {
    // Create a Bangladesh Patient
    patient := bd.NewBDPatient()
    patient.AddNID("1990123456789")
    
    // Create an ICD-11 Diagnosis
    diagnosis := icd11.NewCodeableConcept("BA00", "Essential hypertension")
}
```

## 📏 Development Standards

This project follows the **ZARISH HEALTH** development standards:
* **Naming**: Use `zh-` prefix for all ecosystem components.
* **Localization**: Always prefer localized profiles (`bd.`) for patient demographics in Bangladesh.

## 📞 Resources & Support

* **Organization**: [Zarish Sphere](https://github.com/zs-health)
* **Main Platform**: [zh-core](https://github.com/zs-health/zh-core)
* **Website**: [health.zarishsphere.com](https://health.zarishsphere.com)

---
© 2026 ZARISH HEALTH. All rights reserved.
