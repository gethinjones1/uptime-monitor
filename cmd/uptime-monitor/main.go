package main

import (
	"log"
	"path/filepath"
	"uptime-monitor-cli/internal/storage"

	"github.com/spf13/cobra"
)

var db *storage.DB

func init() {
	// Open (or create) the BoltDB file in the current directory
	path := filepath.Join(".", "endpoints.db")
	var err error
	db, err = storage.OpenDB(path)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
}

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
