package repository

import (
	"errors"
	"sanyathecreator/todo-rest-example/internal/models"
)

type Repository struct {
	tasks map[string]models.Task
}

func New() *Repository {
	return &Repository{
		tasks: make(map[string]models.Task),
	}
}

func (r *Repository) AddTask(task models.Task) {
	r.tasks[task.Title] = task
}

func (r *Repository) GetTask(title string) (models.Task, error) {
	task, ok := r.tasks[title]
	if !ok {
		return models.Task{}, errors.New("Task not found")
	}

	return task, nil
}

func (r *Repository) GetTasks() map[string]models.Task {
	tmp := make(map[string]models.Task, len(r.tasks))

	for k, v := range r.tasks {
		tmp[k] = v
	}

	return tmp
}

func (r *Repository) DeleteTask(title string) error {
	_, ok := r.tasks[title]
	if !ok {
		return errors.New("Task not found")
	}

	delete(r.tasks, title)

	return nil
}
