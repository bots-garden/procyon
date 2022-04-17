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
		//fmt.Println(viper.Get("procyon-registry"))
		//fmt.Println(viper.Get("procyon-registry.url"))

		//  go run main.go registry url 

	},
}

func init() {
	registryCmd.AddCommand(urlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// urlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// urlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
