package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Concept represents a simple terminology concept
type Concept struct {
	Code    string `json:"code"`
	Display string `json:"display"`
	System  string `json:"system"`
}

// TerminologyServer represents a lightweight FHIR Terminology Server
type TerminologyServer struct {
	Concepts map[string][]Concept
}

func NewTerminologyServer() *TerminologyServer {
	return &TerminologyServer{
		Concepts: make(map[string][]Concept),
	}
}

func (s *TerminologyServer) AddConcept(system string, code string, display string) {
	s.Concepts[system] = append(s.Concepts[system], Concept{
		Code:    code,
		Display: display,
		System:  system,
	})
}

func (s *TerminologyServer) HandleExpand(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	filter := r.URL.Query().Get("filter")

	concepts, ok := s.Concepts[url]
	if !ok {
		http.Error(w, "ValueSet not found", http.StatusNotFound)
		return
	}

	var filtered []Concept
	for _, c := range concepts {
		if filter == "" || strings.Contains(strings.ToLower(c.Display), strings.ToLower(filter)) || strings.Contains(strings.ToLower(c.Code), strings.ToLower(filter)) {
			filtered = append(filtered, c)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func StartTerminologyServer(port int) {
	server := NewTerminologyServer()
	
	// Add some sample ICD-11 concepts
	server.AddConcept("http://id.who.int/icd/release/11/mms", "BA00", "Essential hypertension")
	server.AddConcept("http://id.who.int/icd/release/11/mms", "1B10", "Tuberculosis of the lung")
	
	// Add Bangladesh Divisions
	server.AddConcept("https://health.zarishsphere.com/fhir/ValueSet/bd-divisions", "DH", "Dhaka")
	server.AddConcept("https://health.zarishsphere.com/fhir/ValueSet/bd-divisions", "CH", "Chattogram")

	http.HandleFunc("/fhir/ValueSet/$expand", server.HandleExpand)
	
	fmt.Printf("Terminology Server starting on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
