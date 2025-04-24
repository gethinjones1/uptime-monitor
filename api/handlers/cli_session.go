package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CLISession struct {
	Token     string
	CreatedAt time.Time
}

var (
	sessionStore = make(map[string]*CLISession)
	storeLock    sync.Mutex
)

// POST /cli/session
func CreateCLISession(c *gin.Context) {
	id := uuid.New().String()

	storeLock.Lock()
	sessionStore[id] = &CLISession{CreatedAt: time.Now()}
	storeLock.Unlock()

	loginURL := "http://localhost:5173/cli-login?session=" + id

	c.JSON(http.StatusOK, gin.H{
		"session_id": id,
		"login_url":  loginURL,
	})
}

// GET /cli/session/:id/status
func GetCLISessionStatus(c *gin.Context) {
	id := c.Param("id")

	storeLock.Lock()
	session, exists := sessionStore[id]
	storeLock.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	if session.Token == "" {
		c.JSON(http.StatusOK, gin.H{"status": "pending"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "complete", "token": session.Token})
	}
}

// POST /cli/session/:id/complete
func CompleteCLISession(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		Token string `json:"token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}

	storeLock.Lock()
	session, exists := sessionStore[id]
	if exists {
		session.Token = req.Token
	}
	storeLock.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "completed"})
}
