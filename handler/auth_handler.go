package handler

import (
	dbHelper "backend/database/dbhelper"
	"net/http"

	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Name     string `json:"name"`
	PhoneNo  string `json:"phone_no"`
	Password string `json:"password"`
}

type LoginRequest struct {
	PhoneNo  string `json:"phone_no"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {

	var req SignupRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	hashedPassword, err := utils.Hashpassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := models.Users{
		Name:     req.Name,
		PhoneNo:  req.PhoneNo,
		Password: hashedPassword,
	}

	err = dbHelper.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.ID.String())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": gin.H{
			"name":     user.Name,
			"phone_no": user.PhoneNo,
		},
		"token": token,
	})
}

func Login(c *gin.Context) {

	var req LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user, err := dbHelper.GetUserByPhone(req.PhoneNo)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	isValid := utils.VerifyPassword(
		req.Password,
		user.Password,
	)

	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID.String())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       user.ID,
			"name":     user.Name,
			"phone_no": user.PhoneNo,
		},
		"token": token,
	})
}
