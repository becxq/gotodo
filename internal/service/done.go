package service

import (
	"errors"
	"strconv"
)

func (s *TaskService) Done(id string) error {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	tasks, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == ID {
			tasks[i].Status = !tasks[i].Status
			found = true
			break
		}
	}

	if found {
		return s.repo.Save(tasks)
	}

	return errors.New("Not found")
}
