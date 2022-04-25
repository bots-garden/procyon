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

// revisionCmd represents the revision command
var revisionCmd = &cobra.Command{
	Use:   "revision",
	Short: "Activate the default revision",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("💡 switch the default revision")

		functionName, _ := cmd.Flags().GetString("function")
		revisionName, _ := cmd.Flags().GetString("revision")
		switchValue, _ := cmd.Flags().GetString("switch")

		fmt.Println("📝", switchValue, "⛎", functionName, "📦", revisionName)

		/*
			go run main.go functions revision
			  --function hello-world
				--revision rev1
				--switch off

			go run main.go functions revision
			  --function hello-world
				--revision rev2
				--switch on
		*/

		client := resty.New()
		resp, err := client.R().
			Put(viper.GetString("procyon-launcher.url") + "/revisions/" + functionName + "/" + revisionName + "/default/" + switchValue)

		if err != nil {
			fmt.Println("😡", err)
		} else {
			fmt.Println("🙂", resp.StatusCode())
		}

	},
}

func init() {
	var functionName string
	var revisionName string
	var switchValue string

	functionsCmd.AddCommand(revisionCmd)

	revisionCmd.Flags().StringVarP(&functionName, "function", "f", "", "Function name (required)")
	revisionCmd.MarkFlagRequired("function")

	revisionCmd.Flags().StringVarP(&revisionName, "revision", "r", "", "Revision name (required)")
	revisionCmd.MarkFlagRequired("revision")

	revisionCmd.Flags().StringVarP(&switchValue, "switch", "s", "on", "Switch default revision (on/off)")
	//revisionCmd.MarkFlagRequired("switch")

}
