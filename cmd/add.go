/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/giuszeppe/expense-tracker-go-cli/services"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an expense",
	Long:  `Allows you to add an expense with an amount and a description`,
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetFloat64("amount")
		if amount == 0 {
			log.Fatal("Insert a valid amount")
		}
		description, _ := cmd.Flags().GetString("description")
		if description == "" {
			log.Fatal("Insert a valid description")
		}
		services.AddExpense(amount, description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().String("description", "", "to add a description to the expense")
	addCmd.PersistentFlags().Float64("amount", 0, "to add an amount to the expense")
}
