package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sanyathecreator/todo-rest-example/internal/middleware"
	"sanyathecreator/todo-rest-example/internal/models"
	"strconv"
	"strings"
)

func (h *Handlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	var completed *bool

	// check if filtering is provided
	if raw := r.URL.Query().Get("completed"); raw != "" {
		val, err := strconv.ParseBool(raw)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "completed must be true or false")
			return
		}

		completed = &val
	}

	tasks, err := h.repo.GetAllTasks(r.Context(), userID, completed)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, tasks)
}

func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	id, err := parsePath(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "wrong ID provided")
		return
	}

	task, err := h.repo.GetTask(r.Context(), userID, id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "task does not exist")
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var dto models.TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondWithError(w, http.StatusBadRequest, "set incorrect data")
		return
	}

	if strings.TrimSpace(dto.Title) == "" {
		respondWithError(w, http.StatusBadRequest, "title cannot be null")
		return
	}

	task, err := h.repo.AddTask(r.Context(), userID, dto)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, task)
}

func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	id, err := parsePath(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "wrong ID provided")
		return
	}

	_, err = h.repo.GetTask(r.Context(), userID, id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "task does not exist")
		return
	}

	var dto models.UpdateTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondWithError(w, http.StatusBadRequest, "set incorrect data")
		return
	}

	if dto.Title != nil && strings.TrimSpace(*dto.Title) == "" {
		respondWithError(w, http.StatusBadRequest, "title cannot be null")
		return
	}

	task, err := h.repo.UpdateTask(r.Context(), userID, id, dto)
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
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	id, err := parsePath(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "wrong ID provided")
		return
	}

	_, err = h.repo.GetTask(r.Context(), userID, id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("task with id %d not found", id))
		return
	}

	err = h.repo.DeleteTask(r.Context(), userID, id)
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
