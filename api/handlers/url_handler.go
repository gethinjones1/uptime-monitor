package handlers

import (
	"net/http"
	"uptime-monitor/api/models"
	"uptime-monitor/shared/database"

	"github.com/gin-gonic/gin"
)

type CreateURLRequest struct {
	URL  string `json:"url" binding:"required,url"`
	NAME string `json:"name" binding:"required"`
}

func CreateMonitoredURL(c *gin.Context) {
	var req CreateURLRequest

	// Validate input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// Insert into DB
	query := `INSERT INTO monitored_urls (url, name) VALUES ($1, $2) RETURNING id, created_at`
	var url models.MonitoredURL
	url.URL = req.URL
	url.NAME = req.NAME

	err := database.DB.QueryRow(query, url.URL, url.NAME).Scan(&url.ID, &url.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database insert failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, url)
}

func ListMonitoredURLs(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, url, created_at FROM monitored_urls ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed: " + err.Error()})
		return
	}
	defer rows.Close()

	var urls []models.MonitoredURL
	for rows.Next() {
		var u models.MonitoredURL
		if err := rows.Scan(&u.ID, &u.URL, &u.CreatedAt); err == nil {
			urls = append(urls, u)
		}
	}

	c.JSON(http.StatusOK, urls)
}
