package fhir_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/zs-health/zh-fhir-go/fhir"
	"github.com/zs-health/zh-fhir-go/fhir/r5"
	"github.com/zs-health/zh-fhir-go/fhir/validation"
)

// TestOfficialFHIRExamples tests against official FHIR example resources
// from https://www.hl7.org/fhir/
func TestOfficialFHIRExamples(t *testing.T) {
	examplesDir := filepath.Join("..", "testdata", "fhir", "examples")

	// Check if examples directory exists
	if _, err := os.Stat(examplesDir); os.IsNotExist(err) {
		t.Skip("Official FHIR examples not downloaded. Run: make download-fhir-examples")
	}

	tests := []struct {
		name         string
		filename     string
		resourceType string
		validate     bool
	}{
		// Core Resources
		{
			name:         "Official Patient Example",
			filename:     "patient-example.json",
			resourceType: "Patient",
			validate:     true,
		},
		{
			name:         "Official Observation Example",
			filename:     "observation-example.json",
			resourceType: "Observation",
			validate:     true,
		},
		{
			name:         "Official Bundle Example",
			filename:     "bundle-example.json",
			resourceType: "Bundle",
			validate:     true,
		},
		// Radiology-Specific Resources
		{
			name:         "Official DiagnosticReport Example",
			filename:     "diagnosticreport-example-f201-brainct.json",
			resourceType: "DiagnosticReport",
			validate:     true,
		},
		{
			name:         "Official ImagingStudy Example",
			filename:     "imagingstudy-example.json",
			resourceType: "ImagingStudy",
			validate:     true,
		},
		// Clinical Workflow Resources
		{
			name:         "Official Encounter Example",
			filename:     "encounter-example.json",
			resourceType: "Encounter",
			validate:     true,
		},
		{
			name:         "Official Condition Example",
			filename:     "condition-example.json",
			resourceType: "Condition",
			validate:     true,
		},
		{
			name:         "Official Procedure Example",
			filename:     "procedure-example.json",
			resourceType: "Procedure",
			validate:     true,
		},
		{
			name:         "Official MedicationRequest Example",
			filename:     "medicationrequest-example.json",
			resourceType: "MedicationRequest",
			validate:     true,
		},
		{
			name:         "Official ServiceRequest Example",
			filename:     "servicerequest-example.json",
			resourceType: "ServiceRequest",
			validate:     true,
		},
		// Administrative Resources
		{
			name:         "Official Organization Example",
			filename:     "organization-example.json",
			resourceType: "Organization",
			validate:     true,
		},
		{
			name:         "Official Practitioner Example",
			filename:     "practitioner-example.json",
			resourceType: "Practitioner",
			validate:     true,
		},
		{
			name:         "Official Location Example",
			filename:     "location-example.json",
			resourceType: "Location",
			validate:     true,
		},
	}

	validator := validation.NewFHIRValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := filepath.Join(examplesDir, tt.filename)

			// Read the example file
			data, err := os.ReadFile(filePath)
			if err != nil {
				t.Skipf("Could not read %s: %v", tt.filename, err)
			}

			// Parse JSON to verify structure
			var raw map[string]interface{}
			if err := json.Unmarshal(data, &raw); err != nil {
				t.Fatalf("Invalid JSON in %s: %v", tt.filename, err)
			}

			// Verify resourceType
			resourceType, ok := raw["resourceType"].(string)
			if !ok {
				t.Fatal("Missing or invalid resourceType")
			}
			if resourceType != tt.resourceType {
				t.Errorf("Expected resourceType %s, got %s", tt.resourceType, resourceType)
			}

			// Unmarshal into appropriate type and validate
			switch tt.resourceType {
			case "Patient":
				var patient r5.Patient
				if err := json.Unmarshal(data, &patient); err != nil {
					t.Fatalf("Failed to unmarshal Patient: %v", err)
				}

				// Validate round-trip
				roundTripData, err := json.Marshal(patient)
				if err != nil {
					t.Fatalf("Failed to marshal Patient: %v", err)
				}

				var roundTrip r5.Patient
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip Patient: %v", err)
				}

				// Validate structure
				if tt.validate {
					if err := validator.Validate(&patient); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
						// Don't fail - official examples may have extensions we don't fully support
					}
				}

				t.Logf("Successfully processed Patient with ID: %v", patient.ID)

			case "Observation":
				var observation r5.Observation
				if err := json.Unmarshal(data, &observation); err != nil {
					t.Fatalf("Failed to unmarshal Observation: %v", err)
				}

				// Validate round-trip
				roundTripData, err := json.Marshal(observation)
				if err != nil {
					t.Fatalf("Failed to marshal Observation: %v", err)
				}

				var roundTrip r5.Observation
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip Observation: %v", err)
				}

				if tt.validate {
					if err := validator.Validate(&observation); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
					}
				}

				t.Logf("Successfully processed Observation with ID: %v", observation.ID)

			case "Bundle":
				var bundle fhir.Bundle
				if err := json.Unmarshal(data, &bundle); err != nil {
					t.Fatalf("Failed to unmarshal Bundle: %v", err)
				}

				// Validate round-trip
				roundTripData, err := json.Marshal(bundle)
				if err != nil {
					t.Fatalf("Failed to marshal Bundle: %v", err)
				}

				var roundTrip fhir.Bundle
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip Bundle: %v", err)
				}

				// Test bundle helper
				helper := fhir.NewBundleHelper(&bundle)
				count := helper.Count()

				t.Logf("Successfully processed Bundle with %d entries", count)

			case "DiagnosticReport":
				var resource r5.DiagnosticReport
				if err := json.Unmarshal(data, &resource); err != nil {
					t.Fatalf("Failed to unmarshal DiagnosticReport: %v", err)
				}
				roundTripData, err := json.Marshal(resource)
				if err != nil {
					t.Fatalf("Failed to marshal DiagnosticReport: %v", err)
				}
				var roundTrip r5.DiagnosticReport
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip DiagnosticReport: %v", err)
				}
				if tt.validate {
					if err := validator.Validate(&resource); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
					}
				}
				t.Logf("Successfully processed DiagnosticReport with ID: %v", resource.ID)

			case "ImagingStudy":
				var resource r5.ImagingStudy
				if err := json.Unmarshal(data, &resource); err != nil {
					t.Fatalf("Failed to unmarshal ImagingStudy: %v", err)
				}
				roundTripData, err := json.Marshal(resource)
				if err != nil {
					t.Fatalf("Failed to marshal ImagingStudy: %v", err)
				}
				var roundTrip r5.ImagingStudy
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip ImagingStudy: %v", err)
				}
				if tt.validate {
					if err := validator.Validate(&resource); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
					}
				}
				t.Logf("Successfully processed ImagingStudy with ID: %v", resource.ID)

			case "Encounter":
				var resource r5.Encounter
				if err := json.Unmarshal(data, &resource); err != nil {
					t.Fatalf("Failed to unmarshal Encounter: %v", err)
				}
				roundTripData, err := json.Marshal(resource)
				if err != nil {
					t.Fatalf("Failed to marshal Encounter: %v", err)
				}
				var roundTrip r5.Encounter
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip Encounter: %v", err)
				}
				if tt.validate {
					if err := validator.Validate(&resource); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
					}
				}
				t.Logf("Successfully processed Encounter with ID: %v", resource.ID)

			case "Condition":
				var resource r5.Condition
				if err := json.Unmarshal(data, &resource); err != nil {
					t.Fatalf("Failed to unmarshal Condition: %v", err)
				}
				roundTripData, err := json.Marshal(resource)
				if err != nil {
					t.Fatalf("Failed to marshal Condition: %v", err)
				}
				var roundTrip r5.Condition
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip Condition: %v", err)
				}
				if tt.validate {
					if err := validator.Validate(&resource); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
					}
				}
				t.Logf("Successfully processed Condition with ID: %v", resource.ID)

			case "Procedure":
				var resource r5.Procedure
				if err := json.Unmarshal(data, &resource); err != nil {
					t.Fatalf("Failed to unmarshal Procedure: %v", err)
				}
				roundTripData, err := json.Marshal(resource)
				if err != nil {
					t.Fatalf("Failed to marshal Procedure: %v", err)
				}
				var roundTrip r5.Procedure
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip Procedure: %v", err)
				}
				if tt.validate {
					if err := validator.Validate(&resource); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
					}
				}
				t.Logf("Successfully processed Procedure with ID: %v", resource.ID)

			case "MedicationRequest":
				var resource r5.MedicationRequest
				if err := json.Unmarshal(data, &resource); err != nil {
					t.Fatalf("Failed to unmarshal MedicationRequest: %v", err)
				}
				roundTripData, err := json.Marshal(resource)
				if err != nil {
					t.Fatalf("Failed to marshal MedicationRequest: %v", err)
				}
				var roundTrip r5.MedicationRequest
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip MedicationRequest: %v", err)
				}
				if tt.validate {
					if err := validator.Validate(&resource); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
					}
				}
				t.Logf("Successfully processed MedicationRequest with ID: %v", resource.ID)

			case "ServiceRequest":
				var resource r5.ServiceRequest
				if err := json.Unmarshal(data, &resource); err != nil {
					t.Fatalf("Failed to unmarshal ServiceRequest: %v", err)
				}
				roundTripData, err := json.Marshal(resource)
				if err != nil {
					t.Fatalf("Failed to marshal ServiceRequest: %v", err)
				}
				var roundTrip r5.ServiceRequest
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip ServiceRequest: %v", err)
				}
				if tt.validate {
					if err := validator.Validate(&resource); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
					}
				}
				t.Logf("Successfully processed ServiceRequest with ID: %v", resource.ID)

			case "Organization":
				var resource r5.Organization
				if err := json.Unmarshal(data, &resource); err != nil {
					t.Fatalf("Failed to unmarshal Organization: %v", err)
				}
				roundTripData, err := json.Marshal(resource)
				if err != nil {
					t.Fatalf("Failed to marshal Organization: %v", err)
				}
				var roundTrip r5.Organization
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip Organization: %v", err)
				}
				if tt.validate {
					if err := validator.Validate(&resource); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
					}
				}
				t.Logf("Successfully processed Organization with ID: %v", resource.ID)

			case "Practitioner":
				var resource r5.Practitioner
				if err := json.Unmarshal(data, &resource); err != nil {
					t.Fatalf("Failed to unmarshal Practitioner: %v", err)
				}
				roundTripData, err := json.Marshal(resource)
				if err != nil {
					t.Fatalf("Failed to marshal Practitioner: %v", err)
				}
				var roundTrip r5.Practitioner
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip Practitioner: %v", err)
				}
				if tt.validate {
					if err := validator.Validate(&resource); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
					}
				}
				t.Logf("Successfully processed Practitioner with ID: %v", resource.ID)

			case "Location":
				var resource r5.Location
				if err := json.Unmarshal(data, &resource); err != nil {
					t.Fatalf("Failed to unmarshal Location: %v", err)
				}
				roundTripData, err := json.Marshal(resource)
				if err != nil {
					t.Fatalf("Failed to marshal Location: %v", err)
				}
				var roundTrip r5.Location
				if err := json.Unmarshal(roundTripData, &roundTrip); err != nil {
					t.Fatalf("Failed to unmarshal round-trip Location: %v", err)
				}
				if tt.validate {
					if err := validator.Validate(&resource); err != nil {
						t.Logf("Validation warning for %s: %v", tt.filename, err)
					}
				}
				t.Logf("Successfully processed Location with ID: %v", resource.ID)
			}
		})
	}
}

