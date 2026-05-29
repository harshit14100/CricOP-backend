package handler

import (
	"fmt"
	"net/http"

	"backend/models"
	"backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateMatch(c *gin.Context) {

	var req models.CreateMatchRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	hostUserID := userID.(string)

	matchID, err := services.CreateMatch(
		req,
		hostUserID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "match created successfully",
		"id":      matchID,
	})
}

func StartMatchToss(c *gin.Context) {
	matchIDParam := c.Param("id")
	matchID, err := uuid.Parse(matchIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid match id",
		})
		return
	}

	var req models.Toss
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	err = services.StartMatchToss(
		matchID,
		req,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "toss completed successfully",
	})
}

func StartMatch(c *gin.Context) {
	var req models.StartMatchRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	hostUserID := userID.(string)
	matchID, err := services.StartMatch(
		req,
		hostUserID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "match started successfully",
		"match_id": matchID,
	})
}

func SuperSetupMatchHandler(c *gin.Context) {

	var req models.SuperStartMatchRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		fmt.Println("BIND ERROR:", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	userID, exists := c.Get("user_id")

	if !exists {

		fmt.Println("USER ID MISSING")

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})

		return
	}

	hostUserID := userID.(string)

	matchID, err := services.SuperSetupMatch(
		req,
		hostUserID,
	)

	if err != nil {

		fmt.Println("SUPER SETUP ERROR:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "match created successfully",
		"match_id": matchID,
	})
}
