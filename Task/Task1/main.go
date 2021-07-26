package main

import (
	"fmt"
	"task1/api"
	"task1/cmd"

	"github.com/spf13/cobra"
)

var githubUrl = "https://api.github.com/repos"
var rootCmd = &cobra.Command{
	Use:   "Create Table",
	Short: "Create table using this command",
	Long:  `Creates a table if this command is used`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Please provide repository owner name and repo name")
			return
		}
		if args[0] == "" {
			fmt.Println("Please provide repository owner name in first argument")
			return
		}
		owner := args[0]
		if args[1] == "" {
			fmt.Println("Please provide repository name in second argument")
			return
		}
		repo := args[1]

		date, err := api.Call(githubUrl, owner, repo)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(date)
	},
}

func main() {
	cmd.Execute(rootCmd)
}
