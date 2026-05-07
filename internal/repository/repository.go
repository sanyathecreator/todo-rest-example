package repository

import (
	"context"
	"errors"
	"fmt"
	"sanyathecreator/todo-rest-example/internal/models"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) AddTask(ctx context.Context, dto models.TaskDTO) (models.Task, error) {
	var task models.Task

	query := `
        INSERT INTO tasks (title, description)
        VALUES ($1, $2)
        RETURNING id, title, description, completed, created_at, completed_at
    `

	err := r.conn.QueryRow(ctx, query, dto.Title, dto.Description).Scan(
		&task.ID, &task.Title, &task.Description,
		&task.Completed, &task.CreatedAt, &task.CompletedAt,
	)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (r *Repository) UpdateTask(ctx context.Context, id int, dto models.UpdateTaskDTO) (models.Task, error) {
	task, err := r.GetTask(ctx, id)
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
    WHERE id=$5
	`

	_, err = r.conn.Exec(ctx, query,
		task.Title, task.Description, task.Completed, task.CompletedAt, id,
	)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (r *Repository) GetTask(ctx context.Context, id int) (models.Task, error) {
	var task models.Task

	query := `
	SELECT id, title, description, completed, created_at, completed_at 
	FROM tasks 
	where id = $1;
	`

	err := r.conn.QueryRow(ctx, query, id).Scan(
		&task.ID, &task.Title, &task.Description,
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

func (r *Repository) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	query := `
	SELECT id, title, description, completed, created_at, completed_at 
	FROM tasks 
	order by created_at desc;
	`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.ID, &task.Title, &task.Description,
			&task.Completed, &task.CreatedAt, &task.CompletedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	// check mid-iteration errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) DeleteTask(ctx context.Context, id int) error {
	query := `DELETE FROM tasks WHERE id = $1`

	result, err := r.conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rows := result.RowsAffected()

	if rows == 0 {
		return fmt.Errorf("task with id %d not found", id)
	}

	return nil
}
