/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a wasm file to the Procyon Registry",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚢 publishing the wasm file", args)
		//  go run main.go registry publish --path toto.wasm --service heyToto --version 1.2.3
	},
}

func init() {
	
	var pathToWasmFile string
	var serviceName string
	var wasmModuleVersion string

	registryCmd.AddCommand(publishCmd)

	
	publishCmd.Flags().StringVarP(&pathToWasmFile, "path", "p", "", "Path to wasm file (required)")
	publishCmd.MarkFlagRequired("path")

	publishCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Service name (required)")
	publishCmd.MarkFlagRequired("service")

	publishCmd.Flags().StringVarP(&wasmModuleVersion, "version", "v", "", "Wasm module version (required)")
	publishCmd.MarkFlagRequired("version")

	//fmt.Println(pathToWasmFile, serviceName, wasmModuleVersion)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
