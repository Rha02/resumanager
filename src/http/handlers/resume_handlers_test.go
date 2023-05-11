package handlers

import (
	"context"
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
