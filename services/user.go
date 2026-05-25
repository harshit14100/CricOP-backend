package services

import (
	"backend/database/dbHelper"
	"backend/models"

	"golang.org/x/crypto/bcrypt"
)

func ResetPassword(phoneNo string, newPassword string) error {
	hashedpassword, err := bcrypt.GenerateFromPassword(
		[]byte(newPassword), bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	return dbHelper.ResetPassword(phoneNo, string(hashedpassword))
}

func GetALlUsers() ([]models.UserResponse, error) {
	return dbHelper.GetAllUsers()
}
