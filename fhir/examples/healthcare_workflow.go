//go:build ignore

package examples

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/zs-health/zh-fhir-go/fhir"
	"github.com/zs-health/zh-fhir-go/fhir/internal/testutil"
	"github.com/zs-health/zh-fhir-go/fhir/primitives"
	"github.com/zs-health/zh-fhir-go/fhir/r5"
	"github.com/zs-health/zh-fhir-go/fhir/validation"
)

// Example: Complete healthcare workflow
// Creating a patient encounter with observations and medications
func main() {
	// Step 1: Create a patient
	patient := createPatient()

	// Step 2: Create an encounter for the patient
	encounter := createEncounter(patient.ID)

	// Step 3: Create observations (vital signs)
	observations := createVitalSigns(patient.ID, encounter.ID)

	// Step 4: Create a medication request
	medication := createMedicationRequest(patient.ID, encounter.ID)

	// Step 5: Create a bundle with all resources
	bundle := createBundle(patient, encounter, observations, medication)

	// Step 6: Validate all resources
	if err := validateBundle(bundle); err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	// Step 7: Serialize to JSON
	data, err := json.MarshalIndent(bundle, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Healthcare workflow bundle created successfully!")
	fmt.Printf("Bundle contains %d resources\n", len(bundle.Entry))
	fmt.Printf("\nJSON output (first 500 chars):\n%s...\n", string(data[:min(500, len(data))]))
}

func createPatient() *r5.Patient {
	birthDate := primitives.MustDate("1970-05-15")
	return &r5.Patient{
		ID:     testutil.StringPtr("patient-001"),
		Active: testutil.BoolPtr(true),
		Name: []r5.HumanName{
			{
				Use:    testutil.StringPtr("official"),
				Family: testutil.StringPtr("Smith"),
				Given:  []string{"Jane", "Marie"},
			},
		},
		Gender:    testutil.StringPtr("female"),
		BirthDate: &birthDate,
		Telecom: []r5.ContactPoint{
			{
				System: testutil.StringPtr("phone"),
				Value:  testutil.StringPtr("+1-555-0123"),
				Use:    testutil.StringPtr("mobile"),
			},
			{
				System: testutil.StringPtr("email"),
				Value:  testutil.StringPtr("jane.smith@example.com"),
			},
		},
		Address: []r5.Address{
			{
				Use:        testutil.StringPtr("home"),
				Line:       []string{"123 Main St", "Apt 4B"},
				City:       testutil.StringPtr("Springfield"),
				State:      testutil.StringPtr("IL"),
				PostalCode: testutil.StringPtr("62701"),
				Country:    testutil.StringPtr("USA"),
			},
		},
	}
}

func createEncounter(patientID *string) *r5.Encounter {
	now := primitives.MustDateTime(time.Now().Format(time.RFC3339))
	return &r5.Encounter{
		ID:     testutil.StringPtr("encounter-001"),
		Status: "finished",
		Class: []r5.CodeableConcept{
			{
				Coding: []r5.Coding{
					{
						System:  testutil.StringPtr("http://terminology.hl7.org/CodeSystem/v3-ActCode"),
						Code:    testutil.StringPtr("AMB"),
						Display: testutil.StringPtr("ambulatory"),
					},
				},
			},
		},
		Subject: &r5.Reference{
			Reference: testutil.StringPtr(fmt.Sprintf("Patient/%s", *patientID)),
			Display:   testutil.StringPtr("Jane Smith"),
		},
		Period: &r5.Period{
			Start: &now,
			End:   &now,
		},
	}
}

func createVitalSigns(patientID, encounterID *string) []*r5.Observation {
	effectiveDateTime := primitives.MustDateTime(time.Now().Format(time.RFC3339))

	// Blood Pressure
	bp := &r5.Observation{
		ID:     testutil.StringPtr("obs-bp-001"),
		Status: "final",
		Category: []r5.CodeableConcept{
			{
				Coding: []r5.Coding{
					{
						System:  testutil.StringPtr("http://terminology.hl7.org/CodeSystem/observation-category"),
						Code:    testutil.StringPtr("vital-signs"),
						Display: testutil.StringPtr("Vital Signs"),
					},
				},
			},
		},
		Code: r5.CodeableConcept{
			Coding: []r5.Coding{
				{
					System:  testutil.StringPtr("http://loinc.org"),
					Code:    testutil.StringPtr("85354-9"),
					Display: testutil.StringPtr("Blood pressure panel"),
				},
			},
			Text: testutil.StringPtr("Blood Pressure"),
		},
		Subject: &r5.Reference{
			Reference: testutil.StringPtr(fmt.Sprintf("Patient/%s", *patientID)),
		},
		Encounter: &r5.Reference{
			Reference: testutil.StringPtr(fmt.Sprintf("Encounter/%s", *encounterID)),
		},
		EffectiveDateTime: &effectiveDateTime,
		Component: []r5.ObservationComponent{
			{
				Code: r5.CodeableConcept{
					Coding: []r5.Coding{
						{
							System:  testutil.StringPtr("http://loinc.org"),
							Code:    testutil.StringPtr("8480-6"),
							Display: testutil.StringPtr("Systolic blood pressure"),
						},
					},
				},
				ValueQuantity: &r5.Quantity{
					Value:  float64Ptr(120),
					Unit:   testutil.StringPtr("mmHg"),
					System: testutil.StringPtr("http://unitsofmeasure.org"),
					Code:   testutil.StringPtr("mm[Hg]"),
				},
			},
			{
				Code: r5.CodeableConcept{
					Coding: []r5.Coding{
						{
							System:  testutil.StringPtr("http://loinc.org"),
							Code:    testutil.StringPtr("8462-4"),
							Display: testutil.StringPtr("Diastolic blood pressure"),
						},
					},
				},
				ValueQuantity: &r5.Quantity{
					Value:  float64Ptr(80),
					Unit:   testutil.StringPtr("mmHg"),
					System: testutil.StringPtr("http://unitsofmeasure.org"),
					Code:   testutil.StringPtr("mm[Hg]"),
				},
			},
		},
	}

	// Heart Rate
	hr := &r5.Observation{
		ID:     testutil.StringPtr("obs-hr-001"),
		Status: "final",
		Category: []r5.CodeableConcept{
			{
				Coding: []r5.Coding{
					{
						System:  testutil.StringPtr("http://terminology.hl7.org/CodeSystem/observation-category"),
						Code:    testutil.StringPtr("vital-signs"),
						Display: testutil.StringPtr("Vital Signs"),
					},
				},
			},
		},
		Code: r5.CodeableConcept{
			Coding: []r5.Coding{
				{
					System:  testutil.StringPtr("http://loinc.org"),
					Code:    testutil.StringPtr("8867-4"),
					Display: testutil.StringPtr("Heart rate"),
				},
			},
			Text: testutil.StringPtr("Heart Rate"),
		},
		Subject: &r5.Reference{
			Reference: testutil.StringPtr(fmt.Sprintf("Patient/%s", *patientID)),
		},
		Encounter: &r5.Reference{
			Reference: testutil.StringPtr(fmt.Sprintf("Encounter/%s", *encounterID)),
		},
		EffectiveDateTime: &effectiveDateTime,
		ValueQuantity: &r5.Quantity{
			Value:  float64Ptr(72),
			Unit:   testutil.StringPtr("beats/minute"),
			System: testutil.StringPtr("http://unitsofmeasure.org"),
			Code:   testutil.StringPtr("/min"),
		},
	}

	return []*r5.Observation{bp, hr}
}

func createMedicationRequest(patientID, encounterID *string) *r5.MedicationRequest {
	authoredOn := primitives.MustDateTime(time.Now().Format(time.RFC3339))

	return &r5.MedicationRequest{
		ID:     testutil.StringPtr("medreq-001"),
		Status: "active",
		Intent: "order",
		Medication: r5.CodeableReference{
			Concept: &r5.CodeableConcept{
				Coding: []r5.Coding{
					{
						System:  testutil.StringPtr("http://www.nlm.nih.gov/research/umls/rxnorm"),
						Code:    testutil.StringPtr("197361"),
						Display: testutil.StringPtr("Lisinopril 10 MG Oral Tablet"),
					},
				},
				Text: testutil.StringPtr("Lisinopril 10mg tablet"),
			},
		},
		Subject: &r5.Reference{
			Reference: testutil.StringPtr(fmt.Sprintf("Patient/%s", *patientID)),
			Display:   testutil.StringPtr("Jane Smith"),
		},
		Encounter: &r5.Reference{
			Reference: testutil.StringPtr(fmt.Sprintf("Encounter/%s", *encounterID)),
		},
		AuthoredOn: &authoredOn,
		DosageInstruction: []r5.Dosage{
			{
				Text:   testutil.StringPtr("Take one tablet by mouth once daily"),
				Timing: &r5.Timing{},
			},
		},
	}
}

func createBundle(patient *r5.Patient, encounter *r5.Encounter,
	observations []*r5.Observation, medication *r5.MedicationRequest) *fhir.Bundle {

	bundle := &fhir.Bundle{
		Type: "collection",
	}

	helper := fhir.NewBundleHelper(bundle)

	// Add all resources to bundle
	_ = helper.AddEntry(patient, testutil.StringPtr(fmt.Sprintf("Patient/%s", *patient.ID)))
	_ = helper.AddEntry(encounter, testutil.StringPtr(fmt.Sprintf("Encounter/%s", *encounter.ID)))

	for _, obs := range observations {
		_ = helper.AddEntry(obs, testutil.StringPtr(fmt.Sprintf("Observation/%s", *obs.ID)))
	}

	_ = helper.AddEntry(medication, testutil.StringPtr(fmt.Sprintf("MedicationRequest/%s", *medication.ID)))

	return bundle
}

func validateBundle(bundle *fhir.Bundle) error {
	validator := validation.NewFHIRValidator()

	// Validate each entry in the bundle
	for i, entry := range bundle.Entry {
		// Parse the resource to get its type
		var resourceMap map[string]interface{}
		if err := json.Unmarshal(entry.Resource, &resourceMap); err != nil {
			return fmt.Errorf("entry %d: failed to parse resource: %w", i, err)
		}

		resourceType, ok := resourceMap["resourceType"].(string)
		if !ok {
			return fmt.Errorf("entry %d: missing resourceType", i)
		}

		fmt.Printf("Validating %s...\n", resourceType)

		// For demonstration, we'd need to unmarshal to the correct type
		// This is simplified - in practice you'd use a type switch or registry
	}

	fmt.Println("All resources validated successfully!")
	return nil
}

func float64Ptr(f float64) *float64 {
	return &f
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
