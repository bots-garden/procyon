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

// callCmd represents the call command
var callCmd = &cobra.Command{
	Use:   "call",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 calling function")

		method, _ := cmd.Flags().GetString("method")
		functionName, _ := cmd.Flags().GetString("function")
		revisionName, _ := cmd.Flags().GetString("revision")
		data, _ := cmd.Flags().GetString("data")
		reverseUrl := viper.GetString("procyon-reverse.url")

		fmt.Println("📝", method, "⛎", functionName, "📦", revisionName)
		fmt.Println("🌍", viper.Get("procyon-reverse.url"))

		/*
			go run main.go functions call
			  --function hello-world
				--revision rev1
				--method GET

			go run main.go functions call
			  --function hello-world
				--revision rev2
				--method GET

			go run main.go functions call
			  --function hello
				--revision demo
				--method POST
				--data "Bob Morane"
		*/

		client := resty.New()

		if method == "POST" {
			if revisionName=="" {
				resp, err := client.R().
					SetHeader("Content-Type", "application/json").
					SetBody(data).
					Post(reverseUrl + "/functions/" + functionName)
	
				if err != nil {
					fmt.Println("😡", err)
				} else {
					if resp.StatusCode() == 200 {
						fmt.Println("🙂", resp.StatusCode(),":", resp.String())
					} else {
						fmt.Println("😡", resp.StatusCode())
					}				
				}
			} else  {
				resp, err := client.R().
					SetHeader("Content-Type", "application/json").
					SetBody(data).
					Post(reverseUrl + "/functions/" + functionName + "/" + revisionName)
	
				if err != nil {
					fmt.Println("😡", err)
				} else {
					if resp.StatusCode() == 200 {
						fmt.Println("🙂", resp.StatusCode(),":", resp.String())
					} else {
						fmt.Println("😡", resp.StatusCode())
					}				
				}
			}

		} else {
			if method == "GET" {
				//fmt.Println("🎃", reverseUrl, functionName, revisionName, method)
				if revisionName=="" {
					resp, err := client.R().
						Get(reverseUrl + "/functions/" + functionName)
	
					if err != nil {
						fmt.Println("😡", err)
					} else {
						if resp.StatusCode() == 200 {
							fmt.Println("🙂", resp.StatusCode(),":", resp.String())
						} else {
							fmt.Println("😡", resp.StatusCode())
						}					}
				} else  {
					resp, err := client.R().
						Get(reverseUrl + "/functions/" + functionName + "/" + revisionName)
	
					if err != nil {
						fmt.Println("😡", err)
					} else {
						if resp.StatusCode() == 200 {
							fmt.Println("🙂", resp.StatusCode(),":", resp.String())
						} else {
							fmt.Println("😡", resp.StatusCode())
						}
						
					}
				}
			} else {
				// TODO: 🤔
			}
		}


	},
}


func init() {
	var functionName string
	var revisionName string
	var method string
	var data string

	functionsCmd.AddCommand(callCmd)

	callCmd.Flags().StringVarP(&functionName, "function", "f", "", "Function name (required)")
	callCmd.MarkFlagRequired("function")

	callCmd.Flags().StringVarP(&revisionName, "revision", "r", "", "Revision name (required)")
	//callCmd.MarkFlagRequired("revision")

	callCmd.Flags().StringVarP(&method, "method", "m", "POST", "Method (POST or GET)")
	//callCmd.MarkFlagRequired("method")

	callCmd.Flags().StringVarP(&data, "data", "d", "", "Data")
	//callCmd.MarkFlagRequired("data")

}
