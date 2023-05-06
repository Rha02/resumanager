package handlers

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Rha02/resumanager/src/models"
	"github.com/go-chi/chi/v5"
)

type resStruct struct {
	models.Resume
	FileURL string `json:"file_url"`
}

func (m *Repository) GetUserResumes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value("userID").(string)

	resumes, err := m.DB.GetUserResumes(userID)
	if err != nil {
		http.Error(w, "Error fetching resumes from database", http.StatusInternalServerError)
		return
	}

	res := make(map[string]resStruct)

	// add file url to each resume
	for _, resume := range resumes {
		fileURL := m.FileStorage.GetFileURL(resume.FileName)

		res[resume.FileName] = resStruct{
			Resume:  resume,
			FileURL: fileURL,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
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

	fileURL := m.FileStorage.GetFileURL(resume.FileName)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		resStruct{
			Resume:  resume,
			FileURL: fileURL,
		},
	)
}

func (m *Repository) PostResume(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if fileHeader == nil {
		http.Error(w, "Error reading file from form", http.StatusInternalServerError)
		return
	}

	extension := filepath.Ext(fileHeader.Filename)
	if extension != ".pdf" {
		http.Error(w, "File must be a pdf", http.StatusBadRequest)
		return
	}

	// Set max size to 10MB
	r.ParseMultipartForm(10 << 20)

	ctx := r.Context()
	userID := ctx.Value(ContextKey{}).(map[string]interface{})["id"].(float64)
	userIDint := int(userID)

	isMaster := r.FormValue("is_master")
	if isMaster == "" {
		http.Error(w, "Error reading is_master from form", http.StatusInternalServerError)
		return
	}
	isMasterBool, err := strconv.ParseBool(isMaster)
	if err != nil {
		http.Error(w, "Error converting is_master to bool", http.StatusInternalServerError)
		return
	}

	filename, err := m.FileStorage.Upload(file)
	if err != nil {
		http.Error(w, "Error uploading file to storage", http.StatusInternalServerError)
		return
	}

	fileURL := m.FileStorage.GetFileURL(filename)

	resume := models.Resume{
		Name:     fileHeader.Filename,
		FileName: filename,
		UserID:   userIDint,
		Size:     int(fileHeader.Size),
		IsMaster: isMasterBool,
	}

	if err := m.DB.InsertResume(resume); err != nil {
		http.Error(w, "Error creating resume in database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		resStruct{
			Resume:  resume,
			FileURL: fileURL,
		},
	)
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
