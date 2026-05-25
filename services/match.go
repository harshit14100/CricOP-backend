package services

import (
	"backend/database/dbHelper"
	"backend/models"
	"errors"

	"github.com/google/uuid"
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

func StartMatchToss(matchID uuid.UUID,
	req models.Toss,
) error {
	if req.TossDecision != "bat" &&
		req.TossDecision != "bowl" {
		return errors.New("Toss decision must be 'bat' or 'bowl'")
	}
	return dbHelper.StartMatchToss(matchID, req.TossWinnerID, req.TossDecision)
}
