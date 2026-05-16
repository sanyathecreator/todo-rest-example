package handlers

import (
	"encoding/json"
	"net/http"
	"sanyathecreator/todo-rest-example/internal/models"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	user, err := h.repo.AddUser(r.Context(), models.UserDTO{
		Email:    strings.ToLower(dto.Email),
		Password: hash,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (h *Handlers) LoginUser(w http.ResponseWriter, r *http.Request) {
	var dto models.UserDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondWithError(w, http.StatusBadRequest, "set incorrect data")
		return
	}

	user, err := h.repo.GetUserByEmail(r.Context(), strings.ToLower(dto.Email))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.Password)); err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	token, err := generateToken(user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

func generateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("verysecret"))
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
