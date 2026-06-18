package models

import "time"

type Task struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Priority int       `json:"priority"`
	Due      time.Time `json:"time"`
	Status   bool      `json:"status"`
}

type TaskRepository interface {
	Create(task Task) error
	Remove(id string) error
	GetAll() ([]Task, error)
	Check(id string) error
	Clear() error
}
