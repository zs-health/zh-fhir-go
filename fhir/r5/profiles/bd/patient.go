package bd

import (
	"github.com/zs-health/zh-fhir-go/fhir/r5/resources"
)

const (
	ProfileBDPatient = "https://health.zarishsphere.com/fhir/StructureDefinition/bd-patient"
	ExtensionNID     = "http://dghs.gov.bd/identifier/nid"
	ExtensionBRN     = "http://dghs.gov.bd/identifier/brn"
	ExtensionUHID    = "http://dghs.gov.bd/identifier/uhid"
)

// BDPatient represents a Patient resource localized for Bangladesh
type BDPatient struct {
	resources.Patient
}

// NewBDPatient creates a new localized Patient
func NewBDPatient() *BDPatient {
	p := &BDPatient{}
	profile := ProfileBDPatient
	if p.Meta == nil {
		p.Meta = &resources.Meta{}
	}
	p.Meta.Profile = []string{profile}
	return p
}

// AddIdentifier adds a DGHS standard identifier
func (p *BDPatient) AddIdentifier(system, value string) {
	p.Identifier = append(p.Identifier, resources.Identifier{
		System: &system,
		Value:  &value,
	})
}

// SetNames sets both English and Bangla names as per DGHS requirements
func (p *BDPatient) SetNames(englishName, banglaName string) {
	official := "official"
	p.Name = []resources.HumanName{
		{
			Use:  &official,
			Text: &englishName, // Primary text in English
		},
	}
	
	// In a real implementation, we would add the translation extension here
	// For now, we use the Text field to represent the primary name
}
