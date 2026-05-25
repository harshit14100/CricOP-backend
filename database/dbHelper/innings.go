package dbHelper

import (
	"backend/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateMatchStatusLive(ctx context.Context, tx pgx.Tx, matchID string) error {
	query := `UPDATE matches SET status = 'live' WHERE id = $1`
	_, err := tx.Exec(ctx, query, matchID)
	return err
}

func CreateInning(ctx context.Context, tx pgx.Tx, inningID string, matchID string, req models.StartInningRequest) error {
	query := `
		INSERT INTO innings (id, match_id, inning_number, batting_team_id, bowling_team_id)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := tx.Exec(
		ctx,
		query,
		inningID,
		matchID,
		req.InningNumber,
		req.BattingTeamID,
		req.BowlingTeamID,
	)
	return err
}
