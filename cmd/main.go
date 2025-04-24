package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "uptime-monitor",
		Short: "CLI for your central uptime-monitor service",
	}

	apiURL    string
	tokenPath string
)

func getAuthToken() (string, error) {
	data, err := os.ReadFile(tokenPath)
	if err == nil {
		return string(data), nil
	}

	fmt.Print("No auth token found. Would you like to login now? [Y/n]: ")
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = answer[:len(answer)-1]

	if answer == "n" || answer == "N" {
		return "", errors.New("authentication required. Please run 'uptime-monitor login'")
	}

	// Trigger login flow
	if err := loginCmd.RunE(nil, nil); err != nil {
		return "", fmt.Errorf("login failed: %w", err)
	}

	// Retry reading token after login
	data, err = os.ReadFile(tokenPath)
	if err != nil {
		return "", errors.New("failed to read token after login")
	}
	return string(data), nil
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to detect home directory: %v", err)
	}
	tokenPath = filepath.Join(home, ".uptime-monitor-token")

	defaultURL := os.Getenv("UPTIME_API_URL")
	if defaultURL == "" {
		defaultURL = "http://localhost:8080"
	}

	rootCmd.PersistentFlags().StringVar(&apiURL, "api-url", defaultURL,
		"Base URL of uptime API (or set UPTIME_API_URL)")

	// Register commands
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(logoutCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
