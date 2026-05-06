package router

import (
	"net/http"
	"sanyathecreator/todo-rest-example/internal/handlers"
	"sanyathecreator/todo-rest-example/internal/repository"
)

func Register(repo *repository.Repository) http.Handler {
	httpHandlers := handlers.New(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", httpHandlers.GetAllTasks)
	mux.HandleFunc("/tasks/create", httpHandlers.CreateTask)
	mux.HandleFunc("/tasks/", methodHandler(httpHandlers))

	return mux
}

func methodHandler(handlers *handlers.Handlers) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTask(w, r)
		case http.MethodPatch:
			handlers.UpdateTask(w, r)
		case http.MethodDelete:
			handlers.DeleteTask(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	}
}
