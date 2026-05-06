package main

import (
	"fmt"
	"log"
	"net/http"
	"sanyathecreator/todo-rest-example/internal/handlers"
	"sanyathecreator/todo-rest-example/internal/repository"
)

var port = "8080"

func main() {
	repo := repository.New()
	httpHandlers := handlers.New(repo)

	log.Println("Registering handlers")

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", httpHandlers.GetAllTasks)
	mux.HandleFunc("/tasks/create", httpHandlers.CreateTask)
	mux.HandleFunc("/tasks/", methodHandler(httpHandlers))

	log.Printf("Running server on port %s", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Fatal(err)
	}
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