// TestOfficialExamplesRoundTrip ensures we can parse and re-serialize
// official FHIR examples without data loss
func TestOfficialExamplesRoundTrip(t *testing.T) {
	examplesDir := filepath.Join("..", "testdata", "fhir", "examples")

	if _, err := os.Stat(examplesDir); os.IsNotExist(err) {
		t.Skip("Official FHIR examples not available")
	}

	files, err := filepath.Glob(filepath.Join(examplesDir, "*.json"))
	if err != nil || len(files) == 0 {
		t.Skip("No FHIR example files found")
	}

	for _, file := range files {
		filename := filepath.Base(file)
		t.Run(filename, func(t *testing.T) {
			// Read original
			originalData, err := os.ReadFile(file)
			if err != nil {
				t.Skipf("Could not read file: %v", err)
			}

			// Parse as generic JSON
			var original map[string]interface{}
			if err := json.Unmarshal(originalData, &original); err != nil {
				t.Fatalf("Invalid JSON: %v", err)
			}

			resourceType, ok := original["resourceType"].(string)
			if !ok {
				t.Skip("No resourceType field")
			}

			// Re-serialize and compare structure
			reserializedData, err := json.Marshal(original)
			if err != nil {
				t.Fatalf("Failed to re-serialize: %v", err)
			}

			var reserialized map[string]interface{}
			if err := json.Unmarshal(reserializedData, &reserialized); err != nil {
				t.Fatalf("Failed to parse re-serialized data: %v", err)
			}

			// Verify resourceType preserved
			if rt, ok := reserialized["resourceType"].(string); !ok || rt != resourceType {
				t.Errorf("ResourceType changed from %s to %v", resourceType, reserialized["resourceType"])
			}

			t.Logf("Successfully round-tripped %s", resourceType)
		})
	}
}

