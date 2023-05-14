package handlers

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

var getUserResumesTests = []struct {
	name               string
	ctxValue           map[string]interface{}
	expectedStatusCode int
}{
	{
		name: "Valid request",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "Missing user id",
		ctxValue:           map[string]interface{}{},
		expectedStatusCode: http.StatusInternalServerError,
	},
	{
		name: "Error getting resumes",
		ctxValue: map[string]interface{}{
			"id": -1.0,
		},
		expectedStatusCode: http.StatusInternalServerError,
	},
}

func TestGetUserResumes(t *testing.T) {
	handler := getRoutes()

	for _, tt := range getUserResumesTests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/resumes", nil)

			ctx := context.WithValue(req.Context(), ContextKey{}, tt.ctxValue)

			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, tt.expectedStatusCode)
				t.Errorf("response body: %v", rr.Body.String())
			}
		})
	}
}

var getResumeTests = []struct {
	name               string
	ctxValue           map[string]interface{}
	resumeID           string
	expectedStatusCode int
}{
	{
		name: "Valid request",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		resumeID:           "1",
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "Missing user id",
		ctxValue:           map[string]interface{}{},
		resumeID:           "1",
		expectedStatusCode: http.StatusInternalServerError,
	},
	{
		name: "Invalid resume id",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		resumeID:           "invalid",
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name: "Error getting resume",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		resumeID:           "-1",
		expectedStatusCode: http.StatusInternalServerError,
	},
	{
		name: "User does not own resume",
		ctxValue: map[string]interface{}{
			"id": 2.0,
		},
		resumeID:           "1",
		expectedStatusCode: http.StatusForbidden,
	},
}

func TestGetResume(t *testing.T) {
	handler := getRoutes()

	for _, tt := range getResumeTests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/resumes/"+tt.resumeID, nil)

			ctx := context.WithValue(req.Context(), ContextKey{}, tt.ctxValue)

			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, tt.expectedStatusCode)
				t.Errorf("response body: %v", rr.Body.String())
			}
		})
	}
}

var postResumeTests = []struct {
	name               string
	ctxValue           map[string]interface{}
	file               string
	expectedStatusCode int
}{
	{
		name: "Valid request",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		file:               "test.pdf",
		expectedStatusCode: http.StatusOK,
	},
	// {
	// 	name:     "Missing user id",
	// 	ctxValue: map[string]interface{}{},
	// 	file:     "test.pdf",
	// 	params: map[string]string{
	// 		"is_master": "true",
	// 	},
	// 	expectedStatusCode: http.StatusInternalServerError,
	// },
	{
		name: "Missing file",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		file:               "",
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name: "Invalid file not pdf",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		file:               "invalid.txt",
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name: "Error getting user resumes",
		ctxValue: map[string]interface{}{
			"id": -1.0,
		},
		file:               "test.pdf",
		expectedStatusCode: http.StatusInternalServerError,
	},
	{
		name: "Error uploading resume",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		file:               "upload_error.pdf",
		expectedStatusCode: http.StatusInternalServerError,
	},
	{
		name: "Error inserting resume to DB",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		file:               "db_insert_resume_error.pdf",
		expectedStatusCode: http.StatusInternalServerError,
	},
}

func TestPostResume(t *testing.T) {
	handler := getRoutes()

	for _, tt := range postResumeTests {
		t.Run(tt.name, func(t *testing.T) {
			pr, pw := io.Pipe()
			defer pw.Close()
			defer pr.Close()

			writer := multipart.NewWriter(pw)

			go func() {
				defer writer.Close()

				if tt.file == "" {
					return
				}

				part, err := writer.CreateFormFile("file", tt.file)
				if err != nil {
					t.Errorf("error creating form file: %v", err)
				}

				data := "testing"
				if tt.file == "upload_error.pdf" {
					data = "error"
				}

				_, err = part.Write([]byte(data))
				if err != nil {
					t.Errorf("error writing to form file: %v", err)
				}
			}()

			req, _ := http.NewRequest("POST", "/resumes", pr)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			ctx := context.WithValue(req.Context(), ContextKey{}, tt.ctxValue)

			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, tt.expectedStatusCode)
				t.Errorf("response body: %v", rr.Body.String())
			}
		})
	}
}

var deleteResumeTests = []struct {
	name               string
	ctxValue           map[string]interface{}
	resumeID           string
	expectedStatusCode int
}{
	{
		name: "Valid request",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		resumeID:           "1",
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "Missing user id",
		ctxValue:           map[string]interface{}{},
		resumeID:           "1",
		expectedStatusCode: http.StatusInternalServerError,
	},
	{
		name: "Invalid resume id",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		resumeID:           "invalid",
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name: "Error getting resume",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		resumeID:           "-1",
		expectedStatusCode: http.StatusInternalServerError,
	},
	{
		name: "User does not own resume",
		ctxValue: map[string]interface{}{
			"id": 2.0,
		},
		resumeID:           "1",
		expectedStatusCode: http.StatusForbidden,
	},
	{
		name: "Error deleting resume",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		resumeID:           "-3",
		expectedStatusCode: http.StatusInternalServerError,
	},
	{
		name: "Error deleting resume file from storage",
		ctxValue: map[string]interface{}{
			"id": 1.0,
		},
		resumeID:           "-2",
		expectedStatusCode: http.StatusInternalServerError,
	},
}

func TestDeleteResume(t *testing.T) {
	handler := getRoutes()

	for _, tt := range deleteResumeTests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("DELETE", "/resumes/"+tt.resumeID, nil)

			ctx := context.WithValue(req.Context(), ContextKey{}, tt.ctxValue)

			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, tt.expectedStatusCode)
				t.Errorf("response body: %v", rr.Body.String())
			}
		})
	}
}
