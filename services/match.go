package services

import (
	"backend/database/dbHelper"
	"backend/models"
	"errors"

	"github.com/google/uuid"
)

func CreateMatch(req models.CreateMatchRequest, hostuserId string) (string, error) {
	if req.Team1ID == req.Team2ID {
		return "", errors.New("Team1Id and Team2Id cannot be the same")
	}

	if req.Players_per_team <= 0 {
		return "", errors.New("Players per team must be positive")
	}

	matchID, err := dbHelper.CreateMatch(req, hostuserId)
	if err != nil {
		return "", err
	}

	return matchID, nil
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

//func GetAllMatches() ([]models.Match, error) {
//	return dbHelper.GetMatches()
//}

func StartMatch(
	req models.StartMatchRequest,
	hostUserID string,
) (string, error) {

	return dbHelper.StartMatch(req, hostUserID)
}

func SuperSetupMatch(
	req models.SuperStartMatchRequest,
	hostUserID string,
) (string, error) {

	return dbHelper.SuperSetupMatch(
		req,
		hostUserID,
	)
}

func GetMatches() ([]models.MatchListResponse, error) {

	matches, err := dbHelper.GetMatches()

	if err != nil {
		return nil, err
	}

	return matches, nil
}
