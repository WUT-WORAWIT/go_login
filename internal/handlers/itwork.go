package handlers

import (
	"go_login/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetItWork ... Health check endpoint (query work table)
func GetItWork(c *gin.Context) {
	resultChan := make(chan []models.Work)
	errChan := make(chan error)

	go func() {
		var work []models.Work
		err := models.GetItWorkall(&work)
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- work
	}()

	// Wait for result
	select {
	case works := <-resultChan:
		c.JSON(http.StatusOK, works)
	case err := <-errChan:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
