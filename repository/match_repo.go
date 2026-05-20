package repository

import (
	"backend/models"
	"context"

	"backend/database"

	"github.com/google/uuid"
)

func CreateMatch(
	req models.CreateMatchRequest,
	hostUserID string,
) error {

	query := `
	INSERT INTO matches (
		id,
		host_user_id,
		team1_id,
		team2_id,
		venue,
		overs,
		players_per_team
	)
	VALUES (
		$1,$2,$3,$4,$5,$6,$7
	)
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		uuid.New(),
		hostUserID,
		req.Team1Id,
		req.Team2Id,
		req.Venue,
		req.Overs,
		req.Players_per_team,
	)

	return err
}
