package service

import (
	"errors"
	"gotodo/internal/models"
	"time"
)

func (s *TaskService) Add(priority int, name, due string) error {
	// must have --priority (1 to 3) and --due (none or with date, time or date:time)
	// and --name of the task mochiron
	// tasks:
	// 1. create a Task var
	// 2. validate flags
	// 3. save task

	// validate priority
	if priority < 1 || priority > 3 {
		return errors.New("Wrong number of priority")
	}

	// checking len of name
	if len(name) < 3 {
		return errors.New("tast text is too short")
	} else if len(name) > 32 {
		return errors.New("task text is too long")
	}

	// so i'm going to do next formats for time:
	// simple date: time for exact day and hour
	// [num][unit]: add user time to current time

	layout := "2006-01-02 15:03"

	var duration time.Duration

	t, err := time.Parse(layout, due)
	if err != nil {
		duration, err = time.ParseDuration(due)
	}

	if err != nil {
		return errors.New("time format is wrong")
	}

	t = time.Now().Add(duration)

	// getting id of last task
	tasks, err := s.repo.GetAll()
	if err != nil {
		return errors.New("error of getting tasks")
	}

	var id int
	if len(tasks) == 0 {
		id = 0
	} else {
		id = tasks[0].ID + 1
	}

	task := models.Task{ID: id, Name: name, Due: t, Priority: priority, Status: false}
	tasks = append(tasks, task)

	return s.repo.Save(tasks)
}
