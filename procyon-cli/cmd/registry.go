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
	Short: "Interacting with the experimental Procyon Registry",
	Long: `The  Procyon Registry is a quick'n dirty wasm file registry.
For production, you should use wapm.io or the GitLab generic package registry,
or something else`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📦 Procyon Registry")
	},
}

func init() {
	rootCmd.AddCommand(registryCmd)
}
