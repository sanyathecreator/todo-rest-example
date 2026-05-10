package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InitDB(ctx context.Context, databaseURL string) (*pgx.Conn, error) {
	conn, err := Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	err = CreateTables(ctx, conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Connect(ctx context.Context, databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CreateTables(ctx context.Context, conn *pgx.Conn) error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`
	_, err := conn.Exec(ctx, createUsersTable)
	if err != nil {
		return err
	}

	createTasksTable := `
	CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP DEFAULT NULL
);
	`
	_, err = conn.Exec(ctx, createTasksTable)
	if err != nil {
		return err
	}
	return nil
}
