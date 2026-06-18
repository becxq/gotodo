package repository

import (
	"encoding/json"
	"gotodo/internal/models"
	"os"
)

func (r *Repository) Save(task models.Task) error {
	data, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}
