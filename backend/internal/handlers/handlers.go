package handlers

import (
	"encoding/json"
	"net/http"
	"sanyathecreator/todo-rest-example/internal/repository"
	"strconv"
	"strings"
)

type Handlers struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handlers {
	return &Handlers{
		repo: repo,
	}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}

func parsePath(r *http.Request) (int, error) {
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tasks/"), "/")
	idStr := pathParts[0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}
