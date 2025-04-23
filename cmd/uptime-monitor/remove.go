package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove an endpoint",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		// TODO: call storage.Remove(name)
		fmt.Printf("Removed %s", name)
	},
}
