package main

import (
		"fmt"
		"gotodo/internal/service"
		"os"
)

func main(){
		if err := service.RootCmd.Execute(); err != nil{
				fmt.Println(err)
				os.Exit(1)
		}
}
