package bd

import (
	"github.com/zs-health/zh-fhir-go/fhir/r5/resources"
)

const (
	SystemBDDivisions = "https://health.zarishsphere.com/fhir/ValueSet/bd-divisions"
	SystemBDDistricts = "https://health.zarishsphere.com/fhir/ValueSet/bd-districts"
)

// Bangladesh Divisions
var Divisions = map[string]string{
	"DH": "Dhaka",
	"CH": "Chattogram",
	"RJ": "Rajshahi",
	"KH": "Khulna",
	"BR": "Barishal",
	"SY": "Sylhet",
	"RG": "Rangpur",
	"MY": "Mymensingh",
}

// GetDivisionCoding returns a FHIR Coding for a Bangladesh division
func GetDivisionCoding(code string) *resources.Coding {
	if display, ok := Divisions[code]; ok {
		system := SystemBDDivisions
		return &resources.Coding{
			System:  &system,
			Code:    &code,
			Display: &display,
		}
	}
	return nil
}
