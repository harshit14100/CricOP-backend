package dbHelper

import (
	"backend/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func InsertDelivery(ctx context.Context, tx pgx.Tx, inningID string, req models.RecordDeliveryRequest, totalRuns int) error {
	query := `
		INSERT INTO deliveries (
			id, inning_id, over_number, ball_number, striker_id, non_striker_id, bowler_id, 
			runs_bat, extras, extra_type, total_runs, wicket, wicket_type, fielder_id, player_out_id, is_free_hit
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	`
	_, err := tx.Exec(ctx, query,
		uuid.New(), inningID, req.OverNumber, req.BallNumber, req.StrikerID, req.NonStrikerID, req.BowlerID,
		req.RunsBat, req.Extras, req.ExtraType, totalRuns, req.Wicket, req.WicketType, req.FielderID, req.PlayerOutID, req.IsFreeHit,
	)
	return err
}

func UpdateInningStats(ctx context.Context, tx pgx.Tx, inningID string, totalRuns int, extras int, isLegalDelivery bool, isWicket bool) error {
	query := `
		UPDATE innings 
		SET total_runs = total_runs + $1,
		    extras = extras + $2,
		    wickets = wickets + $3,
		    balls_in_current_over = balls_in_current_over + $4
		WHERE id = $5
	`

	wicketCount := 0
	if isWicket {
		wicketCount = 1
	}

	ballCount := 0
	if isLegalDelivery {
		ballCount = 1
	}

	_, err := tx.Exec(ctx, query, totalRuns, extras, wicketCount, ballCount, inningID)
	return err
}

func UpdateBattingScorecard(ctx context.Context, tx pgx.Tx, inningID string, playerID string, runs int, isLegalDelivery bool, isWicket bool, dismissalType *string) error {
	query := `
		INSERT INTO batting_scorecards (id, inning_id, player_id, runs, balls_faced, fours, sixes, is_out, dismissal_type)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (inning_id, player_id) DO UPDATE SET
			runs = batting_scorecards.runs + EXCLUDED.runs,
			balls_faced = batting_scorecards.balls_faced + EXCLUDED.balls_faced,
			fours = batting_scorecards.fours + EXCLUDED.fours,
			sixes = batting_scorecards.sixes + EXCLUDED.sixes,
			is_out = EXCLUDED.is_out,
			dismissal_type = EXCLUDED.dismissal_type
	`

	ballsFaced, fours, sixes := 0, 0, 0
	if isLegalDelivery {
		ballsFaced = 1
	}
	if runs == 4 {
		fours = 1
	}
	if runs == 6 {
		sixes = 1
	}

	_, err := tx.Exec(ctx, query, uuid.New(), inningID, playerID, runs, ballsFaced, fours, sixes, isWicket, dismissalType)
	return err
}

func UpdateBowlingScorecard(ctx context.Context, tx pgx.Tx, inningID string, bowlerID string, totalRuns int, isLegalDelivery bool, isWicket bool) error {
	query := `
		INSERT INTO bowling_scorecards (id, inning_id, player_id, runs_conceded, wickets)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (inning_id, player_id) DO UPDATE SET
			runs_conceded = bowling_scorecards.runs_conceded + EXCLUDED.runs_conceded,
			wickets = bowling_scorecards.wickets + EXCLUDED.wickets
	`
	wicketCount := 0
	if isWicket {
		wicketCount = 1
	}

	_, err := tx.Exec(ctx, query, uuid.New(), inningID, bowlerID, totalRuns, wicketCount)
	return err
}
