package services

import (
	"backend/models"
	"backend/repository"
	"errors"
)

func CreateTeam(req models.CreateTeamRequest, createdBy string) error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	return repository.CreateTeam(req, createdBy)

}

func AddPlayerToTeam(teamId string, req models.AddPlayersToTeamRequest) error {
	if len(req.PlayerIDs) <= 0 {
		return errors.New("player_id is required")
	}
	teamExists, err := repository.TeamExists(teamId)
	if err != nil {
		return err
	}
	if !teamExists {
		return errors.New("team does not exist")
	}

	for _, playerId := range req.PlayerIDs {
		playerExists, err := repository.PlayerExists(playerId)
		if err != nil {
			return err
		}
		if !playerExists {
			return errors.New("player does not exist")
		}
		err = repository.AddPlayerToTeam(teamId, playerId)
		if err != nil {
			return err
		}
	}
	return nil
}
