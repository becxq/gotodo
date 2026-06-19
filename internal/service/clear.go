package service

func (s *TaskService) Clear() error {
	return s.repo.Clear()
}
