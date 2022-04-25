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
var tasksListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of all tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("📝 get the tasks' list")
		/*
			go run main.go tasks list
			TODO: add some filters
			TODO: add some flag to get a json output, table output ...
		*/
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(viper.GetString("procyon-launcher.url") + "/tasks")

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

			fmt.Fprintln(writer, "task-id\tmodule\tstates\tfunction")
			fmt.Fprintln(writer, "-------------------------------------\t-------------------------\t------\t----------------------------")

			var result map[string]interface{}
			json.Unmarshal([]byte(jsonString), &result)

			for key, value := range result {
				content := value.(map[string]interface{})
				config := content["Config"].(map[string]interface{})

				row := key + "\t" + config["WasmFileName"].(string) + "\t" + strconv.FormatFloat(content["State"].(float64), 'f', -1, 64) + "|" + strconv.FormatFloat(content["PreviousState"].(float64), 'f', -1, 64) + "\t" + config["FunctionName"].(string) + "\t" + config["FunctionRevision"].(string) + "(" + strconv.FormatBool(config["DefaultRevision"].(bool)) + ")"

				fmt.Fprintln(writer, row)
			}
			fmt.Fprintln(writer)
			writer.Flush()

		}

	},
}

func init() {
	tasksCmd.AddCommand(tasksListCmd)

}
