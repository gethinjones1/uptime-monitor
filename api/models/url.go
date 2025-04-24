package models

import "time"

type MonitoredURL struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	NAME      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
