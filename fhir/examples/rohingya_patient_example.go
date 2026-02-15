package main

import (
	"encoding/json"
	"fmt"
	"github.com/zs-health/zh-fhir-go/fhir/r5/profiles/bd"
	"github.com/zs-health/zh-fhir-go/fhir/r5/resources"
)

func main() {
	// 1. Create a specialized Rohingya Patient
	patient := bd.NewRohingyaPatient()
	patient.ID = stringPtr("rohingya-example-001")

	// 2. Add Specialized Identifiers
	// FCN (Family Counting Number), Progress ID, and MRN (Medical Record Number)
	patient.AddRohingyaIdentifiers(
		"123-456-789", // FCN
		"PROG-998877", // Progress ID
		"MRN-BD-5544", // MRN
	)

	// 3. Set Name
	patient.Name = []resources.HumanName{
		{
			Family: stringPtr("Ali"),
			Given:  []string{"Mohammad"},
		},
	}

	// 4. Set Detailed Shelter Location
	// Camp, Block, Sub-block, and Shelter Number
	patient.SetShelterLocation(
		"Camp 1E",   // Camp
		"Block A",   // Block
		"Sub-Block 2", // Sub-block
		"S-105",     // Shelter Number
	)

	// 5. Add General Bangladesh Address
	patient.Address = []resources.Address{
		{
			City:    stringPtr("Ukhiya"),
			State:   stringPtr("Chattogram"),
			Country: stringPtr("Bangladesh"),
		},
	}

	// 6. Print the localized Rohingya Patient as JSON
	patientJSON, _ := json.MarshalIndent(patient, "", "  ")
	fmt.Println("Localized Rohingya Refugee Patient Resource:")
	fmt.Println(string(patientJSON))
}

func stringPtr(s string) *string { return &s }
