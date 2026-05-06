package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context, databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
