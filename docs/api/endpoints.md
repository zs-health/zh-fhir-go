# REST Endpoints

Complete reference for all FHIR REST API endpoints.

## Base URL

```
http://localhost:8080/fhir
```

## Resource Operations

### Create Resource

Create a new FHIR resource.

**Request**

```http
POST /fhir/{resourceType}
Content-Type: application/fhir+json

{
  "resourceType": "Patient",
  "active": true,
  "name": [{
    "family": "Chowdhury",
    "given": ["Rahima"]
  }]
}
```

**Response**

```http
HTTP/1.1 201 Created
Content-Type: application/fhir+json

{
  "resourceType": "Patient",
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "active": true,
  "name": [{
    "family": "Chowdhury",
    "given": ["Rahima"]
  }]
}
```

---

### Read Resource

Read a specific resource by ID.

**Request**

```http
GET /fhir/{resourceType}/{id}
```

**Response**

```http
HTTP/1.1 200 OK
Content-Type: application/fhir+json

{
  "resourceType": "Patient",
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "active": true,
  ...
}
```

If the resource is not found:

```http
HTTP/1.1 404 Not Found
```

---

### Update Resource

Update an existing resource.

**Request**

```http
PUT /fhir/{resourceType}/{id}
Content-Type: application/fhir+json

{
  "resourceType": "Patient",
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "active": true,
  "name": [{
    "family": "Chowdhury",
    "given": ["Rahima", "Begum"]
  }]
}
```

---

### Delete Resource

Delete a resource.

**Request**

```http
DELETE /fhir/{resourceType}/{id}
```

**Response**

```http
HTTP/1.1 204 No Content
```

---

### Search Resources

Search for resources of a specific type.

**Request**

```http
GET /fhir/{resourceType}
```

**Response**

```http
HTTP/1.1 200 OK
Content-Type: application/fhir+json

{
  "resourceType": "Bundle",
  "type": "searchset",
  "total": 2,
  "entry": [
    {
      "resource": { ... }
    },
    {
      "resource": { ... }
    }
  ]
}
```

---

## Terminology Endpoints

### Expand ValueSet

Expand a ValueSet to get all included concepts.

**Request**

```http
GET /fhir/ValueSet/$expand?url={system}
```

**Example**

```bash
# Expand ICD-11 codes
curl "http://localhost:8080/fhir/ValueSet/\$expand?url=http://id.who.int/icd/release/11/mms"

# Filter results
curl "http://localhost:8080/fhir/ValueSet/\$expand?url=http://id.who.int/icd/release/11/mms&filter=hyper"
```

**Response**

```json
[
  {
    "code": "BA00",
    "display": "Essential hypertension",
    "system": "http://id.who.int/icd/release/11/mms"
  }
]
```

## Error Handling

### Bad Request

```http
HTTP/1.1 400 Bad Request
Content-Type: application/fhir+json

{
  "error": "Invalid JSON"
}
```

### Not Found

```http
HTTP/1.1 404 Not Found
```
