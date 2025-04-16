package models

import "time"

type MonitoredURL struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
