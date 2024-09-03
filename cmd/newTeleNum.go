/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newTeleNumCmd represents the newTeleNum command
var newTeleNumCmd = &cobra.Command{
	Use:   "newTeleNum",
	Aliases: []string{"ntn"},
	Short: "Alias: 'ntn'.  Desc: Create a new user and add telephone number.",
	Long: `Use this when needed to create a User and add associated telephone number.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("newTeleNum called")
	},
}

func init() {
	rootCmd.AddCommand(newTeleNumCmd)
	// Here you will define your flags and configuration settings.
        updateTeleNumCmd.PersistentFlags().StringP("placeholder", "p", "1234567890", "A short hand and default placeholder for user telephone number.")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newTeleNumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newTeleNumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
