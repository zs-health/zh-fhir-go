package bd

import (
	"github.com/zs-health/zh-fhir-go/fhir"
	"github.com/zs-health/zh-fhir-go/fhir/r5"
)

const (
	ProfileBDPatient = "https://health.zarishsphere.com/fhir/StructureDefinition/bd-patient"
	ExtensionNID     = "http://dghs.gov.bd/identifier/nid"
	ExtensionBRN     = "http://dghs.gov.bd/identifier/brn"
	ExtensionUHID    = "http://dghs.gov.bd/identifier/uhid"
)

// BDPatient represents a r5.Patient resource localized for Bangladesh
type BDPatient struct {
	r5.Patient
}

// NewBDPatient creates a new localized r5.Patient
func NewBDPatient() *BDPatient {
	p := &BDPatient{}
	profile := ProfileBDPatient
	if p.Meta == nil {
		p.Meta = &fhir.Meta{}
	}
	p.Meta.Profile = []string{profile}
	return p
}

// AddIdentifier adds a DGHS standard identifier
func (p *BDPatient) AddIdentifier(system, value string) {
	p.Identifier = append(p.Identifier, r5.Identifier{
		System: &system,
		Value:  &value,
	})
}

// SetNames sets both English and Bangla names as per DGHS requirements
func (p *BDPatient) SetNames(englishName, banglaName string) {
	official := "official"
	p.Name = []r5.HumanName{
		{
			Use:  &official,
			Text: &englishName, // Primary text in English
		},
	}
}
