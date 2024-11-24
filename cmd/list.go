/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/giuszeppe/expense-tracker-go-cli/services"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your expense in a table format",
	Long:  `Will list all your expense in a table format like that: TODO add example`,
	Run: func(cmd *cobra.Command, args []string) {
		services.ListExpenses()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
