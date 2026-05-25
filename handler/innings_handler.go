package handler

import (
	"net/http"

	"backend/models"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func StartInning(c *gin.Context) {
	matchID := c.Param("id")

	var req models.StartInningRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	inningID, err := services.StartInning(matchID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "inning started successfully, match is now live",
		"inning_id": inningID,
	})
}
