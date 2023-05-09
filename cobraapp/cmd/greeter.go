/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name string

// greeterCmd represents the greeter command
var greeterCmd = &cobra.Command{
	Use:   "greeter",
	Short: "I will greet you",
	Long:  `This will greet the user`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Hello %s!", name)
	},
}

func init() {
	rootCmd.AddCommand(greeterCmd)
	greeterCmd.Flags().StringVarP(&name, "name", "n", "", "Enter your name")
	greeterCmd.Flags().BoolP("greet", "t", false, "Help message for toggle")
}
