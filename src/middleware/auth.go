package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
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

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			http.Error(w, "Error parsing token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), ContextKey{}, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		http.Error(w, "Invalid token", http.StatusUnauthorized)
	})
}
