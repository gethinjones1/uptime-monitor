package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all monitored endpoints and their health",
	Run: func(cmd *cobra.Command, args []string) {
		eps, err := db.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to list endpoints: %v\n", err)
			os.Exit(1)
		}
		if len(eps) == 0 {
			fmt.Println("No endpoints configured.")
			return
		}

		stats, err := db.ListStatuses()
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to list statuses: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%-10s │ %-30s │ %-7s │ %s\n",
			"NAME", "URL", "STATUS", "LAST CHECKED")
		fmt.Println(strings.Repeat("─", 10), "┼", strings.Repeat("─", 30),
			"┼", strings.Repeat("─", 7), "┼", strings.Repeat("─", 20))

		for name, url := range eps {
			st, ok := stats[name]
			status := "UNKNOWN"
			ts := ""
			if ok {
				if st.Healthy {
					status = "UP"
				} else {
					status = "DOWN"
				}
				ts = st.LastChecked.Format(time.RFC822)
			}
			fmt.Printf("%-10s │ %-30s │ %-7s │ %s\n",
				name, url, status, ts)
		}
	},
}
