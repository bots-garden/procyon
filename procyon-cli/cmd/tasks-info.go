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

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get details of a task",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📝 get details of a task")

		taskId, _ := cmd.Flags().GetString("task-id")

		/*
			go run main.go tasks info --task-id 1126678e-c2ca-4556-9dbf-57ec06922d39
		*/
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("PROCYON_ADMIN_TOKEN", viper.GetString("procyon-launcher.admin-token")).
			Get(viper.GetString("procyon-launcher.url") + "/tasks/"+taskId)

		if err != nil {
			fmt.Println("😡", err)
		} else {

			// eg 401 Unauthorized
			if resp.IsError() {
				fmt.Println("😡", resp.Status())
			} else {
				jsonString := resp.String()
				// Json format
				//fmt.Println("🙂", resp.StatusCode(),":", jsonString) 
	
				var result map[string]interface{}
				json.Unmarshal([]byte(jsonString), &result)
	
				config := result["Config"].(map[string]interface{})
	
				// Id  WasmFileName State PreviousState
				// Config: FunctionName FunctionRevision":DefaultRevision
				fmt.Println("🙂 [", resp.StatusCode(), "]", result["Id"])
	
	
				fmt.Println("📦 Task Id:", result["Id"].(string))
				fmt.Println("📝 Wasm Module:", config["WasmFileName"].(string))
				fmt.Println("✅ States:", strconv.FormatFloat(result["State"].(float64), 'f', -1, 64) + "|" + strconv.FormatFloat(result["PreviousState"].(float64), 'f', -1, 64))
	
				fmt.Println("⛎ Function:", config["FunctionName"].(string),  config["FunctionRevision"].(string) + "(" + strconv.FormatBool(config["DefaultRevision"].(bool)) + ")")
	
			}

		}
	},
}

func init() {
	var taskId string

	tasksCmd.AddCommand(infoCmd)

	infoCmd.Flags().StringVarP(&taskId, "task-id", "t", "", "Task id (required)")
	infoCmd.MarkFlagRequired("task-id")

}
