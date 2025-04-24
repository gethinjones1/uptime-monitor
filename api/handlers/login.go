package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	EMAIL    string `json:"email" binding:"required,"`
	PASSWORD string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token": "yeng12345",
	})
}
