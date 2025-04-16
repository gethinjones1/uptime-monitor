package main

import (
	"log"
	"os"
	"uptime-monitor/api/routes"
	"uptime-monitor/shared/database"
)

func main() {
	database.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := routes.SetupRouter()

	log.Printf("Starting API on port %s...\n", port)
	r.Run(":" + port)
}
