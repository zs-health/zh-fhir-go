package examples

import (
	"fmt"

	"github.com/zs-health/zh-fhir-go/fhir/r5/profiles/bd"
	"github.com/zs-health/zh-fhir-go/fhir/r5/terminology/icd11"
)

// BangladeshPatientExample demonstrates using localized R5 profiles for Bangladesh
func BangladeshPatientExample() {
	// Create a new Bangladesh Patient profile
	patient := bd.NewBDPatient()
	patient.ID = stringPtrBD("patient-bd-001")

	// Set DGHS standard identifiers
	patient.AddIdentifier("http://dghs.gov.bd/identifier/nid", "19901234567890123")

	// Set Names (English text)
	patient.SetNames("Abul Bashar", "আবুল বাশার")

	// Create an ICD-11 diagnosis
	diagnosis := icd11.NewCodeableConcept("BA00", "Essential hypertension")

	// In a real scenario, this would be attached to an Encounter or Condition
	// Here we just print the result
	fmt.Printf("Created Bangladesh Patient: %s\n", *patient.ID)
	fmt.Printf("Diagnosis: %s (%s)\n", *diagnosis.Text, *diagnosis.Coding[0].Code)
}

func stringPtrBD(s string) *string {
	return &s
}
