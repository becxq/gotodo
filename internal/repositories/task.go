package repository

import (
	"encoding/json"
	"errors"
	"gotodo/internal/models"
	"os"
)

type Repository struct {
	filePath string
}

func NewRepository(filePath string) *Repository {
	return &Repository{filePath: filePath}
}

func (r *Repository) Contact() {}

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

func (r *Repository) Save(task models.Task) error {
	tasks, err := r.GetAll()
	if err != nil {
		return err
	}

	tasks = append(tasks, task)

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}

func (r *Repository) Delete(id string) error {
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

	data, err := json.MarshalIndent(updatedTasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}
