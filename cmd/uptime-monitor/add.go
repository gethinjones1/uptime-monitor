package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [name] [url]",
	Short: "Add a new endpoint to monitor",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		// 1) Get the name
		var name string
		if len(args) >= 1 {
			name = args[0]
		} else {
			fmt.Print("👉 Enter a name for the endpoint: ")
			input, _ := reader.ReadString('\n')
			name = strings.TrimSpace(input)
		}

		// 2) Get the URL
		var url string
		if len(args) >= 2 {
			url = args[1]
		} else {
			fmt.Print("👉 Enter the URL to monitor: ")
			input, _ := reader.ReadString('\n')
			url = strings.TrimSpace(input)
		}

		if err := db.Add(name, url); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to add: %v\n", err)
			os.Exit(1)
		}

		// TODO: call storage.Add(name, url)
		fmt.Printf("✅ Added %s (%s)\n", name, url)
	},
}
