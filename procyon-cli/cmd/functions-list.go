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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of all functions",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📝 get the functions' list")

		/*
			go run main.go functions list
			TODO: add some filters
			TODO: add some flag to get a json output, table output ...
		*/
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(viper.GetString("procyon-launcher.url") + "/functions")

		if err != nil {
			fmt.Println("😡", err)
		} else {
			fmt.Println("🙂", resp.StatusCode(),":", resp.String()) // TODO: less verbose
		}
		
	},
}
// ${PROCYON_URL}/functions
func init() {
	functionsCmd.AddCommand(listCmd)

}
