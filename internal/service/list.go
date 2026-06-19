package service

import "gotodo/internal/models"

func (s *TaskService) List() ([]models.Task, error) {
	return s.repo.GetAll()
}
