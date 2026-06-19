package main

import (
	"fmt"
	"gotodo/internal/commands"
	"gotodo/internal/repository"
	"gotodo/internal/service"
)

func main() {
	Repository := repository.NewRepository("tasks.json") // must run or build app with: go run cmd/api/main.go
	Service := service.NewTaskService(Repository)
	Runner := commands.NewCommandManager(Service)

	rootCmd := Runner.NewRootCmd()
	rootCmd.AddCommand(Runner.NewAddCmd())
	rootCmd.AddCommand(Runner.NewRmCmd())
	rootCmd.AddCommand(Runner.NewDoneCmd())
	rootCmd.AddCommand(Runner.NewListCmd())
	rootCmd.AddCommand(Runner.NewClearCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
