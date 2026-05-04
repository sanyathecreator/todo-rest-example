package handlers

import (
	"encoding/json"
	"net/http"
	"sanyathecreator/todo-rest-example/internal/models"
	"sanyathecreator/todo-rest-example/internal/repository"
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
	tasks := h.repo.GetAllTasks()

	respondWithJSON(w, http.StatusOK, tasks)
}

func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {
	title := parsePath(r)

	task, err := h.repo.GetTask(title)
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

	task := models.NewTask(dto.Title, dto.Description)
	h.repo.AddTask(task)

	respondWithJSON(w, http.StatusOK, task)
}

func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	title := parsePath(r)

	_, err := h.repo.GetTask(title)
	if err != nil {
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

	task, err := h.repo.UpdateTask(dto)
	if err != nil {
		if strings.Contains(err.Error(), "Task not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	title := parsePath(r)

	_, err := h.repo.GetTask(title)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Task does not exist")
		return
	}

	err = h.repo.DeleteTask(title)
	if err != nil {
		if strings.Contains(err.Error(), "Task not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "succes"})
}

func parsePath(r *http.Request) string {
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tasks/"), "/")
	title := pathParts[0]

	return title
}
