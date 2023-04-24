package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Rha02/resumanager/src/middleware"
	"github.com/Rha02/resumanager/src/models"
	"github.com/golang-jwt/jwt"
)

var (
	signingMethod                        = jwt.SigningMethodHS256
	accessTokenExpireTime  time.Duration = 15 * 60          // 15 minutes
	refreshTokenExpireTime               = 24 * 7 * 60 * 60 // 24 hours
)

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var user models.User

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	// Check if request has right fields
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username or password is empty", http.StatusBadRequest)
		return
	}

	key := []byte(os.Getenv("JWT_SECRET"))

	// Create access token
	accessToken, err := jwt.NewWithClaims(signingMethod, jwt.MapClaims{
		"id":       1,
		"username": user.Username,
		"exp":      time.Now().Add(accessTokenExpireTime * time.Second).Unix(),
	}).SignedString(key)
	if err != nil {
		http.Error(w, "Error signing access token", http.StatusInternalServerError)
		return
	}

	// Create refresh token
	refreshToken, err := jwt.NewWithClaims(signingMethod, jwt.MapClaims{
		"id":       1,
		"username": user.Username,
		"exp":      refreshTokenExpireTime,
	}).SignedString(key)
	if err != nil {
		http.Error(w, "Error signing refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register"))
}

func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}

func (m *Repository) Refresh(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Refresh"))
}

func (m *Repository) CheckAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims := ctx.Value(middleware.ContextKey{}).(jwt.MapClaims)

	res := map[string]interface{}{
		"id":       claims["id"],
		"username": claims["username"],
		"exp":      claims["exp"],
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
