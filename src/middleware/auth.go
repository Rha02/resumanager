package middleware

import (
	"context"
	"net/http"

	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
)

type ContextKey struct{}

// RequiresAuthentication is a middleware that checks if the request has a valid JWT token
func RequiresAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || len(authHeader) < 7 {
			http.Error(w, "No authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[7:]

		claims, err := authtokenservice.Repo.ParseToken(tokenString)
		if err != nil {
			http.Error(w, "Error parsing token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKey{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
