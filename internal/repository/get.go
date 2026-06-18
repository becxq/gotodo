package repository

import (
	"errors"
	"gotodo/internal/models"
)

func (r *Repository) Get(id string) (models.Task, error) {
	tasks, err := r.GetAll()
	if err != nil {
		return models.Task{}, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return models.Task{}, errors.New("task not found")
}
