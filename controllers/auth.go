package controllers

import (
	"api-gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginHandler is the handler for the login route.
func LoginHandler(c *gin.Context) {
	var credentials models.LoginRequest
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Simulate authentication.
	if credentials.Username == "admin" && credentials.Password == "1234" {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "username": credentials.Username})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password is 1234 don't be noob"})
	}
}
