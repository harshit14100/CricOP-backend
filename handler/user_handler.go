package handler

import (
	"backend/models"
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

func ResetPassword(c *gin.Context) {
	var req models.PasswordRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = services.ResetPassword(
		req.PhoneNo,
		req.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "password reset successfully",
	})
}
