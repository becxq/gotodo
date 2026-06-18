package repository

import (
	"encoding/json"
	"gotodo/internal/models"
	"os"
)

func (r *Repository) GetAll() ([]models.Task, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
