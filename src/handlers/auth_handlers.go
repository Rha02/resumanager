package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

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
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var user models.User

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
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
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(authHeader) < 7 {
		http.Error(w, "No authorization header", http.StatusUnauthorized)
		return
	}
	tokenString := authHeader[7:]

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrNoCookie
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		http.Error(w, "Error parsing token", http.StatusUnauthorized)
		return
	}

	// Check if token is valid
	if !token.Valid {
		http.Error(w, "Token is not valid", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Error parsing claims", http.StatusUnauthorized)
		return
	}

	userID := int(claims["id"].(float64))

	res := map[string]interface{}{
		"id":       userID,
		"username": claims["username"],
		"exp":      claims["exp"],
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
