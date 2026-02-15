package server

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/zs-health/zh-fhir-go/fhir/primitives"
	"github.com/zs-health/zh-fhir-go/fhir/r4"
	"github.com/zs-health/zh-fhir-go/internal/ig"
)

// TerminologyServer handles FHIR terminology operations
type TerminologyServer struct {
	loader *ig.Loader
}

func NewTerminologyServer(loader *ig.Loader) *TerminologyServer {
	return &TerminologyServer{
		loader: loader,
	}
}

// HandleExpand handles the ValueSet/$expand operation
func (s *TerminologyServer) HandleExpand(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	filter := r.URL.Query().Get("filter")

	vs, ok := s.loader.ValueSets[url]
	if !ok {
		// If not found in ValueSets, check if it's a CodeSystem and expand it fully
		cs, ok := s.loader.CodeSystems[url]
		if !ok {
			http.Error(w, "Terminology resource not found", http.StatusNotFound)
			return
		}
		vs = s.expandCodeSystem(cs)
	}

	// Apply filter
	if filter != "" {
		vs = s.filterValueSet(vs, filter)
	}

	w.Header().Set("Content-Type", "application/fhir+json")
	json.NewEncoder(w).Encode(vs)
}

func (s *TerminologyServer) expandCodeSystem(cs *r4.CodeSystem) *r4.ValueSet {
	vs := &r4.ValueSet{
		URL:    cs.URL,
		Status: cs.Status,
		Title:  cs.Title,
	}

	contains := make([]r4.ValueSetExpansionContains, 0, len(cs.Concept))
	for _, c := range cs.Concept {
		code := c.Code
		contains = append(contains, r4.ValueSetExpansionContains{
			System:  cs.URL,
			Code:    &code,
			Display: c.Display,
		})
	}

	vs.Expansion = &r4.ValueSetExpansion{
		Timestamp: primitives.FromTimeDateTime(time.Now()),
		Contains:  contains,
	}

	return vs
}

func (s *TerminologyServer) filterValueSet(vs *r4.ValueSet, filter string) *r4.ValueSet {
	if vs.Expansion == nil {
		return vs
	}

	filter = strings.ToLower(filter)
	newContains := make([]r4.ValueSetExpansionContains, 0)
	for _, c := range vs.Expansion.Contains {
		if (c.Code != nil && strings.Contains(strings.ToLower(*c.Code), filter)) ||
			(c.Display != nil && strings.Contains(strings.ToLower(*c.Display), filter)) {
			newContains = append(newContains, c)
		}
	}

	newVS := *vs
	newVS.Expansion = &r4.ValueSetExpansion{
		Contains: newContains,
	}
	return &newVS
}

// RegisterHandlers registers terminology routes
func (s *TerminologyServer) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/fhir/ValueSet/$expand", s.HandleExpand)
}
