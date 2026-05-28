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

func AddPlayerToTeam(teamID string, req models.AddPlayersToTeamRequest) error {
	if len(req.PlayerIDs) <= 0 {
		return errors.New("player_id is required")
	}
	teamExists, err := dbHelper.TeamExists(teamID)
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
		err = dbHelper.AddPlayerToTeam(teamID, playerId)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTeamPlayers(teamID string) ([]models.UserResponse, error) {
	return dbHelper.GetTeamPlayers(teamID)
}

func GetTeams() ([]models.Team, error) {

	return dbHelper.GetTeams()
}
