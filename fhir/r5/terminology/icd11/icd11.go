package icd11

import (
	"github.com/zs-health/zh-fhir-go/fhir/r5/resources"
)

const (
	SystemICD11 = "http://id.who.int/icd/release/11/mms"
)

// NewCoding creates a new FHIR Coding for an ICD-11 code
func NewCoding(code, display string) resources.Coding {
	system := SystemICD11
	return resources.Coding{
		System:  &system,
		Code:    &code,
		Display: &display,
	}
}

// NewCodeableConcept creates a new FHIR CodeableConcept for an ICD-11 code
func NewCodeableConcept(code, display string) resources.CodeableConcept {
	coding := NewCoding(code, display)
	return resources.CodeableConcept{
		Coding: []resources.Coding{coding},
		Text:   &display,
	}
}
