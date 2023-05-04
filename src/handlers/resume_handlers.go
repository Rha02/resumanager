package handlers

import (
	"encoding/json"
	"net/http"
)

func (m *Repository) GetUserResumes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value("userID").(string)

	resumes, err := m.DB.GetUserResumes(userID)
	if err != nil {
		http.Error(w, "Error fetching resumes from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resumes)
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
