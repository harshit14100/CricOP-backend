package handler

import (
	"net/http"

	"backend/models"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func CreateTeams(c *gin.Context) {

	var req models.CreateTeamRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
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

	createdBy := userID.(string)

	err = services.CreateTeam(
		req,
		createdBy,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "team created successfully",
	})
}

func AddPlayerToTeam(c *gin.Context) {

	var req models.AddPlayersToTeamRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	teamID := c.Param("id")

	err = services.AddPlayerToTeam(
		teamID,
		req,
	)

	if err != nil {

		switch err.Error() {

		case "team does not exist":
			c.JSON(http.StatusNotFound, gin.H{
				"message": "team not found",
			})

		case "player does not exist":
			c.JSON(http.StatusNotFound, gin.H{
				"message": "player not found",
			})

		case "player already exists":
			c.JSON(http.StatusConflict, gin.H{
				"message": "player already exists",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "players added successfully",
	})
}
