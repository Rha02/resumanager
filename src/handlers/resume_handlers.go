package handlers

import (
	"encoding/json"
	"net/http"
)

func (m *Repository) GetUserResumes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("GetUserResumes")
}

func (m *Repository) GetResume(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("GetResume")
}

func (m *Repository) PostResume(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("PostResume")
}

func (m *Repository) DeleteResume(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("DeleteResume")
}
