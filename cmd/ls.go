package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type endpointInfo struct {
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	Healthy     bool      `json:"healthy"`
	LastChecked time.Time `json:"lastChecked"`
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all endpoints and their health",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := getAuthToken()
		if err != nil {
			log.Fatalf("auth failed: %v", err)
		}

		endpoint := strings.TrimRight(apiURL, "/") + "/status"
		token, err := getAuthToken()
		if err != nil {
			log.Fatalf("auth failed: %v", err)
		}

		req, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			log.Fatalf("request build failed: %v", err)
		}
		req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			log.Fatalf("API error (%d): %s", resp.StatusCode, strings.TrimSpace(string(body)))
		}

		var eps []endpointInfo
		if err := json.NewDecoder(resp.Body).Decode(&eps); err != nil {
			log.Fatalf("invalid JSON: %v", err)
		}

		if len(eps) == 0 {
			fmt.Println("No endpoints registered.")
			return
		}

		fmt.Printf("%-12s  %-30s  %-6s  %s\n",
			"NAME", "URL", "STATE", "LAST CHECKED")
		fmt.Println(strings.Repeat("-", 12), strings.Repeat("-", 30),
			strings.Repeat("-", 6), strings.Repeat("-", 20))

		for _, e := range eps {
			state := "DOWN"
			if e.Healthy {
				state = "UP"
			}
			fmt.Printf("%-12s  %-30s  %-6s  %s\n",
				e.Name, e.URL, state, e.LastChecked.Format(time.RFC822))
		}
	},
}
