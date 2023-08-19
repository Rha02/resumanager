package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Rha02/resumanager/src/models"
)

type response struct {
	models.User
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Email    string `json:"email" validate:"required,min=3,max=254,email"`
		Password string `json:"password" validate:"required,min=8,max=32"`
	}{}

	defer r.Body.Close()
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &body)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	if err = m.validator.Struct(body); err != nil {
		http.Error(w, "Error validating request body", http.StatusBadRequest)
		return
	}

	user, err := m.DB.GetUserByEmail(body.Email)
	if err != nil || !m.hashRepo.ComparePasswords(user.Password, body.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, err := m.AuthTokenRepo.CreateAccessToken(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
	})
	if err != nil {
		http.Error(w, "Error signing access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := m.AuthTokenRepo.CreateRefreshToken(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
	})
	if err != nil {
		http.Error(w, "Error signing refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Email    string `json:"email" validate:"required,min=3,max=254,email"`
		Username string `json:"username" validate:"required,min=3,max=32"`
		Password string `json:"password" validate:"required,min=8,max=32"`
	}{}

	defer r.Body.Close()
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(reqBody, &body); err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	if err = m.validator.Struct(body); err != nil {
		log.Println(err.Error())
		http.Error(w, "Error validating request body", http.StatusBadRequest)
		return
	}

	passwordHash, err := m.hashRepo.HashPassword(body.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusBadRequest)
		return
	}

	user := models.User{
		Email:    body.Email,
		Username: body.Username,
		Password: passwordHash,
	}

	id, err := m.DB.CreateUser(user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	accessToken, err := m.AuthTokenRepo.CreateAccessToken(map[string]interface{}{
		"id":       id,
		"username": user.Username,
	})
	if err != nil {
		http.Error(w, "Error signing access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := m.AuthTokenRepo.CreateRefreshToken(map[string]interface{}{
		"id":       id,
		"username": user.Username,
	})
	if err != nil {
		http.Error(w, "Error signing refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(authHeader) < 7 {
		http.Error(w, "No authorization header", http.StatusUnauthorized)
		return
	}

	refreshToken := authHeader[7:]

	if _, err := m.AuthTokenRepo.ParseRefreshToken(refreshToken); err != nil {
		http.Error(w, "Error parsing token", http.StatusUnauthorized)
		return
	}

	// add refresh token to blacklist
	if err := m.Blacklist.Set(refreshToken, "", 24*7*60*60); err != nil {
		http.Error(w, "Error caching token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully logged out",
	})
}

func (m *Repository) Refresh(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(authHeader) < 7 {
		http.Error(w, "No authorization header", http.StatusUnauthorized)
		return
	}
	refreshToken := authHeader[7:]

	// check if user token is not blacklisted
	v, err := m.Blacklist.Get(refreshToken)
	if err != nil {
		http.Error(w, "Error getting token from blacklist", http.StatusInternalServerError)
		return
	}
	if v != "" {
		http.Error(w, "Token is blacklisted", http.StatusUnauthorized)
		return
	}

	claims, err := m.AuthTokenRepo.ParseRefreshToken(refreshToken)
	if err != nil {
		http.Error(w, "Error parsing token", http.StatusUnauthorized)
		return
	}

	accessToken, err := m.AuthTokenRepo.CreateAccessToken(claims)
	if err != nil {
		http.Error(w, "Error signing access token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": accessToken,
	})
}

func (m *Repository) CheckAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims := ctx.Value(ContextKey{}).(map[string]interface{})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(claims)
}
