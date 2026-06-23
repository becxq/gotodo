package commands

import (
	"fmt"
	"gotodo/internal/service"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	str2duration "github.com/xhit/go-str2duration/v2"
)

func (c *CommandManager) NewAddCmd() *cobra.Command {
	var (
		priority int
 		due string
 		name string
	)

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
	var id string
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
	var id string
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
	var (
		p int
		s bool
		d string
	)

	listCmd := &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			filters := service.FilterKeys{
				Priority: nil,
				Status: nil,
				Due: nil,
			}

			if cmd.Flags().Changed("priority") {
				filters.Priority = &p
			}

			if d != ""{
				dur, err := str2duration.ParseDuration(d)
				if err != nil{
					fmt.Println("Error", err)
					return
				}

				filters.Due = &dur
			}

			if cmd.Flags().Changed("status"){
				filters.Status = &s
			}

			tasks, err := c.TaskManager.List(filters)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			for _, task := range tasks {
				priority := "Low"
				switch task.Priority {
				case 1:
					priority = "Low"
				case 2:
					priority = "Medium"
				case 3:
					priority = "High"
				}

				fmt.Printf("ID: %d| %s Task: %s is ", task.ID, priority, task.Name)

				if task.Status {
					color.New(color.FgGreen).Print("Done ")
				} else {
					color.New(color.FgYellow).Print("Undone ")
				}

				fmt.Printf("till %s\n", task.Due.Format("2006-01-02 15:04"))
			}
		},
	}

	listCmd.Flags().IntVarP(&p, "priority", "p", 0, "Set a priority for a tasks; from 1 to 3")
	listCmd.Flags().BoolVarP(&s, "status", "s", false, "Set a status for a tasks; =true or =false")
	listCmd.Flags().StringVarP(&d, "due", "d", "", "Set a due for a task; [num][unit]")

	listCmd.Flags().Lookup("status").DefValue = ""

	return listCmd
}

func (c *CommandManager) NewClearCmd() *cobra.Command {
	return &cobra.Command{
		Use: "clear",
		Run: func(cmd *cobra.Command, args []string) {
			c.TaskManager.Clear()
		},
	}
}
