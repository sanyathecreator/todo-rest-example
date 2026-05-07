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

// TODOs:
// - move DB connection code to database file
// - rewrite all repository function for using DB with SQL queries
// - create DB table
// - wire conn into repository.New()
// - add queries in handlers for filtering(completed/uncompleted tasks)

// - Fix the ToggleCompletion bug in task.go:39 —
// 		*t.CompletedAt panics because the pointer is nil, you need to create a new

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

	repo := repository.New() // FIX: still passes nothing, so the repo still uses the in-memory map

	log.Println("Registering handlers")

	router := router.Register(repo)

	log.Printf("Running server on port %s", port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatal(err)
	}
}
