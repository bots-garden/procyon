/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Interacting with the running tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🧾 Procyon Launcher Tasks")
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)

}
