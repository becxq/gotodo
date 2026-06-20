package commands

import (
	"fmt"

	"github.com/fatih/color"
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
				priority := "Low"
				switch tasks[i].Priority {
				case 1:
					priority = "Low"
				case 2:
					priority = "Medium"
				case 3:
					priority = "High"
				}

				fmt.Printf("ID: %d| %s Task: %s is ", tasks[i].ID, priority, tasks[i].Name)

				if tasks[i].Status {
					color.New(color.FgGreen).Print("Done ")
				} else {
					color.New(color.FgYellow).Print("Undone ")
				}

				fmt.Printf("till %s\n", tasks[i].Due.Format("2026-01-02 15:04:02"))
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
