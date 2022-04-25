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

// tasksKillCmd represents the tasksKill command
var tasksKillCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill a specific task",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📝 kill a task")

		taskId, _ := cmd.Flags().GetString("task-id")

		/*
			go run main.go tasks kill --task-id 8c2d4226-8382-4a53-b06d-35f0ad3e34a2
			TODO: lists of functions and tasks are not updated
		*/
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			Delete(viper.GetString("procyon-launcher.url") + "/tasks/"+taskId)

		if err != nil {
			fmt.Println("😡", err)
		} else {
			//fmt.Println("🙂", resp.StatusCode(),":", resp.String()) // TODO: less verbose
			fmt.Println("🙂 [", resp.StatusCode(), "] task", taskId, "is killed")
		}

	},
}

func init() {
	var taskId string

	tasksCmd.AddCommand(tasksKillCmd)

	tasksKillCmd.Flags().StringVarP(&taskId, "task-id", "t", "", "Task id (required)")
	tasksKillCmd.MarkFlagRequired("task-id")

}
