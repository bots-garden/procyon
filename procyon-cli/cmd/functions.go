/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// functionsCmd represents the functions command
var functionsCmd = &cobra.Command{
	Use:   "functions",
	Short: "Interacting with functions",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 functions cmd")
	},
}

func init() {
	rootCmd.AddCommand(functionsCmd)
}
