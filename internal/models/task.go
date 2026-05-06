package models

import "time"

type Task struct {
	ID          int        `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Completed   bool       `json:"completed" db:"completed"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	CompletedAt *time.Time `json:"completed_at" db:"updated_at"`
}

type TaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskDTO struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
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

// BUG: panics when CompletedAt is nil (which it always is for a new task) *t.CompletedAt = time.Now()
func (t *Task) ToggleCompletion() {
	if !t.Completed {
		*t.CompletedAt = time.Now()
	} else {
		t.CompletedAt = nil
	}
	t.Completed = !t.Completed
}
