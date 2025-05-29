package handlers

import (
	"go_login/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetProfile(c *gin.Context) {
	// Get claims from middleware
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "No token claims found",
		})
		return
	}

	// Type assert claims
	jwtClaims, ok := claims.(*jwt.RegisteredClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to parse token claims",
		})
		return
	}

	// Get user from database using the username from claims
	var user models.User
	if err := models.GetUserByID(&user, jwtClaims.Subject); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "User not found",
		})
		return
	}

	// Return user profile (excluding sensitive data)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"username": user.Username,
			// Add other non-sensitive user fields here
		},
	})
}
