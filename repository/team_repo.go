package repository

import (
	"backend/database"
	"backend/models"
	"context"

	"github.com/google/uuid"
)

func TeamExists(teamId string) (bool, error) {

	query := `
	SELECT EXISTS(
		SELECT 1
		FROM teams
		WHERE team_id = $1
		AND archived_at IS NULL
	)
	`

	var exists bool

	err := database.DB.QueryRow(
		context.Background(),
		query,
		teamId,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func PlayerExists(playerId string) (bool, error) {
	query := `
	SELECT EXISTS(
	SELECT 1
	FROM players
	WHERE id = $1
	)`

	var exists bool

	err := database.DB.QueryRow(
		context.Background(),
		query,
		playerId,
	).Scan(&exists)

	if err != nil {
		return false, err
	}
	return exists, nil
}

func AddPlayerToTeam(teamId string, playerID string) error {
	query := `
INSERT INTO team_players(id, team_id, player_id)
VALUES ($1, $2, $3)
on conflict (team_id, player_id) do nothing`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		uuid.New(),
		teamId,
		playerID,
	)

	return err
}

func CreateTeam(
	req models.CreateTeamRequest,
	hostUserID string,
) error {

	query := `
	INSERT INTO teams (
		team_id,
	    name,
	    created_by
	)
	VALUES (
		$1, $2, $3
	)
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		uuid.New(),
		req.Name,
		hostUserID,
	)

	return err
}
