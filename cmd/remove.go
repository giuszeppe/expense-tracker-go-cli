/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/giuszeppe/expense-tracker-go-cli/services"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an example given an ID",
	Long:  `Delete an example given an ID, returns an error string if ID not found`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if id == -1 {
			log.Fatal("Must provide an ID")
		}
		services.RemoveExpense(id)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.PersistentFlags().Int(
		"id", -1, "The id to remove",
	)
}
