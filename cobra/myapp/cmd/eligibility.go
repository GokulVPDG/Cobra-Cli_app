/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var age int

var eligibilityCmd = &cobra.Command{
	Use:   "eligibility",
	Short: "Check eligibility to vote",
	Long:  `Check whether a person is eligible to vote based on their age.`,

	Run: func(cmd *cobra.Command, args []string) {

		if age >= 18 {

			fmt.Println("You are Eligible to Vote")
		} else {

			fmt.Println("You are not Eligible to Vote")
		}
	},
}

func init() {

	eligibilityCmd.Flags().IntVarP(&age, "age", "a", 0, "Enter your age")
	rootCmd.AddCommand(eligibilityCmd)

}
