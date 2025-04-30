package routes

import (
	"time"
	"uptime-monitor/api/handlers"
	"uptime-monitor/api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Or "*" for dev
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", handlers.Login)
	r.POST("/signup", handlers.Signup)

	r.POST("/cli/session", handlers.CreateCLISession)
	r.GET("/cli/session/:id/status", handlers.GetCLISessionStatus)
	r.POST("/cli/session/:id/complete", handlers.CompleteCLISession)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthRequired())
	{
		auth.POST("/urls", handlers.CreateMonitoredURL)
		auth.GET("/urls", handlers.ListMonitoredURLs)
		auth.GET("/status", handlers.GetStatuses)
	}

	return r
}
