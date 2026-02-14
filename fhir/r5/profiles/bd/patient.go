package bd

import (
	"github.com/zs-health/zh-fhir-go/fhir/r5/resources"
)

const (
	ProfileBDPatient = "https://health.zarishsphere.com/fhir/StructureDefinition/bd-patient"
	ExtensionNID     = "https://health.zarishsphere.com/fhir/StructureDefinition/bd-nid"
	ExtensionBRN     = "https://health.zarishsphere.com/fhir/StructureDefinition/bd-brn"
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

// AddNID adds a National ID extension to the patient
func (p *BDPatient) AddNID(nid string) {
	url := ExtensionNID
	ext := resources.Extension{
		URL:        &url,
		ValueString: &nid,
	}
	p.Extension = append(p.Extension, ext)
}
