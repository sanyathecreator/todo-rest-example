package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sanyathecreator/todo-rest-example/internal/middleware"
	"sanyathecreator/todo-rest-example/internal/repository"
	"sanyathecreator/todo-rest-example/internal/router"

	"github.com/joho/godotenv"
)

// TODOs:
// - add queries in handlers for filtering(completed/uncompleted tasks)
// - add mutexes

// - middleware - request logging(method, path, status code, latency)
// - Graceful shutdown - Handle SIGTERM/SIGINT, drain in-flight requests, close the DB connection cleanly
// - Sentinel errors - Define ErrNotFound, use errors.Is() in handlers
// - Tests — integration tests against a real test DB
// - Docker Compose — package the app + postgres together

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
	conn, err := repository.InitDB(ctx, dbURL)
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	repo := repository.New(conn)
	log.Println("Registering handlers")

	router := router.Register(repo)
	handler := middleware.Logger(router)

	log.Printf("Running server on port %s", port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	if err != nil {
		log.Fatal(err)
	}
}
