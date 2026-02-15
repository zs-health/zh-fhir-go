//go:build ignore

package examples

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/zs-health/zh-fhir-go/fhir/internal/testutil"
	"github.com/zs-health/zh-fhir-go/fhir/primitives"
	"github.com/zs-health/zh-fhir-go/fhir/r4/resources"
)

func main() {
	// Create a blood pressure observation
	observation := createBloodPressureObservation()

	// Marshal to JSON
	data, err := json.MarshalIndent(observation, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling observation: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}

func createBloodPressureObservation() r5.Observation {
	effectiveDateTime := primitives.MustDateTime("2024-01-15T10:30:00Z")

	return r5.Observation{
		ID:     testutil.StringPtr("blood-pressure"),
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
			Text: testutil.StringPtr("Blood pressure systolic & diastolic"),
		},
		Subject: r5.Reference{
			Reference: testutil.StringPtr("Patient/example"),
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
		Interpretation: []r5.CodeableConcept{
			{
				Coding: []r5.Coding{
					{
						System:  testutil.StringPtr("http://terminology.hl7.org/CodeSystem/v3-ObservationInterpretation"),
						Code:    testutil.StringPtr("N"),
						Display: testutil.StringPtr("Normal"),
					},
				},
			},
		},
	}
}
