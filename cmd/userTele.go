/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"github.com/spf13/cobra"
)

// userTeleCmd represents the userTele command
var userTeleCmd = &cobra.Command{
	Use:   "userTele",
	Aliases: []string{"ut"},
	Short: "Retrieve user information based on name or telephone number.",
	Long: `Command to retrieve information on a user either based on input of
telephone or name.

This will call local database and obtain what information has been instantiated.`,
	Run: func(cmd *cobra.Command, args []string) {
	    placeholder, _ := cmd.Flags().GetString("placeholder")
        fmt.Println(placeholder)
        // Check to see if user has input placeholder value, if true update telephone number with placeholder.
	    if placeholder != "" {
	        fmt.Println("You have input a placeholder. Would you like to update telephone number for this user?")
	        placeScanner := bufio.NewScanner(os.Stdin)
	        placeScanner.Scan()

	        err := placeScanner.Err()
	        if err != nil {
	            log.Fatal(err)
	        }

            if placeScanner.Text() == "y" || placeScanner.Text() == "yes" {
                fmt.Println("Updated user's telephone number with placeholder.")
            } else {
                fmt.Println("Did not update user telephone. Retrieving User information.")
            }
	    }

        // Make db call to retrieve user information

		fmt.Println("userTele called")
	},
}

func init() {
	rootCmd.AddCommand(userTeleCmd)
	// Here you will define your flags and configuration settings.
       userTeleCmd.PersistentFlags().StringP("placeholder", "p", "", "A placeholder for input user telephone number.")
       userTeleCmd.PersistentFlags().StringP("name", "n", "", "Telephone user's name.")
       userTeleCmd.PersistentFlags().StringP("tele", "t", "", "User's telephone number.")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userTeleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userTeleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
