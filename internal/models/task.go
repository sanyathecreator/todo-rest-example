package models

import "time"

type Task struct {
	ID          int        `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Completed   bool       `json:"completed" db:"completed"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	CompletedAt *time.Time `json:"completed_at" db:"completed_at"`
}

type TaskDTO struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type UpdateTaskDTO struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Completed   *bool   `json:"completed" db:"completed"`
}

func NewTask(title, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) ToggleCompletion() {
	if !t.Completed {
		now := time.Now()
		t.CompletedAt = &now
	} else {
		t.CompletedAt = nil
	}
	t.Completed = !t.Completed
}
