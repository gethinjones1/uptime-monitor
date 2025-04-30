package handlers

import (
	"net/http"
	"time"
	"uptime-monitor/shared/database"

	"github.com/gin-gonic/gin"
)

type URLStatusResponse struct {
	ID           int       `json:"id"`
	URL          string    `json:"url"`
	NAME         string    `json:"name"`
	StatusCode   int       `json:"status_code"`
	IsUp         bool      `json:"is_up"`
	ResponseTime int       `json:"response_time"`
	CheckedAt    time.Time `json:"checked_at"`
	Availability int       `json:"availability"`
}

func GetStatuses(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID := userIDVal.(int)
	query := `
			SELECT DISTINCT ON (u.id)
			u.id,
			u.url,
			u.name,
			COALESCE(s.status_code, 0),
			COALESCE(s.is_up, false),
			COALESCE(s.checked_at, now()),
			COALESCE(s.response_time, 0),
			COALESCE(av.availability, 0) AS availability
		FROM monitored_urls u
		LEFT JOIN url_statuses s
			ON u.id = s.url_id
		LEFT JOIN (
			SELECT 
				url_id,
				ROUND(100.0 * SUM(CASE WHEN is_up THEN 1 ELSE 0 END) / COUNT(*)) AS availability
			FROM url_statuses
			GROUP BY url_id
		) av
			ON u.id = av.url_id

		where u.user_id = $1
		ORDER BY u.id, s.checked_at DESC
	`

	rows, err := database.DB.Query(query, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch statuses"})
		return
	}
	defer rows.Close()

	var results []URLStatusResponse
	for rows.Next() {
		var r URLStatusResponse
		if err := rows.Scan(&r.ID, &r.URL, &r.NAME, &r.StatusCode, &r.IsUp, &r.CheckedAt, &r.ResponseTime, &r.Availability); err == nil {
			results = append(results, r)
		}
	}

	c.JSON(http.StatusOK, results)
}
