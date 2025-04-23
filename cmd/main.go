package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "uptime-monitor",
		Short: "Monitor and notify uptime status",
	}

	// Pass db into subcommands via closures or package-level var
	rootCmd.AddCommand(addCmd, lsCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
