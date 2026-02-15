package examples

import (
	"fmt"

	"github.com/zs-health/zh-fhir-go/fhir/r5/profiles/bd"
)

// RohingyaPatientExample demonstrates using localized R5 profiles for the Rohingya Response
func RohingyaPatientExample() {
	// Create a new Rohingya Patient profile
	patient := bd.NewRohingyaPatient()
	patient.ID = stringPtrRohingya("rohingya-patient-001")

	// Set standard Identifiers with Rohingya-specific extensions
	patient.AddRohingyaIdentifiers(
		"FCN-123456",    // Family Counting Number
		"PID-987654321", // Progress ID
		"MRN-BD-001",    // Medical Record Number
	)

	// Set detailed shelter location
	patient.SetShelterLocation(
		"Camp 1E",     // Camp
		"Block A",     // Block
		"Sub-block 1", // Sub-block
		"Shelter 101", // Shelter/House Number
	)

	fmt.Printf("Created Rohingya Patient: %s\n", *patient.ID)
}

func stringPtrRohingya(s string) *string {
	return &s
}
