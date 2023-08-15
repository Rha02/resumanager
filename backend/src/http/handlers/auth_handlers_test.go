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
		requestBody:        `{"email": "user@test.loc", "password": "testpassword"}`,
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "Wrong request body format",
		requestBody:        `{"email": "user@test.loc", "password": "test"`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Missing email",
		requestBody:        `{"password": "testpassword"}`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Missing password",
		requestBody:        `{"email": "user@test.loc"}`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Not in email format",
		requestBody:        `{"email": "notemail", password": "testpassword"}`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Short password",
		requestBody:        `{"email": "user@test.loc", password": "test"}`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Invalid credentials",
		requestBody:        `{"email": "db_get_user_error@test.loc", "password": "testpassword" }`,
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Error creating access token",
		requestBody: `{
			"email": "access_token_error@test.loc",
			"password": "testpassword"
		}`,
		expectedStatusCode: http.StatusInternalServerError,
	},
	{
		name: "Error creating refresh token",
		requestBody: `{
			"email": "refresh_token_error@test.loc",
			"password": "testpassword"
		}`,
		expectedStatusCode: http.StatusInternalServerError,
	},
}

// TestLogin tests the Login handler.
func TestLogin(t *testing.T) {
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
		name: "Error getting token from blacklist",
		requestHeaders: map[string]string{
			"Authorization": "Bearer cache_error",
		},
		expectedStatusCode: http.StatusInternalServerError,
	},
	{
		name: "Token found in blacklist",
		requestHeaders: map[string]string{
			"Authorization": "Bearer blacklisted_token",
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

var registerTests = []struct {
	name               string
	requestBody        string
	expectedStatusCode int
}{
	{
		name: "Valid register",
		requestBody: `{
			"email": "user@test.loc",
			"username": "testuser",
			"password": "testpassword"
		}`,
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "Wrong request body format",
		requestBody:        `{"email": "user@test.loc", "username": "test", "password": "test"`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Missing email",
		requestBody:        `{"username": "testuser","password": "testpassword"}`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Missing username",
		requestBody:        `{"email": "user@test.loc", "password": "testpassword"}`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Missing password",
		requestBody:        `{"email": "user@test.loc", "username": "testuser"}`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Not in email format",
		requestBody:        `{"email": "notemail", "username": "test", "password": "test"}`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Short username",
		requestBody:        `{"email": "user@test.loc", "username": "te", password": "testpassword"}`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Short password",
		requestBody:        `{"email": "user@test.loc", "username": "testuser", password": "test"}`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name:               "Hashing error",
		requestBody:        `{"email": "user@test.loc", "username": "testuser", "password": "hash_error"}`,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name: "Error creating user",
		requestBody: `{
			"email": "user@test.loc",
			"username": "db_create_user_error",
			"password": "testpassword"
		}`,
		expectedStatusCode: http.StatusInternalServerError,
	},
}

func TestRegister(t *testing.T) {
	for _, tt := range registerTests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/register", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, tt.expectedStatusCode)
			}
		})
	}
}

func TestLogout(t *testing.T) {
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
