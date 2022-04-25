/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	resty "github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a function (wasm module) to the Procyon Launcher",
	Long: `Deploy a function (wasm module) to the Procyon Launcher:
# With wapm.io:
procyon-cli functions deploy \
--wasm k33g/hello-world/1.0.1/hello-world.wasm \
--function hello-world \
--revision rev1

# With Procyon Registry:
procyon-cli functions deploy \
--wasm hello-world.1.0.1.wasm \
--function hello-world \
--revision rev1
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 deployment of wasm module in progress ")

		urlToWasmFile, _ := cmd.Flags().GetString("wasm")
		functionName, _ := cmd.Flags().GetString("function")
		revisionName, _ := cmd.Flags().GetString("revision")

		fmt.Println("📝", urlToWasmFile, "⛎", functionName, "📦", revisionName)
		fmt.Println("🌍 registry:", viper.Get("wasm-registry.url"))

		/*
			# wapm.io
			go run main.go functions deploy
				--wasm k33g/hello-world/1.0.1/hello-world.wasm
				--function hello-world
				--revision rev1

			# procyon registry
			go run main.go functions deploy
				--wasm hello-world.1.0.1.wasm
				--function hello-world
				--revision rev1
		*/

		body := map[string]interface{}{
			"executor":         1, // I plan to be able to deploy other kinds of Runnables Launchers (1==Sat)
			"defaultRevision":  false,
			"wasmFileName":     urlToWasmFile,
			"wasmRegistryUrl":  viper.GetString("wasm-registry.url") + "/" + urlToWasmFile,
			"functionName":     functionName,
			"functionRevision": revisionName,
		}

		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post(viper.GetString("procyon-launcher.url") + "/tasks")

		if err != nil {
			fmt.Println("😡", err)
		} else {
			jsonString := resp.String()
			// Json format
			//fmt.Println("🙂", resp.StatusCode(), ":", jsonString) 

			var result map[string]interface{}
			json.Unmarshal([]byte(jsonString), &result)

			config := result["Config"].(map[string]interface{})

			fmt.Println("📦 WasmRegistryUrl:", config["WasmRegistryUrl"].(string))
			fmt.Println("🌍 WasmFunctionHttpPort:", strconv.FormatFloat(config["WasmFunctionHttpPort"].(float64), 'f', -1, 64))
			fmt.Println("⛎ FunctionName:", config["FunctionName"].(string))
			fmt.Println("📝 FunctionRevision:", config["FunctionRevision"].(string), "DefaultRevision:", strconv.FormatBool(config["DefaultRevision"].(bool)))


			fmt.Println("🌍", functionName, "["+revisionName+"]", "url:", viper.GetString("procyon-reverse.url")+"/functions/"+functionName+"/"+revisionName)
		}

	},
}

func init() {
	var urlToWasmFile string
	var functionName string
	var revisionName string
	//var wasmModuleVersion string included in urlToWasmFile

	functionsCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVarP(&urlToWasmFile, "wasm", "w", "", "Path(URL) to wasm module (required)")
	deployCmd.MarkFlagRequired("wasm")

	deployCmd.Flags().StringVarP(&functionName, "function", "f", "", "Function name (required)")
	deployCmd.MarkFlagRequired("function")

	deployCmd.Flags().StringVarP(&revisionName, "revision", "r", "", "Revision name (required)")
	deployCmd.MarkFlagRequired("revision")

}
