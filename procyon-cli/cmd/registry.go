/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// registryCmd represents the registry command
var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "interacting with the experimental Procyon Registry",
	Long: `The  Procyon Registry is a quick'n dirty wasm file registry.
For production, you should use wapm.io or the GitLab generic package registry,
or something else`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📦 Procyon Registry")
	},
}

func init() {
	rootCmd.AddCommand(registryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
