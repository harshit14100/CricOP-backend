package services

import (
	"backend/database/dbHelper"
	"backend/models"
	"errors"
)

func CreateMatch(req models.CreateMatchRequest, hostuserId string) error {
	if req.Team1ID == req.Team2ID {
		return errors.New("Team1Id and Team2Id cannot be the same")
	}
	if req.Players_per_team <= 0 {
		return errors.New("Players per team must be positive")

	}
	return dbHelper.CreateMatch(req, hostuserId)

}
