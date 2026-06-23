package repository

import (
	"errors"
	"github.com/becxq/gotodo/internal/models"
)

func (r *Repository) Delete(id int) error {
	tasks, err := r.GetAll()
	if err != nil {
		return err
	}

	var updatedTasks []models.Task
	found := false

	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !found {
		return errors.New("cannot delete: task not found")
	}

	return r.Save(updatedTasks)
}
