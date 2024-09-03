/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateTeleNumCmd represents the updateTeleNum command
var updateTeleNumCmd = &cobra.Command{
	Use:   "updateTeleNum",
	Aliases: []string{"utn"},
	Short: "Alias: 'utn'.  Desc: Updates User telephone information.",
	Long: `Use this command to update any user's telephone number associated
to name.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateTeleNum called")
	},
}

func init() {
	rootCmd.AddCommand(updateTeleNumCmd)
	// Here you will define your flags and configuration settings.
        updateTeleNumCmd.PersistentFlags().StringP("placeholder", "p", "1234567890", "A short hand and default placeholder for user telephone number.")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateTeleNumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateTeleNumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
