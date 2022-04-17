/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/* 
publishCmd represents the publish command. 
Publish a wasm file to the Procyon Registry

go run main.go registry publish --path ../samples/satellites/forty-two/forty-two.wasm --service forty-two --version 0.0.0
*/
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a wasm file to the Procyon Registry",
	Long: `Publish a wasm file to the Procyon Registry:
procyon-cli registry publish \
--path ../samples/satellites/forty-two/forty-two.wasm \
--service forty-two \
--version 0.0.0
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚢 publishing the wasm file", args)
		
		pathToWasmFile, _ := cmd.Flags().GetString("path")
		serviceName, _ := cmd.Flags().GetString("service")
		wasmModuleVersion, _ := cmd.Flags().GetString("version")

		fmt.Println("📝", pathToWasmFile, "⛎", serviceName, "📦", wasmModuleVersion)
		fmt.Println("🌍", viper.Get("procyon-registry.url"))

		//  go run main.go registry publish --path ../samples/satellites/forty-two/forty-two.wasm --service forty-two --version 0.0.0

		client := resty.New()

		resp, err := client.R().
			SetFile(serviceName, pathToWasmFile).
			Post(viper.GetString("procyon-registry.url")+"/publish/"+wasmModuleVersion)
		
		if err != nil {
			fmt.Println("😡", err)
		} else {
			fmt.Println("🙂", resp.RawResponse)
		}

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


}
