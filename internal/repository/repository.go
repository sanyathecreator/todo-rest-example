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

func (r *Repository) AddTask(task models.Task) error {
	if _, ok := r.tasks[task.Title]; ok {
		return errors.New("Task already exists")
	}

	r.tasks[task.Title] = task

	return nil
}

func (r *Repository) UpdateTask(dto models.UpdateTaskDTO) (models.Task, error) {
	if _, ok := r.tasks[*dto.Title]; !ok {
		return models.Task{}, errors.New("Task not found")
	}

	task := r.tasks[*dto.Title]

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

	r.tasks[*dto.Title] = task

	return task, nil
}

func (r *Repository) GetTask(title string) (models.Task, error) {
	task, ok := r.tasks[title]
	if !ok {
		return models.Task{}, errors.New("Task not found")
	}

	return task, nil
}

func (r *Repository) GetAllTasks() map[string]models.Task {
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
