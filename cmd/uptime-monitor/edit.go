package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [name] [new-url]",
	Short: "Edit an existing endpoint's URL",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name, newURL := args[0], args[1]
		// TODO: call storage.Edit(name, newURL)
		fmt.Printf("Updated %s to %s", name, newURL)
	},
}
