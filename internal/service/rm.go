package service

import "strconv"

func (s *TaskService) Rm(id string) error {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ID)
}
