package router

import (
	"net/http"
	"sanyathecreator/todo-rest-example/internal/handlers"
	"sanyathecreator/todo-rest-example/internal/middleware"
	"sanyathecreator/todo-rest-example/internal/repository"
)

func Register(repo *repository.Repository) http.Handler {
	httpHandlers := handlers.New(repo)

	mux := http.NewServeMux()
	// public routes
	mux.HandleFunc("/register", httpHandlers.RegisterUser)
	mux.HandleFunc("/login", httpHandlers.LoginUser)

	// protected routes
	mux.Handle("/tasks", middleware.Auth(http.HandlerFunc(httpHandlers.GetAllTasks)))
	mux.Handle("/tasks/create", middleware.Auth(http.HandlerFunc(httpHandlers.CreateTask)))
	mux.Handle("/tasks/", middleware.Auth(methodHandler(httpHandlers)))

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
