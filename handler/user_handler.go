package handler

import (
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
