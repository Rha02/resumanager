package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// requiresAuthenticationTests is a table-driven test suite for the RequiresAuthentication middleware
var requiresAuthenticationTests = []struct {
	name               string
	requestHeaders     map[string]string
	expectedStatusCode int
}{
	{
		name: "Valid login",
		requestHeaders: map[string]string{
			"Authorization": "Bearer access_token",
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "No authorization header",
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
			"Authorization": "Bearer " + "error",
		},
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Not an access token",
		requestHeaders: map[string]string{
			"Authorization": "Bearer " + "refresh_token",
		},
		expectedStatusCode: http.StatusUnauthorized,
	},
}

// TestRequiresAuthentication tests the RequiresAuthentication middleware
func TestRequiresAuthentication(t *testing.T) {
	handler := getRoutes()

	for _, tt := range requiresAuthenticationTests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/test-middleware/auth", strings.NewReader(""))
			req.Header.Set("Content-Type", "application/json")
			for k, v := range tt.requestHeaders {
				req.Header.Set(k, v)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, tt.expectedStatusCode)
				t.Errorf(rr.Body.String())
			}
		})
	}
}
