package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out and remove stored CLI authentication token",
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
			fmt.Println("You're already logged out.")
			return nil
		}

		if err := os.Remove(tokenPath); err != nil {
			return fmt.Errorf("failed to remove token: %w", err)
		}

		fmt.Println("✅ Successfully logged out.")
		return nil
	},
}
