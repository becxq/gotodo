package service

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "gotodo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root Command")
	},
}

var AddCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add Command")
	},
}
var RmCmd = &cobra.Command{
	Use: "rm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Rm Command")
	},
}
var ListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List Command")
	},
}
var DoneCmd = &cobra.Command{
	Use: "done",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Done Command")
	},
}
var ClearCmd = &cobra.Command{
	Use: "clear",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Clear Command")
	},
}
