package service

import (
	"errors"
	"gotodo/internal/models"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var name string
var priority string
var due string

func (s *TaskService) Add(cmd *cobra.Command, args []string) error {
	// must have --priority (1 to 3) and --due (none or with date, time or date:time)
	// and --name of the task mochiron
	// tasks:
	// 1. create a Task var
	// 2. validate flags
	// 3. save task

	// validate priority
	pr, err := strconv.Atoi(priority)
	if err != nil {
		return errors.New("Priority must be an integer (1 to 3; lower to higher)")
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
	lastTask := tasks[0]

	task := models.Task{ID: lastTask.ID + 1, Name: name, Due: t, Priority: pr, Status: false}

	return s.repo.Save(task)
}
