package main

import (
	"encoding/json"
	"fmt"
	"github.com/zs-health/zh-fhir-go/fhir/r5/profiles/bd"
	"github.com/zs-health/zh-fhir-go/fhir/r5"
	"github.com/zs-health/zh-fhir-go/fhir/r5/terminology/icd11"
)

func main() {
	// 1. Create a Bangladesh-localized Patient
	patient := bd.NewBDPatient()
	patient.ID = stringPtr("bd-patient-example")
	
	// Add National ID (NID)
	patient.AddNID("1990123456789")
	
	// Set Name
	patient.Name = []resources.HumanName{
		{
			Family: stringPtr("Rahman"),
			Given:  []string{"Ariful"},
		},
	}
	
	// Set Address with Bangladesh Division
	divCoding := bd.GetDivisionCoding("DH")
	patient.Address = []resources.Address{
		{
			City:     stringPtr("Dhaka"),
			District: stringPtr("Dhaka"),
			State:    divCoding.Display,
			Country:  stringPtr("Bangladesh"),
		},
	}

	// 2. Create a Condition using ICD-11
	// Example: Essential hypertension (BA00)
	condition := resources.Condition{
		Subject: &resources.Reference{
			Reference: stringPtr("Patient/bd-patient-example"),
		},
		Code: icd11.NewCodeableConcept("BA00", "Essential hypertension"),
	}

	// 3. Print the localized Patient as JSON
	patientJSON, _ := json.MarshalIndent(patient, "", "  ")
	fmt.Println("Localized Bangladesh Patient:")
	fmt.Println(string(patientJSON))

	// 4. Print the ICD-11 Condition as JSON
	conditionJSON, _ := json.MarshalIndent(condition, "", "  ")
	fmt.Println("\nICD-11 Condition:")
	fmt.Println(string(conditionJSON))
}

func stringPtr(s string) *string { return &s }
func boolPtr(b bool) *bool       { return &b }
