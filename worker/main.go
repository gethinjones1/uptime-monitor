package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	"uptime-monitor/shared/database"
)

type MonitoredURL struct {
	ID  int
	URL string
}

func main() {
	database.Init()
	db := database.DB

	defer db.Close()

	for {
		log.Println("🔍 Checking all monitored URLs...")
		checkAll(db)
		time.Sleep(60 * time.Second)
	}
}

func checkAll(db *sql.DB) {
	rows, err := db.Query("SELECT id, url FROM monitored_urls")
	if err != nil {
		log.Println("❌ Failed to fetch URLs:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u MonitoredURL
		if err := rows.Scan(&u.ID, &u.URL); err == nil {
			checkAndLog(db, u)
		}
	}
}

func checkAndLog(db *sql.DB, u MonitoredURL) {
	client := http.Client{Timeout: 10 * time.Second}
	start := time.Now() // Start timer
	resp, err := client.Get(u.URL)

	statusCode := 0
	isUp := false

	if err == nil {
		statusCode = resp.StatusCode
		isUp = statusCode < 500
		resp.Body.Close()
	}
	duration := time.Since(start).Milliseconds() // Calculate duration

	_, err = db.Exec(`INSERT INTO url_statuses (url_id, status_code, is_up, response_time) VALUES ($1, $2, $3, $4)`,
		u.ID, statusCode, isUp, duration,
	)

	if err != nil {
		log.Printf("⚠️  Failed to log status for %s: %v\n", u.URL, err)
	} else {
		log.Printf("✅ %s -> %d (up: %t)\n", u.URL, statusCode, isUp)
	}
}
