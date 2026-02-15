# 🏥 ZARISH HEALTH - FHIR R5 & HL7 v2 Infrastructure

**Status**: Integrated with zh-core ✅  
**Current Phase**: Phase 4 - Full FHIR Infrastructure Deployment  
**Last Updated**: February 14, 2026  
**Version**: 0.4.0-alpha

## 📋 Table of Contents

1. [Vision & Mission](#-vision--mission)
2. [Project Identity](#-project-identity)
3. [FHIR Infrastructure](#-fhir-infrastructure)
4. [Bangladesh & Rohingya Localization](#-bangladesh--rohingya-localization)
5. [Build & Publish Guide](#-build--publish-guide)
6. [Resources & Support](#-resources--support)

## 🎯 Vision & Mission

### Vision
To provide a complete, localized FHIR infrastructure for the **ZARISH HEALTH** ecosystem, supporting the **Bangladesh National Health Data Standards** and the **Rohingya Refugee Response**.

## 🏗 FHIR Infrastructure

This repository now includes a complete set of tools for FHIR implementation:

### 1. Terminology Server
A lightweight, built-in terminology server for expanding ValueSets and validating codes (ICD-11, local geography).
*   **Run**: `./zh-fhir -term-server -port 8080`
*   **Endpoint**: `http://localhost:8080/fhir/ValueSet/$expand`

### 2. DGHS Standard Profiles
Localized profiles based on the **Bangladesh National FHIR IG**:
*   **BDPatient**: Supports NID, BRN, UHID, and bilingual names.
*   **BDAddress**: Standardized Bangladesh administrative levels.

### 3. Rohingya Refugee Support
Specialized extensions for humanitarian response:
*   **Identifiers**: FCN, Progress ID, MRN.
*   **Location**: Camp, Block, Sub-block, Shelter Number.

## 🛠 Build & Publish Guide

### For Non-Coders:
1.  **Sync Local Machine**:
    ```bash
    git pull origin main
    ```
2.  **Run Everything**:
    ```bash
    ./scripts/run_fhir.sh
    ```
    *This will build the tool and start the terminology server.*

### GitHub Actions:
Every push to this repository is automatically built and tested via **GitHub Actions** to ensure quality and compliance.

## 📁 Project Structure

```text
.
├── .github/workflows/    # Automated build & test
├── cmd/zh-fhir/          # CLI & Terminology Server
├── fhir/r5/              # FHIR R5 Resources & Profiles
│   ├── profiles/bd/      # Localized BD & Rohingya Profiles
│   └── terminology/      # ICD-11 & Terminology Logic
├── scripts/              # Infrastructure scripts
└── README.md
```

---
© 2026 ZARISH HEALTH. All rights reserved.
