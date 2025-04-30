package handlers

import (
	"database/sql"
	"net/http"
	"time"
	"uptime-monitor/shared/database"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type SignupResponse struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func Signup(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email already exists
	var existingID int
	err := database.DB.QueryRow("SELECT id FROM users WHERE email = $1", req.Email).Scan(&existingID)
	switch {
	case err == nil:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already in use"})
		return
	case err != sql.ErrNoRows:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email: " + err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Insert new user
	var createdUser SignupResponse
	query := `
		INSERT INTO users (first_name, last_name, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`
	err = database.DB.QueryRow(query, req.FirstName, req.LastName, req.Email, string(hashedPassword)).
		Scan(&createdUser.ID, &createdUser.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database insert failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
