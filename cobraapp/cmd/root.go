/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A brief description of your application",
	Long:  `A long description of your application`,
	// Version: "1.0.0",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Hello World!")
	},
}

var input string

var uppercaseCmd = &cobra.Command{

	Use:     "uppercase",
	Short:   "uppercase a string",
	Long:    "Input a lowercase string, I will change it to Uppercase",
	Aliases: []string{"upper"},
	Args:    cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		result := uppercase(input)

		fmt.Println(result)

	},
}

var str string

var reverseCmd = &cobra.Command{

	Use:     "reverse",
	Short:   "reverse a string",
	Long:    "Input a string i will reverse it for you",
	Aliases: []string{"reverse"},
	Args:    cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		result := reverse(str)

		fmt.Println(result)

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand()
	// rootCmd.AddCommand(reverseCmd)
	// rootCmd.AddCommand(uppercaseCmd)
	// rootCmd.AddCommand(unzipCmd)
	reverseCmd.Flags().StringVarP(&str, "reverse", "r", "", "Reverse a String")
	uppercaseCmd.Flags().StringVarP(&input, "uppercase", "u", "", "Capitalize a String")

}
