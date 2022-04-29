/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	resty "github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of all functions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("📝 functions list")

		/*
			go run main.go functions list
		*/

		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("PROCYON_ADMIN_TOKEN", viper.GetString("procyon-launcher.admin-token")).
			Get(viper.GetString("procyon-launcher.url") + "/functions")

		if err != nil {
			fmt.Println("😡", err)
		} else {
			jsonString := resp.String()

			// Json format
			//fmt.Println("🙂", resp.StatusCode(),":", jsonString)

			// Decoding JSON to Maps - Unstructured Data
			writer := new(tabwriter.Writer)
			// Format in tab-separated columns with a tab stop of 8.
			writer.Init(os.Stdout, 0, 8, 0, '\t', 0)

			fmt.Fprintln(writer, "function-rev\ttask-id\tdefault-revision\thttp-port")
			fmt.Fprintln(writer, "----------------------\t--------------------------------------\t-----------------\t-----------------")

			var result map[string]interface{}
			json.Unmarshal([]byte(jsonString), &result)

			for key, value := range result {
				content := value.(map[string]interface{})

				row := key + "\t" + content["TaskId"].(string) + "\t" + strconv.FormatBool(content["DefaultRevision"].(bool)) + "\t" + strconv.FormatFloat(content["WasmFunctionHttpPort"].(float64), 'f', -1, 64)

				fmt.Fprintln(writer, row)
			}
			fmt.Fprintln(writer)
			writer.Flush()
		}

	},
}

// ${PROCYON_URL}/functions
func init() {
	functionsCmd.AddCommand(listCmd)

}
