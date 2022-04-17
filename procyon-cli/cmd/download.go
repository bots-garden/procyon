/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download <path_to_wasm_file>",
	Short: "Download a wasm file from the Procyon Registry",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📦 downloading the wasm file")
	},
}

func init() {

	var wasmFileName string
	var pathToWasmFileOutput string

	registryCmd.AddCommand(downloadCmd)

	publishCmd.Flags().StringVarP(&wasmFileName, "wasm", "w", "", "Wasm file name (required)")
	publishCmd.MarkFlagRequired("wasm")

	publishCmd.Flags().StringVarP(&pathToWasmFileOutput, "output", "o", "", "Output wasm file path (required)")
	publishCmd.MarkFlagRequired("output")
}
