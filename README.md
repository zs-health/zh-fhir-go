# zh-fhir-go

A comprehensive Go library for healthcare interoperability standards focusing on **FHIR R5** and **HL7 v2** resources.

[![Go Version](https://img.shields.io/badge/Go-1.25.4+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Overview

`zh-fhir-go` provides robust, production-ready implementations of healthcare standards with a focus on type safety, performance, and developer experience. Built for clinical systems integration and healthcare data exchange.

## Features

### FHIR R5 Support

Complete implementation of HL7 FHIR Release 5 specification:

- **158 FHIR Resources** - All R5 resources generated from official HL7 StructureDefinitions
- **Type-Safe API** - Compile-time safety with Go structs matching FHIR specification exactly
- **Comprehensive Validation** - Built-in validation for cardinality, mandatory elements, and choice types.
- **SMART on FHIR** - OAuth2-based authorization framework for FHIR applications.

### HL7 v2 Support

Implementation of HL7 Version 2.x messaging standard:

- **Message Parsing** - Parse HL7 v2.x messages (ADT, ORM, ORU, etc.)
- **Message Generation** - Build HL7 messages programmatically with a fluent API.
- **MLLP Protocol** - Minimal Lower Layer Protocol client/server for message transmission.

## Installation

```bash
go get github.com/zs-health/zh-fhir-go
```

## Quick Start

### Working with FHIR R5

```go
package main

import (
"fmt"
"github.com/zs-health/zh-fhir-go/fhir/r5"
)

func main() {
patient := &r5.Patient{
Name: []r5.HumanName{
{
Family: "Doe",
Given:  []string{"John"},
},
},
}
fmt.Printf("Created patient: %s\n", patient.Name[0].Family)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
