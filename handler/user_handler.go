package handler

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "protected profile route working",
	})
}

func GetUserByUsername(c *gin.Context) {

	username := c.Param("username")

	c.JSON(200, gin.H{
		"username": username,
	})
}
func GetAllUsers(c *gin.Context) {

	users, err := services.GetALlUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
