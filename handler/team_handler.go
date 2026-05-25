package handler

import (
	"backend/database"
	"backend/database/dbHelper"
	"context"
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

func GetTeam(c *gin.Context) {
	teamID := c.Param("team_id")
	team, err := dbHelper.GetTeam(teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "team not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"team": team,
	})
}

func GetTeamPlayers(teamID *gin.Context) {

	query := `
	SELECT 
		u.id,
		u.name,
		u.phone_no
	FROM team_players tp
	INNER JOIN users u
	ON tp.player_id = u.id
	WHERE tp.team_id = $1
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		teamID,
	)

	if err != nil {
		return
	}

	defer rows.Close()

	var players []models.UserResponse

	for rows.Next() {

		var player models.UserResponse

		err := rows.Scan(
			&player.ID,
			&player.Name,
			&player.PhoneNo,
		)

		if err != nil {
			return
		}

		players = append(players, player)
	}

	return
}
