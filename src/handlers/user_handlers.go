package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Rha02/resumanager/src/models"
)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
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

	w.Write([]byte("Hello " + user.Username))
}
