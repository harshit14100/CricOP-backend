package services

import (
	"backend/models"
	"backend/repository"
	"errors"
)

func CreateMatch(req models.CreateMatchRequest, hostuserId string) error {
	if req.Team1Id == req.Team2Id {
		return errors.New("Team1Id and Team2Id cannot be the same")
	}
	if req.Players_per_team <= 0 {
		return errors.New("Players per team must be positive")

	}
	return repository.CreateMatch(req, hostuserId)

}
