package service

import (
		"fmt"
		"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
		Use: "gotodo",
		Run: func(cmd *cobra.Command, args []string){
				fmt.Println("Root Command")
		},
}
