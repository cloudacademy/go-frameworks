/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var launchCmd = &cobra.Command{
	Use:   "launch rocket <name>",
	Short: "Launches a rocket",
	Long: `Launches a rocket into space for its mission. For example:

	rocketctl launch rocket r1 --countdown=10
	`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("launch called")
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
