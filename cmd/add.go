package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [name] [url]",
	Short: "Register a new endpoint with the uptime service",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := getAuthToken()
		if err != nil {
			log.Fatalf("auth failed: %v", err)
		}

		reader := bufio.NewReader(os.Stdin)

		// --- prompt for missing inputs ---
		var name, url string
		if len(args) > 0 {
			name = args[0]
		} else {
			fmt.Print("👉 name: ")
			name, _ = reader.ReadString('\n')
			name = strings.TrimSpace(name)
		}
		if len(args) > 1 {
			url = args[1]
		} else {
			fmt.Print("👉 url: ")
			url, _ = reader.ReadString('\n')
			url = strings.TrimSpace(url)
		}

		// --- POST to the API ---
		payload, _ := json.Marshal(map[string]string{
			"name": name,
			"url":  url,
		})
		endpoint := strings.TrimRight(apiURL, "/") + "/urls"
		resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			log.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			body, _ := io.ReadAll(resp.Body)
			log.Fatalf("API error (%d): %s", resp.StatusCode, strings.TrimSpace(string(body)))
		}

		fmt.Printf("✅ Added %q → %q\n", name, url)
	},
}
