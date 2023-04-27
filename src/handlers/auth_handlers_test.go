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
	{
		name: "Error creating access token",
		requestBody: `{
			"username": "access_token_error",
			"password": "test"
		}`,
		expectedStatusCode: 500,
	},
	{
		name: "Error creating refresh token",
		requestBody: `{
			"username": "refresh_token_error",
			"password": "test"
		}`,
		expectedStatusCode: 500,
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

var refreshTests = []struct {
	name               string
	requestHeaders     map[string]string
	expectedStatusCode int
}{
	{
		name: "Valid refresh",
		requestHeaders: map[string]string{
			"Authorization": "Bearer refresh_token",
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "Missing authorization header",
		requestHeaders:     map[string]string{},
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Short authorization header",
		requestHeaders: map[string]string{
			"Authorization": "Bearer",
		},
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Invalid token",
		requestHeaders: map[string]string{
			"Authorization": "Bearer error",
		},
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Error creating access token",
		requestHeaders: map[string]string{
			"Authorization": "Bearer creating_access_token_error",
		},
		expectedStatusCode: http.StatusInternalServerError,
	},
}

func TestRefresh(t *testing.T) {
	handler := getRoutes()

	for _, tt := range refreshTests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/refresh", strings.NewReader(""))
			req.Header.Set("Content-Type", "application/json")

			for k, v := range tt.requestHeaders {
				req.Header.Set(k, v)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, tt.expectedStatusCode)
				t.Errorf("response body: %v", rr.Body.String())
			}
		})
	}
}

var logoutTests = []struct {
	name               string
	requestHeaders     map[string]string
	expectedStatusCode int
}{
	{
		name: "Valid logout",
		requestHeaders: map[string]string{
			"Authorization": "Bearer refresh_token",
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name: "Missing authorization header",
		requestHeaders: map[string]string{
			"Authorization": "",
		},
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Short authorization header",
		requestHeaders: map[string]string{
			"Authorization": "Bearer",
		},
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Invalid token",
		requestHeaders: map[string]string{
			"Authorization": "Bearer error",
		},
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Error deleting refresh token",
		requestHeaders: map[string]string{
			"Authorization": "Bearer cache_error",
		},
		expectedStatusCode: http.StatusInternalServerError,
	},
}

func TestLogout(t *testing.T) {
	handler := getRoutes()

	for _, tt := range logoutTests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/logout", strings.NewReader(""))
			req.Header.Set("Content-Type", "application/json")

			for k, v := range tt.requestHeaders {
				req.Header.Set(k, v)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, tt.expectedStatusCode)
				t.Errorf("response body: %v", rr.Body.String())
			}
		})
	}
}
