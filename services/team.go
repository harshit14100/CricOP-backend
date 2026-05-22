package services

import (
	"backend/database/dbHelper"
	"backend/models"
	"errors"
)

func CreateTeam(req models.CreateTeamRequest, createdBy string) error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	return dbHelper.CreateTeam(req, createdBy)

}

func AddPlayerToTeam(teamId string, req models.AddPlayersToTeamRequest) error {
	if len(req.PlayerIDs) <= 0 {
		return errors.New("player_id is required")
	}
	teamExists, err := dbHelper.TeamExists(teamId)
	if err != nil {
		return err
	}
	if !teamExists {
		return errors.New("team does not exist")
	}

	for _, playerId := range req.PlayerIDs {
		playerExists, err := dbHelper.PlayerExists(playerId)
		if err != nil {
			return err
		}
		if !playerExists {
			return errors.New("player does not exist")
		}
		err = dbHelper.AddPlayerToTeam(teamId, playerId)
		if err != nil {
			return err
		}
	}
	return nil
}
