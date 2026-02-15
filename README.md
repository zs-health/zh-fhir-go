# 🏥 ZARISH HEALTH - FHIR R5 & HL7 v2 Go Library

**Status**: Integrated with zh-core ✅  
**Current Phase**: Phase 3 - Rohingya Refugee Response Integration  
**Last Updated**: February 14, 2026  
**Version**: 0.3.0-alpha

## 📋 Table of Contents

1. [Vision & Mission](#-vision--mission)
2. [Project Identity](#-project-identity)
3. [Bangladesh & Rohingya Localization](#-bangladesh--rohingya-localization)
4. [Features](#-features)
5. [Project Structure](#-project-structure)
6. [Quick Start Guide](#-quick-start-guide)
7. [Build & Publish Guide](#-build--publish-guide)
8. [Resources & Support](#-resources--support)

## 🎯 Vision & Mission

### Vision
To provide the foundational data structures and interoperability tools for the **ZARISH HEALTH** ecosystem, ensuring seamless data exchange across humanitarian health settings, specifically for **Bangladesh** and the **Rohingya Refugee Response**.

### Mission
* **Standardize**: Full FHIR R5 and HL7 v2 specifications.
* **Localize**: Specialized support for NID, BRN, FCN, Progress ID, and detailed Camp/Shelter tracking.
* **Empower**: Enable clinicians in humanitarian settings with standardized data tools.

## 🇧🇩 Bangladesh & Rohingya Localization

This library includes specialized support for the humanitarian context in Bangladesh:

*   **Rohingya Identifiers**: Support for **FCN** (Family Counting Number), **Progress ID**, and **MRN**.
*   **Shelter Tracking**: Detailed extensions for **Camp**, **Block**, **Sub-block**, and **Shelter Number**.
*   **Bangladesh Identifiers**: Extensions for National ID (NID) and Birth Registration Number (BRN).
*   **Terminology**: Built-in support for **ICD-11** coding systems.

## 🚀 Quick Start Guide

### Rohingya Refugee Patient Example

```go
package main

import (
    "github.com/zs-health/zh-fhir-go/fhir/r5/profiles/bd"
)

func main() {
    patient := bd.NewRohingyaPatient()
    
    // Add Specialized Identifiers
    patient.AddRohingyaIdentifiers("FCN-123", "PROG-456", "MRN-789")
    
    // Set Shelter Location
    patient.SetShelterLocation("Camp 1E", "Block A", "Sub-1", "S-101")
}
```

## 🛠 Build & Publish Guide

For users who are not coders, follow these simple steps to keep your library updated on GitHub.

### 1. Update Your Local Files
If you have made changes or want to sync with the latest version:
```bash
cd ~/Desktop/zh-fhir-go
git pull origin main
```

### 2. Verify Your Changes
To ensure everything is working correctly, you can run the built-in tests:
```bash
go test ./...
```

### 3. Build the CLI Tool
To create a usable program (`zh-fhir`) from the code:
```bash
go build -o zh-fhir ./cmd/zh-fhir
```

### 4. Publish to GitHub
To save your latest changes to your GitHub organization:
```bash
git add .
git commit -m "Update: Added new Rohingya refugee extensions"
git push origin main
```

## 📁 Project Structure

```text
.
├── cmd/
│   └── zh-fhir/          # CLI tool
├── fhir/
│   ├── r5/
│   │   ├── profiles/bd/  # Localized profiles (BDPatient, RohingyaPatient)
│   │   ├── valuesets/bd/ # Localized ValueSets (Divisions, Camps)
│   │   └── terminology/  # ICD-11 helpers
├── hl7/                  # HL7 v2.x support
└── README.md
```

---
© 2026 ZARISH HEALTH. All rights reserved.
