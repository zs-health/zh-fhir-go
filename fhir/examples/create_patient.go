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
	// Example 1: Create a living patient
	patient := createPatient()

	// Marshal to JSON with indentation
	data, err := json.MarshalIndent(patient, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling patient: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("=== Living Patient ===")
	fmt.Println(string(data))

	// Example 2: Create a patient with deceased[x] choice type (boolean)
	deceasedPatient := createDeceasedPatient()

	data, err = json.MarshalIndent(deceasedPatient, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling deceased patient: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n=== Deceased Patient (using deceasedBoolean choice) ===")
	fmt.Println(string(data))

	// Example 3: Patient with deceased date/time
	deceasedWithDate := createDeceasedPatientWithDate()

	data, err = json.MarshalIndent(deceasedWithDate, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling deceased patient with date: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n=== Deceased Patient (using deceasedDateTime choice) ===")
	fmt.Println(string(data))
}

func createPatient() r5.Patient {
	active := true
	birthDate := primitives.MustDate("1974-12-25")

	return r5.Patient{
		ID:     testutil.StringPtr("example"),
		Active: &active,
		Name: []r5.HumanName{
			{
				Use:    testutil.StringPtr("official"),
				Family: testutil.StringPtr("Chalmers"),
				Given:  []string{"Peter", "James"},
			},
			{
				Use:   testutil.StringPtr("usual"),
				Given: []string{"Jim"},
			},
		},
		Gender:    testutil.StringPtr("male"),
		BirthDate: &birthDate,
		Telecom: []r5.ContactPoint{
			{
				System: testutil.StringPtr("phone"),
				Value:  testutil.StringPtr("(03) 5555 6473"),
				Use:    testutil.StringPtr("work"),
				Rank:   testutil.IntPtr(1),
			},
			{
				System: testutil.StringPtr("phone"),
				Value:  testutil.StringPtr("(03) 3410 5613"),
				Use:    testutil.StringPtr("mobile"),
				Rank:   testutil.IntPtr(2),
			},
		},
		Address: []r5.Address{
			{
				Use:        testutil.StringPtr("home"),
				Type:       testutil.StringPtr("both"),
				Line:       []string{"534 Erewhon St"},
				City:       testutil.StringPtr("PleasantVille"),
				State:      testutil.StringPtr("Vic"),
				PostalCode: testutil.StringPtr("3999"),
				Period: &r5.Period{
					Start: datetimePtr("1974-12-25"),
				},
			},
		},
		MaritalStatus: &r5.CodeableConcept{
			Coding: []r5.Coding{
				{
					System:  testutil.StringPtr("http://terminology.hl7.org/CodeSystem/v3-MaritalStatus"),
					Code:    testutil.StringPtr("M"),
					Display: testutil.StringPtr("Married"),
				},
			},
		},
		Contact: []r5.PatientContact{
			{
				Relationship: []r5.CodeableConcept{
					{
						Coding: []r5.Coding{
							{
								System: testutil.StringPtr("http://terminology.hl7.org/CodeSystem/v2-0131"),
								Code:   testutil.StringPtr("N"),
							},
						},
					},
				},
				Name: &r5.HumanName{
					Family: testutil.StringPtr("du Marché"),
					Given:  []string{"Bénédicte"},
				},
				Telecom: []r5.ContactPoint{
					{
						System: testutil.StringPtr("phone"),
						Value:  testutil.StringPtr("+33 (237) 998327"),
					},
				},
				Address: &r5.Address{
					Use:        testutil.StringPtr("home"),
					Type:       testutil.StringPtr("both"),
					Line:       []string{"534 Erewhon St"},
					City:       testutil.StringPtr("PleasantVille"),
					State:      testutil.StringPtr("Vic"),
					PostalCode: testutil.StringPtr("3999"),
					Period: &r5.Period{
						Start: datetimePtr("1974-12-25"),
					},
				},
				Gender: testutil.StringPtr("female"),
				Period: &r5.Period{
					Start: datetimePtr("2012"),
				},
			},
		},
	}
}

func datetimePtr(s string) *primitives.DateTime {
	dt := primitives.MustDateTime(s)
	return &dt
}

// createDeceasedPatient demonstrates using the deceasedBoolean choice type field.
// This shows type-safe handling of FHIR choice r5.
func createDeceasedPatient() r5.Patient {
	active := false
	birthDate := primitives.MustDate("1950-03-15")
	deceased := true // Using boolean choice

	return r5.Patient{
		ID:              testutil.StringPtr("deceased-example"),
		Active:          &active,
		DeceasedBoolean: &deceased, // One of the deceased[x] choices
		Name: []r5.HumanName{
			{
				Use:    testutil.StringPtr("official"),
				Family: testutil.StringPtr("Smith"),
				Given:  []string{"John"},
			},
		},
		Gender:    testutil.StringPtr("male"),
		BirthDate: &birthDate,
	}
}

// createDeceasedPatientWithDate demonstrates using the deceasedDateTime choice type field.
// This shows how to choose a different type for the same choice type field.
// NOTE: Only ONE of deceasedBoolean or deceasedDateTime should be set.
func createDeceasedPatientWithDate() r5.Patient {
	active := false
	birthDate := primitives.MustDate("1950-03-15")
	deceasedDate := primitives.MustDateTime("2023-06-15T14:30:00Z") // Using dateTime choice

	return r5.Patient{
		ID:               testutil.StringPtr("deceased-date-example"),
		Active:           &active,
		DeceasedDateTime: &deceasedDate, // Different choice from deceasedBoolean
		Name: []r5.HumanName{
			{
				Use:    testutil.StringPtr("official"),
				Family: testutil.StringPtr("Johnson"),
				Given:  []string{"Mary"},
			},
		},
		Gender:    testutil.StringPtr("female"),
		BirthDate: &birthDate,
	}
}
