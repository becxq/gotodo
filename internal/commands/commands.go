package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var priority int
var due string
var name string
var id string

func (c *CommandManager) NewAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use: "add",
		Run: func(cmd *cobra.Command, args []string) {
			c.TaskManager.Add(priority, name, due)
		},
	}

	addCmd.Flags().IntVarP(&priority, "priority", "p", 1, "Set a priority for a task; from 1 to 3")
	addCmd.Flags().StringVarP(&due, "due", "d", "30m", "Set a due for a task; [num][unit]")
	addCmd.Flags().StringVarP(&name, "name", "n", "Test", "Give a name for a task")

	addCmd.MarkFlagRequired("name")

	return addCmd
}

func (c *CommandManager) NewRmCmd() *cobra.Command {
	rmCmd := &cobra.Command{
		Use: "rm",
		Run: func(cmd *cobra.Command, args []string) {
			c.TaskManager.Rm(id)
		},
	}

	rmCmd.Flags().StringVarP(&id, "id", "i", "no", "ID")

	rmCmd.MarkFlagRequired("id")

	return rmCmd
}

func (c *CommandManager) NewDoneCmd() *cobra.Command {
	doneCmd := &cobra.Command{
		Use: "done",
		Run: func(cmd *cobra.Command, args []string) {
			c.TaskManager.Done(id)
		},
	}

	doneCmd.Flags().StringVarP(&id, "id", "i", "no", "ID")

	doneCmd.MarkFlagRequired("id")

	return doneCmd
}

func (c *CommandManager) NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := c.TaskManager.List()
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			for i := range tasks {
				fmt.Println(tasks[i].ID, tasks[i].Name, tasks[i].Due, tasks[i].Priority, tasks[i].Status)
			}
		},
	}
}

func (c *CommandManager) NewClearCmd() *cobra.Command {
	return &cobra.Command{
		Use: "clear",
		Run: func(cmd *cobra.Command, args []string) {
			c.TaskManager.Clear()
		},
	}
}