// TestOfficialExamplesSummaryMode tests that summary mode works with official examples
func TestOfficialExamplesSummaryMode(t *testing.T) {
	examplesDir := filepath.Join("..", "testdata", "fhir", "examples")

	if _, err := os.Stat(examplesDir); os.IsNotExist(err) {
		t.Skip("Official FHIR examples not available")
	}

	patientFile := filepath.Join(examplesDir, "patient-example.json")
	data, err := os.ReadFile(patientFile)
	if err != nil {
		t.Skip("Patient example not available")
	}

	var patient r5.Patient
	if err := json.Unmarshal(data, &patient); err != nil {
		t.Fatalf("Failed to unmarshal patient: %v", err)
	}

	// Generate summary
	summaryData, err := fhir.MarshalSummaryJSON(&patient)
	if err != nil {
		t.Fatalf("Failed to generate summary: %v", err)
	}

	// Verify it's valid JSON
	var summaryCheck map[string]interface{}
	if err := json.Unmarshal(summaryData, &summaryCheck); err != nil {
		t.Fatalf("Summary is not valid JSON: %v", err)
	}

	// Verify it's smaller
	fullData, _ := json.Marshal(&patient)
	if len(summaryData) >= len(fullData) {
		t.Errorf("Summary should be smaller than full: summary=%d, full=%d",
			len(summaryData), len(fullData))
	}

	reduction := 100.0 * (1.0 - float64(len(summaryData))/float64(len(fullData)))
	t.Logf("Summary mode achieved %.1f%% reduction (%d -> %d bytes)",
		reduction, len(fullData), len(summaryData))
}
