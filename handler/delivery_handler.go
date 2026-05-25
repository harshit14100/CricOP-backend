package handler

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDelivery(c *gin.Context) {
	inningID := c.Param("inning_id")

	var req models.RecordDeliveryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := services.RecordDelivery(inningID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record delivery: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Delivery recorded successfully"})
}
