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

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a wasm file from the Procyon Registry",
	Long: `Download a wasm file from the Procyon Registry:
procyon-cli registry download \
--wasm forty-two.0.0.0.wasm \
--output ../forty-two.0.0.0.wasm
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📦 downloading the wasm file")

		wasmFileName, _ := cmd.Flags().GetString("wasm")
		pathToWasmFileOutput, _ := cmd.Flags().GetString("output")

		fmt.Println("📝", wasmFileName, "📦", pathToWasmFileOutput)
		fmt.Println("🌍", viper.Get("procyon-registry.url"))

		/*
		go run main.go registry download --wasm forty-two.0.0.0.wasm --output ../forty-two.0.0.0.wasm
		go run main.go registry download --wasm hello-world.1.0.1.wasm --output ../hello-world.1.0.1.wasm
		go run main.go registry download --wasm hello-world.1.0.2.wasm --output ../hello-world.1.0.2.wasm
		*/

		client := resty.New()
		resp, err := client.R().SetOutput(pathToWasmFileOutput).Get(viper.GetString("procyon-registry.url")+"/get/"+wasmFileName)

		if err != nil {
			fmt.Println("😡", err)
		} else {
			fmt.Println("🙂", resp.RawResponse)
		}


	},
}

func init() {

	var wasmFileName string
	var pathToWasmFileOutput string

	registryCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&wasmFileName, "wasm", "w", "", "Wasm file name (required)")
	downloadCmd.MarkFlagRequired("wasm")

	downloadCmd.Flags().StringVarP(&pathToWasmFileOutput, "output", "o", "", "Output wasm file path (required)")
	downloadCmd.MarkFlagRequired("output")
}
