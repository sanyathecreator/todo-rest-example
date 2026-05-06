package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sanyathecreator/todo-rest-example/internal/repository"
	"sanyathecreator/todo-rest-example/internal/router"

	"github.com/joho/godotenv"
)

var port = "8080"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatal("DATABASE_URL is not set")
	}

	ctx := context.Background()
	conn, err := repository.Connect(ctx, dbURL)
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	repo := repository.New()

	log.Println("Registering handlers")

	router := router.Register(repo)

	log.Printf("Running server on port %s", port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatal(err)
	}
}
