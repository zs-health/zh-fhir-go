package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/zs-health/zh-fhir-go/internal/ig"
)

// Server represents the main FHIR server
type Server struct {
	mu        sync.RWMutex
	resources map[string]map[string]any
	loader    *ig.Loader
	term      *TerminologyServer
}

func NewServer(loader *ig.Loader) *Server {
	return &Server{
		resources: make(map[string]map[string]any),
		loader:    loader,
		term:      NewTerminologyServer(loader),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	// Handle Terminology Service
	if path == "fhir/ValueSet/$expand" {
		s.term.HandleExpand(w, r)
		return
	}

	// Handle Resource operations (/fhir/ResourceName/...)
	if len(parts) >= 2 && parts[0] == "fhir" {
		resourceType := parts[1]
		if len(parts) == 2 {
			if r.Method == http.MethodPost {
				s.handleCreate(w, r, resourceType)
				return
			}
			if r.Method == http.MethodGet {
				s.handleSearch(w, r, resourceType)
				return
			}
		} else if len(parts) == 3 {
			id := parts[2]
			if r.Method == http.MethodGet {
				s.handleRead(w, r, resourceType, id)
				return
			}
			if r.Method == http.MethodPut {
				s.handleUpdate(w, r, resourceType, id)
				return
			}
			if r.Method == http.MethodDelete {
				s.handleDelete(w, r, resourceType, id)
				return
			}
		}
	}

	http.NotFound(w, r)
}

func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request, resourceType string) {
	var resource map[string]any
	if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	resource["id"] = id
	resource["resourceType"] = resourceType

	s.mu.Lock()
	if _, ok := s.resources[resourceType]; !ok {
		s.resources[resourceType] = make(map[string]any)
	}
	s.resources[resourceType][id] = resource
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

func (s *Server) handleRead(w http.ResponseWriter, r *http.Request, resourceType, id string) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	res, ok := s.resources[resourceType][id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/fhir+json")
	json.NewEncoder(w).Encode(res)
}

func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request, resourceType, id string) {
	var resource map[string]any
	if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	resource["id"] = id
	resource["resourceType"] = resourceType

	s.mu.Lock()
	if _, ok := s.resources[resourceType]; !ok {
		s.resources[resourceType] = make(map[string]any)
	}
	s.resources[resourceType][id] = resource
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/fhir+json")
	json.NewEncoder(w).Encode(resource)
}

func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request, resourceType, id string) {
	s.mu.Lock()
	delete(s.resources[resourceType], id)
	s.mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request, resourceType string) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	type Bundle struct {
		ResourceType string `json:"resourceType"`
		Type         string `json:"type"`
		Total        int    `json:"total"`
		Entry        []any  `json:"entry"`
	}

	bundle := Bundle{
		ResourceType: "Bundle",
		Type:         "searchset",
		Entry:        make([]any, 0),
	}

	if resources, ok := s.resources[resourceType]; ok {
		for _, res := range resources {
			bundle.Entry = append(bundle.Entry, struct {
				Resource any `json:"resource"`
			}{Resource: res})
		}
	}

	bundle.Total = len(bundle.Entry)
	w.Header().Set("Content-Type", "application/fhir+json")
	json.NewEncoder(w).Encode(bundle)
}

func (s *Server) Start(port int) {
	addr := fmt.Sprintf(":%d", port)
	log.Printf("FHIR Server starting on %s...", addr)
	log.Fatal(http.ListenAndServe(addr, s))
}
