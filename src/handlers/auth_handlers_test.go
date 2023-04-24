package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// loginTests is a table of test cases for the Login handler.
var loginTests = []struct {
	name               string
	requestBody        string
	expectedStatusCode int
}{
	{
		name:               "Valid login",
		requestBody:        `{"username": "test", "password": "test"}`,
		expectedStatusCode: 200,
	},
	{
		name:               "Wrong request body format",
		requestBody:        `{"username": "test", "password": "test"`,
		expectedStatusCode: 400,
	},
	{
		name:               "Missing username",
		requestBody:        `{"password": "test"}`,
		expectedStatusCode: 400,
	},
	{
		name:               "Missing password",
		requestBody:        `{"username": "test"}`,
		expectedStatusCode: 400,
	},
}

// TestLogin tests the Login handler.
func TestLogin(t *testing.T) {
	handler := getRoutes()

	for _, tt := range loginTests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/login", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, tt.expectedStatusCode)
			}
		})
	}
}
