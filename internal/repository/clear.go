package repository

import (
	"os"
)

func (r *Repository) Clear() error {
	err := os.WriteFile(r.filePath , []byte("[]"), 0644)
	if err != nil {
		return err
	}

	return nil
}

