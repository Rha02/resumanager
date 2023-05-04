package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (m *Repository) GetUserResumes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value("userID").(string)

	resumes, err := m.DB.GetUserResumes(userID)
	if err != nil {
		http.Error(w, "Error fetching resumes from database", http.StatusInternalServerError)
		return
	}

	// add file url to each resume
	for _, resume := range resumes {
		resume.FileURL, err = m.FileStorage.GetFileURL(resume.FileName)
		if err != nil {
			http.Error(w, "Error fetching resumes from database", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resumes)
}

func (m *Repository) GetResume(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value("userID").(string)

	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Error converting userID to int", http.StatusInternalServerError)
		return
	}

	resumeID := chi.URLParam(r, "resumeID")

	resume, err := m.DB.GetResume(resumeID)
	if err != nil {
		http.Error(w, "Error fetching resume from database", http.StatusInternalServerError)
		return
	}

	if resume.UserID != userIDint {
		http.Error(w, "Unauthorized to access this resume", http.StatusUnauthorized)
		return
	}

	resume.FileURL, err = m.FileStorage.GetFileURL(resume.FileName)
	if err != nil {
		http.Error(w, "Error fetching resume from database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resume)
}

func (m *Repository) PostResume(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("PostResume")
}

func (m *Repository) DeleteResume(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value("userID").(string)
	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Error converting userID to int", http.StatusInternalServerError)
		return
	}

	resumeID := chi.URLParam(r, "resumeID")

	resume, err := m.DB.GetResume(resumeID)
	if err != nil {
		http.Error(w, "Error fetching resume from database", http.StatusInternalServerError)
		return
	}

	if resume.UserID != userIDint {
		http.Error(w, "Unauthorized to access this resume", http.StatusForbidden)
		return
	}

	if err := m.DB.DeleteResume(resumeID); err != nil {
		http.Error(w, "Error deleting resume from database", http.StatusInternalServerError)
		return
	}

	if _, err := m.FileStorage.Delete(resume.FileName); err != nil {
		http.Error(w, "Error deleting resume from database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully deleted resume",
	})
}
