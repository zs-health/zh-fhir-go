# Quick Start

This guide will help you get started with zh-fhir-go in just a few minutes.

## Start the FHIR Server

The fastest way to get started is to run the built-in FHIR server:

```bash
# From the project root
./zh-fhir --server --port 8080
```

You should see output like:

```
2026/02/15 12:00:00 Loading IG data from ./BD-Core-FHIR-IG...
2026/02/15 12:00:00 Loaded 10 CodeSystems and 25 ValueSets
2026/02/15 12:00:00 FHIR Server starting on :8080...
```

## Test the Server

### 1. Check Server Metadata

```bash
curl http://localhost:8080/fhir
```

### 2. Create a Patient

```bash
curl -X POST http://localhost:8080/fhir/Patient \
  -H "Content-Type: application/fhir+json" \
  -d '{
    "resourceType": "Patient",
    "active": true,
    "name": [{
      "family": "Chowdhury",
      "given": ["Rahima"]
    }],
    "gender": "female",
    "birthDate": "1990-05-15"
  }'
```

### 3. Read the Patient

```bash
# Replace {id} with the ID from the previous response
curl http://localhost:8080/fhir/Patient/{id}
```

### 4. Search for Patients

```bash
curl http://localhost:8080/fhir/Patient
```

### 5. Update a Patient

```bash
curl -X PUT http://localhost:8080/fhir/Patient/{id} \
  -H "Content-Type: application/fhir+json" \
  -d '{
    "resourceType": "Patient",
    "id": "{id}",
    "active": true,
    "name": [{
      "family": "Chowdhury",
      "given": ["Rahima", "Begum"]
    }],
    "gender": "female",
    "birthDate": "1990-05-15"
  }'
```

### 6. Delete a Patient

```bash
curl -X DELETE http://localhost:8080/fhir/Patient/{id}
```

## Use the Terminology Server

Start the standalone terminology server:

```bash
./zh-fhir --term-server --port 8081
```

### Expand a ValueSet

```bash
# Get all ICD-11 concepts
curl "http://localhost:8081/fhir/ValueSet/\$expand?url=http://id.who.int/icd/release/11/mms"

# Filter results
curl "http://localhost:8081/fhir/ValueSet/\$expand?url=http://id.who.int/icd/release/11/mms&filter=hyper"
```

## Use the FHIR Library

### Create a Patient Resource

```go
package main

import (
    "encoding/json"
    "fmt"
    
    "github.com/zs-health/zh-fhir-go/fhir/r5"
    "github.com/zs-health/zh-fhir-go/fhir/primitives"
)

func main() {
    active := true
    birthDate := primitives.MustDate("1990-05-15")
    
    patient := r5.Patient{
        Active: &active,
        Name: []r5.HumanName{
            {
                Use:    r5.HumanNameUseOfficial,
                Family: "Chowdhury",
                Given:  []string{"Rahima"},
            },
        },
        Gender:    r5.PatientGenderFemale,
        BirthDate: &birthDate,
    }
    
    // Marshal to JSON
    data, _ := json.MarshalIndent(patient, "", "  ")
    fmt.Println(string(data))
}
```

## Next Steps

- [API Reference](/api/overview) - Full API documentation
- [FHIR Resources](/fhir/overview) - Learn about supported resources
- [Terminology Service](/terminology/overview) - Explore terminology operations
