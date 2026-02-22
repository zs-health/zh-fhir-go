---
layout: home

hero:
  name: "zh-fhir-go"
  text: "FHIR R5 Implementation for Bangladesh Healthcare"
  tagline: A comprehensive Go library and CLI tool for working with HL7 FHIR R5, with Bangladesh and Rohingya refugee localization support.
  actions:
    - theme: brand
      text: Get Started
      link: /guide/introduction
    - theme: alt
      text: View FHIR IG
      link: https://zs-health.github.io/zh-fhir-go/
    - theme: alt
      text: View on GitHub
      link: https://github.com/zs-health/zh-fhir-go
  image:
    src: /logo.svg
    alt: zh-fhir-go

features:
  - title: FHIR R5 Compliant
    details: Complete implementation of FHIR R5 resources, types, and bundles. Generated from official HL7 StructureDefinitions.
  - title: Bangladesh Localization
    details: Native support for Bangladesh healthcare identifiers (NID, BRN, UHID) and administrative divisions.
  - title: Rohingya Support
    details: Specialized extensions for humanitarian response including FCN, Progress ID, and Camp location data.
  - title: Terminology Services
    details: Built-in terminology server with ICD-11 support and ValueSet/$expand operation.
  - title: Type-Safe Primitives
    details: Custom Date, DateTime, Time, and Instant types with validation and partial precision support.
  - title: RESTful FHIR Server
    details: Ready-to-use FHIR server with CRUD operations, search, and terminology integration.
---

<div style="text-align: center; padding: 2rem 0;">

## Quick Start

```bash
# Clone the repository
git clone https://github.com/zs-health/zh-fhir-go.git

# Build the CLI
cd zh-fhir-go
go build -o zh-fhir ./cmd/zh-fhir

# Start the FHIR server
./zh-fhir --server --port 8080
```

</div>
