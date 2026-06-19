package repository

import (
	"encoding/json"
	"gotodo/internal/models"
	"os"
)

func (r *Repository) Save(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}
