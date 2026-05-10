package repository

import (
	"context"
	"errors"
	"fmt"
	"sanyathecreator/todo-rest-example/internal/models"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) AddTask(ctx context.Context, userID int, dto models.TaskDTO) (models.Task, error) {
	var task models.Task

	query := `
        INSERT INTO tasks (user_id, title, description)
        VALUES ($1, $2, $3)
        RETURNING id, user_id, title, description, completed, created_at, completed_at
    `

	err := r.conn.QueryRow(ctx, query, userID, dto.Title, dto.Description).Scan(
		&task.ID, &task.UserID, &task.Title, &task.Description,
		&task.Completed, &task.CreatedAt, &task.CompletedAt,
	)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (r *Repository) UpdateTask(ctx context.Context, userID int, id int, dto models.UpdateTaskDTO) (models.Task, error) {
	task, err := r.GetTask(ctx, userID, id)
	if err != nil {
		return models.Task{}, err
	}

	if dto.Title != nil {
		task.Title = *dto.Title
	}
	if dto.Description != nil {
		task.Description = *dto.Description
	}
	if dto.Completed != nil {
		if task.Completed != *dto.Completed {
			task.ToggleCompletion()
		}
	}

	query := `
        UPDATE tasks
        SET title=$1, description=$2, completed=$3, completed_at=$4
        WHERE user_id=$5 AND id=$6
    `

	_, err = r.conn.Exec(ctx, query,
		task.Title, task.Description, task.Completed, task.CompletedAt, userID, id,
	)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (r *Repository) GetTask(ctx context.Context, userID int, id int) (models.Task, error) {
	var task models.Task

	query := `
        SELECT id, user_id, title, description, completed, created_at, completed_at
        FROM tasks
        WHERE user_id = $1 AND id = $2
    `

	err := r.conn.QueryRow(ctx, query, userID, id).Scan(
		&task.ID, &task.UserID, &task.Title, &task.Description,
		&task.Completed, &task.CreatedAt, &task.CompletedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Task{}, fmt.Errorf("task with id %d not found", id)
		}
		return models.Task{}, err
	}

	return task, nil
}

func (r *Repository) GetAllTasks(ctx context.Context, userID int, completed *bool) ([]models.Task, error) {
	query := `
        SELECT id, user_id, title, description, completed, created_at, completed_at
        FROM tasks
        WHERE user_id = $1
    `

	args := []any{userID}

	if completed != nil {
		query += " AND completed = $2 ORDER BY created_at DESC"
		args = append(args, *completed)
	} else {
		query += " ORDER BY created_at DESC"
	}

	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.ID, &task.UserID, &task.Title, &task.Description,
			&task.Completed, &task.CreatedAt, &task.CompletedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) DeleteTask(ctx context.Context, userID int, id int) error {
	query := `DELETE FROM tasks WHERE user_id = $1 AND id = $2`

	result, err := r.conn.Exec(ctx, query, userID, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("task with id %d not found", id)
	}

	return nil
}
