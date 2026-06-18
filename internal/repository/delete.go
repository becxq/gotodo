package repository

import (
	"encoding/json"
	"errors"
	"gotodo/internal/models"
	"os"
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
			break
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !found {
		return errors.New("cannot delete: task not found")
	}

	data, err := json.MarshalIndent(updatedTasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}
