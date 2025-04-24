package routes

import (
	"uptime-monitor/api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/login", handlers.Login)

	r.POST("/cli/session", handlers.CreateCLISession)
	r.GET("/cli/session/:id/status", handlers.GetCLISessionStatus)
	r.POST("/cli/session/:id/complete", handlers.CompleteCLISession)

	r.POST("/urls", handlers.CreateMonitoredURL)
	r.GET("/urls", handlers.ListMonitoredURLs)
	r.GET("/status", handlers.GetStatuses)

	return r
}
