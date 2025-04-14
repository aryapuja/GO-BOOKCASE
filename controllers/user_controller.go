package controllers

import (
	"go-bookcase/models"
	"go-bookcase/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	var credentials models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Authenticate user and get JWT token
	token, err := services.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the JWT token in response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
