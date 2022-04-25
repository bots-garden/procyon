/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Display the URL of the Procyon Registry",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
    

		fmt.Println("🌍", viper.Get("procyon-registry.url"))


	},
}

func init() {
	registryCmd.AddCommand(urlCmd)
}
