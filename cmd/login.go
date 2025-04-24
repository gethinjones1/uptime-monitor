package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate via browser and store CLI token",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. Request session from API
		resp, err := http.Post(apiURL+"/cli/session", "application/json", nil)
		if err != nil {
			return fmt.Errorf("failed to create session: %w", err)
		}
		defer resp.Body.Close()

		var session struct {
			SessionID string `json:"session_id"`
			LoginURL  string `json:"login_url"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
			return err
		}

		fmt.Println("Opening browser for login...")
		if err := browser.OpenURL(session.LoginURL); err != nil {
			fmt.Printf("Please open the URL manually: %s\n", session.LoginURL)
		}

		// 2. Poll for token
		fmt.Println("Waiting for authentication...")

		for {
			time.Sleep(3 * time.Second)

			statusResp, err := http.Get(apiURL + "/cli/session/" + session.SessionID + "/status")
			if err != nil {
				return err
			}
			defer statusResp.Body.Close()

			var status struct {
				Status string `json:"status"`
				Token  string `json:"token"`
			}
			if err := json.NewDecoder(statusResp.Body).Decode(&status); err != nil {
				return err
			}

			if status.Status == "complete" {
				if err := os.WriteFile(tokenPath, []byte(status.Token), 0600); err != nil {
					return err
				}
				fmt.Println("✅ Login successful! Token saved.")
				break
			}

			fmt.Print(".") // progress indicator
		}

		return nil
	},
}
