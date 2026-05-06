package main

import (
	"fmt"
	"log"
	"net/http"
	"sanyathecreator/todo-rest-example/internal/repository"
	"sanyathecreator/todo-rest-example/internal/router"
)

var port = "8080"

func main() {
	repo := repository.New()

	log.Println("Registering handlers")

	router := router.Register(repo)

	log.Printf("Running server on port %s", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatal(err)
	}
}
