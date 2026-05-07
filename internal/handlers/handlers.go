package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sanyathecreator/todo-rest-example/internal/models"
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

func (h *Handlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.repo.GetAllTasks(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, tasks)
}

func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := parsePath(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "wrong ID provided")
		return
	}

	task, err := h.repo.GetTask(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Task does not exist")
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	var dto models.TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondWithError(w, http.StatusBadRequest, "Set incorrect data")
		return
	}

	if strings.TrimSpace(dto.Title) == "" {
		respondWithError(w, http.StatusBadRequest, "Title cannot be null")
		return
	}

	task, err := h.repo.AddTask(r.Context(), dto)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, task)
}

func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := parsePath(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "wrong ID provided")
		return
	}

	_, err = h.repo.GetTask(r.Context(), id)
	if err != nil {
		// TODO: Go convention is lowercase error strings
		respondWithError(w, http.StatusNotFound, "Task does not exist")
		return
	}

	var dto models.UpdateTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondWithError(w, http.StatusBadRequest, "Set incorrect data")
		return
	}

	if dto.Title != nil && strings.TrimSpace(*dto.Title) == "" {
		respondWithError(w, http.StatusBadRequest, "Title cannot be null")
		return
	}

	task, err := h.repo.UpdateTask(r.Context(), id, dto)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("task with id %d not found", id)) {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := parsePath(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "wrong ID provided")
		return
	}

	_, err = h.repo.GetTask(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("task with id %d not found", id))
		return
	}

	err = h.repo.DeleteTask(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("task with id %d not found", id)) {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "success"})
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
