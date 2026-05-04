package models

import "time"

type Task struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
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

func (t *Task) ToggleCompletion() {
	if !t.Completed {
		*t.CompletedAt = time.Now()
	} else {
		t.CompletedAt = nil
	}
	t.Completed = !t.Completed
}
