package repository

import (
	"context"
	"errors"
	"fmt"
	"sanyathecreator/todo-rest-example/internal/models"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) AddUser(ctx context.Context, dto models.UserDTO) (models.User, error) {
	var user models.User

	query := `
        INSERT INTO users (email, password_hash)
        VALUES ($1, $2)
        RETURNING id, email, password_hash, created_at
    `

	err := r.conn.QueryRow(ctx, query, dto.Email, dto.Password).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	query := `
	SELECT id, email, password_hash, created_at
	FROM users 
	where email = $1;
	`

	err := r.conn.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, fmt.Errorf("user with email %s not found", email)
		}
		return models.User{}, err
	}

	return user, nil
}
