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
	resp, err := client.Get(u.URL)

	statusCode := 0
	isUp := false

	if err == nil {
		statusCode = resp.StatusCode
		isUp = statusCode < 500
		resp.Body.Close()
	}

	_, err = db.Exec(`INSERT INTO url_statuses (url_id, status_code, is_up) VALUES ($1, $2, $3)`,
		u.ID, statusCode, isUp,
	)

	if err != nil {
		log.Printf("⚠️  Failed to log status for %s: %v\n", u.URL, err)
	} else {
		log.Printf("✅ %s -> %d (up: %t)\n", u.URL, statusCode, isUp)
	}
}
