package routes

import (
	"uptime-monitor/api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/urls", handlers.CreateMonitoredURL)
	r.GET("/urls", handlers.ListMonitoredURLs)
	r.GET("/status", handlers.GetStatuses)

	return r
}
