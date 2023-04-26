package middleware

import (
	"context"
	"net/http"

	"github.com/Rha02/resumanager/src/handlers"
)

func RequiresAuthentication(repo *handlers.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || len(authHeader) < 7 {
				http.Error(w, "No authorization header", http.StatusUnauthorized)
				return
			}
			tokenString := authHeader[7:]

			claims, err := repo.AuthTokenRepo.ParseAccessToken(tokenString)
			if err != nil {
				http.Error(w, "Error parsing token", http.StatusUnauthorized)
				return
			}

			claims["token"] = tokenString

			ctx := context.WithValue(r.Context(), handlers.ContextKey{}, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
