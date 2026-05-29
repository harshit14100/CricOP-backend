package handler

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLiveMatchState(c *gin.Context) {

	matchID := c.Param("id")

	data, err := services.GetLiveMatch(matchID)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
