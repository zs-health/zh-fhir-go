# API Overview

The zh-fhir-go project provides two main API components:

1. **FHIR REST Server** - Full RESTful API for FHIR resources
2. **Terminology Server** - ValueSet expansion and code lookup

## FHIR REST Server

The FHIR REST server implements the standard FHIR REST API with CRUD operations:

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/fhir` | Server capability statement (metadata) |
| `POST` | `/fhir/{resourceType}` | Create a new resource |
| `GET` | `/fhir/{resourceType}` | Search for resources |
| `GET` | `/fhir/{resourceType}/{id}` | Read a specific resource |
| `PUT` | `/fhir/{resourceType}/{id}` | Update a resource |
| `DELETE` | `/fhir/{resourceType}/{id}` | Delete a resource |

## Terminology Service

The terminology service provides FHIR terminology operations:

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/fhir/ValueSet/$expand` | Expand a ValueSet |

## Quick Links

- [FHIR Server Details](/api/server) - Server configuration and options
- [REST Endpoints](/api/endpoints) - Complete endpoint reference
- [Terminology Overview](/terminology/overview) - Terminology service details
