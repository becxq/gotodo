package main

import (
	"fmt"
	"gotodo/internal/repository"
	"gotodo/internal/service"
	"os"
)

func main() {
	service.RootCmd.AddCommand(service.AddCmd)
	service.RootCmd.AddCommand(service.RmCmd)
	service.RootCmd.AddCommand(service.ListCmd)
	service.RootCmd.AddCommand(service.DoneCmd)
	service.RootCmd.AddCommand(service.ClearCmd)

	if err := service.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Manager := repository.NewRepository("tasks.json") // must run or build app with: go run cmd/api/main.go
}
