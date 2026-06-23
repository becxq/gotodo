package service

import (
	"github.com/becxq/gotodo/internal/models"
	"time"
)

type Filter func(task models.Task) bool

type FilterKeys struct{
	Priority *int
	Status *bool
	Due *time.Duration
}

func NewPriorityFilter(p int) Filter {return func(task models.Task) bool { return task.Priority == p }}

func NewStatusFilter(s bool) Filter { return func(task models.Task) bool { return task.Status == s } }

func NewDueFilter(d time.Duration) Filter {
	deadline := time.Now().Add(d)
	return func(task models.Task) bool {
		return task.Due.Before(deadline)
	}
}

func (s *TaskService) List(f FilterKeys) ([]models.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil{
		return nil, err
	}

	result := make([]models.Task, 0, len(tasks))

	filters := make([]Filter, 0, 3)

	if f.Priority != nil && *f.Priority >= 1 && *f.Priority <= 3{
		filters = append(filters, NewPriorityFilter(*f.Priority))
	}

	if f.Due != nil{
		filters = append(filters, NewDueFilter(*f.Due))
	}

	if f.Status != nil{
		filters = append(filters, NewStatusFilter(*f.Status))
	}

	for _, task := range tasks {
		pass := true

		for _, filter := range filters{
			if !filter(task){
				pass = false
				break
			}
		}

		if !pass{
			continue
		}

		result = append(result, task)
	}

	return result, nil
}
