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

	err = CreateTable(ctx, conn)
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

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	createTasksTable := `
	CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP DEFAULT NULL
);
	`
	_, err := conn.Exec(ctx, createTasksTable)
	if err != nil {
		return err
	}
	return nil
}
