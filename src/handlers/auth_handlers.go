package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Rha02/resumanager/src/models"
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
)

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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

	accessToken, err := authtokenservice.Repo.CreateAccessToken(map[string]interface{}{
		"id":       1,
		"username": user.Username,
	})
	if err != nil {
		http.Error(w, "Error signing access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := authtokenservice.Repo.CreateRefreshToken(map[string]interface{}{
		"id":       1,
		"username": user.Username,
	})
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

}

func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}

func (m *Repository) Refresh(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Refresh"))
}

func (m *Repository) CheckAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims := ctx.Value(ContextKey{}).(map[string]interface{})

	token := claims["token"].(string)

	// check if user token is not blacklisted
	m.CacheRepo.Get(token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(claims)
}
