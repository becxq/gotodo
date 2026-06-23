package commands

import (
	"fmt"
	"github.com/becxq/gotodo/internal/service"

	"github.com/spf13/cobra"
)

type CommandManager struct {
	TaskManager *service.TaskService
}

func NewCommandManager(taskManager *service.TaskService) *CommandManager {
	return &CommandManager{TaskManager: taskManager}
}

func (c *CommandManager) NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use: "gotodo",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please, use args")
		},
	}
}

