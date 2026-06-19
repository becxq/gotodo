package repository

import (
	"errors"
	"os"
)

func (r *Repository) Clear() error {
	err := os.Truncate("example.txt", 0)
	if err != nil {
		return errors.New("Failed to clear json file")
	}

	return nil
}

