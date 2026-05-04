package models

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
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
