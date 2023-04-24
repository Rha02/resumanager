package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	signingMethod                       = jwt.SigningMethodHS256
	accessTokenExpireTime time.Duration = 15 * 60 // 15 minutes
	jwtKey                              = []byte(os.Getenv("JWT_SECRET"))
)

func createJWTAccessToken(key []byte, expiration int64, sm jwt.SigningMethod) string {
	// Create access token
	accessToken, err := jwt.NewWithClaims(sm, jwt.MapClaims{
		"id":       1,
		"username": "User1",
		"exp":      expiration,
	}).SignedString(key)
	if err != nil {
		panic(err)
	}

	return accessToken
}

// requiresAuthenticationTests is a table-driven test suite for the RequiresAuthentication middleware
var requiresAuthenticationTests = []struct {
	name               string
	requestHeaders     map[string]string
	expectedStatusCode int
}{
	{
		name: "Valid login",
		requestHeaders: map[string]string{
			"Authorization": "Bearer " + createJWTAccessToken(
				jwtKey,
				time.Now().Add(accessTokenExpireTime*time.Second).Unix(),
				signingMethod,
			),
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "No authorization header",
		requestHeaders:     map[string]string{},
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Invalid authorization header",
		requestHeaders: map[string]string{
			"Authorization": "Bearer invalid_token",
		},
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Invalid signing method",
		requestHeaders: map[string]string{
			"Authorization": "Bearer " + createJWTAccessToken(
				jwtKey,
				time.Now().Add(accessTokenExpireTime*time.Second).Unix(),
				jwt.SigningMethodHS512,
			),
		},
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Expired token",
		requestHeaders: map[string]string{
			"Authorization": createJWTAccessToken(
				jwtKey,
				time.Now().Add(-accessTokenExpireTime*time.Second).Unix(),
				signingMethod,
			),
		},
		expectedStatusCode: http.StatusUnauthorized,
	},
	{
		name: "Invalid token",
		requestHeaders: map[string]string{
			"Authorization": "Bearer " + createJWTAccessToken(
				[]byte("invalid"),
				time.Now().Add(accessTokenExpireTime*time.Second).Unix(),
				signingMethod,
			),
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
