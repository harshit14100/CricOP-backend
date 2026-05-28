package dbHelper

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
	FROM users
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

func GetTeamPlayers(teamID string) ([]models.UserResponse, error) {

	query := `
	SELECT 
		u.id,
		u.name,
		u.phone_no
	FROM team_players tp
	INNER JOIN users u
	ON tp.player_id = u.id
	WHERE tp.team_id = $1
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		teamID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var players []models.UserResponse

	for rows.Next() {

		var player models.UserResponse

		err := rows.Scan(
			&player.ID,
			&player.Name,
			&player.PhoneNo,
		)

		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}
func GetTeam(teamID string) (*models.Team, error) {
	var team models.Team

	query := `
		SELECT team_id, name, created_by, created_at, updated_at
		FROM teams
		WHERE team_id = $1
	`

	err := database.DB.QueryRow(
		context.Background(),
		query,
		teamID,
	).Scan(
		&team.ID,
		&team.Name,
		&team.CreatedBy,
		&team.CreatedAt,
		&team.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &team, nil
}

func GetTeams() ([]models.Team, error) {

	query := `
	SELECT
		team_id,
		name,
		created_at
	FROM teams
	WHERE archived_at IS NULL
	ORDER BY created_at DESC
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var teams []models.Team

	for rows.Next() {

		var team models.Team

		err := rows.Scan(
			&team.ID,
			&team.Name,
			&team.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}

	if teams == nil {
		teams = []models.Team{}
	}

	return teams, nil
}
