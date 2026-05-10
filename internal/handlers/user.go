package handlers

import (
	"encoding/json"
	"net/http"
	"sanyathecreator/todo-rest-example/internal/models"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var dto models.UserDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondWithError(w, http.StatusBadRequest, "set incorrect data")
		return
	}

	if strings.TrimSpace(dto.Email) == "" || strings.TrimSpace(dto.Password) == "" {
		respondWithError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	hash, err := hashPassword(dto.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	dto.Password = hash

	user, err := h.repo.AddUser(r.Context(), dto)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
