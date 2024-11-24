/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/giuszeppe/expense-tracker-go-cli/services"
	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A summary of your expenses in the last year",
	Long:  `A summary of your expenses in the last year. Can also be used to summarise a month`,
	Run: func(cmd *cobra.Command, args []string) {
		month, _ := cmd.Flags().GetInt("month")

		if month > 12 || month < 0 {
			log.Fatal("Month must be between 1 and 12")
		}

		services.Summary(month)
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.PersistentFlags().Int("month", 0, "Specifies a month to use for the summary. 1 is january, 12 is december")
}
