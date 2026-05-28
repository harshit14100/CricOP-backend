package dbHelper

import (
	"backend/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func InsertDelivery(
	ctx context.Context,
	tx pgx.Tx,
	delivery models.DeliveryRecord,
) error {

	query := `
INSERT INTO deliveries (
id,
inning_id,
over_number,
ball_number,
striker_id,
non_striker_id,
bowler_id,
runs_bat,
extras,
extra_type,
total_runs,
wicket,
wicket_type,
fielder_id,
player_out_id,
is_free_hit
)
VALUES (
$1, $2, $3, $4, $5, $6, $7,
$8, $9, $10, $11, $12, $13,
$14, $15, $16
)
`

	_, err := tx.Exec(
		ctx,
		query,
		uuid.New(),
		delivery.InningID,
		delivery.OverNumber,
		delivery.BallNumber,
		delivery.StrikerID,
		delivery.NonStrikerID,
		delivery.BowlerID,
		delivery.RunsBat,
		delivery.Extras,
		delivery.ExtraType,
		delivery.TotalRuns,
		delivery.Wicket,
		delivery.WicketType,
		delivery.FielderID,
		delivery.PlayerOutID,
		delivery.IsFreeHit,
	)

	return err
}

func UpdateInningStats(
	ctx context.Context,
	tx pgx.Tx,
	inningID string,
	totalRuns int,
	extras int,
	isLegalDelivery bool,
	isWicket bool,
) error {

	query := `
UPDATE innings
SET total_runs = total_runs + $1,
extras = extras + $2,
wickets = wickets + $3,
updated_at = NOW()
WHERE id = $4
`

	wicketCount := 0
	if isWicket {
		wicketCount = 1
	}

	legalBallIncrement := 0
	if isLegalDelivery {
		legalBallIncrement = 1
	}

	_, err := tx.Exec(
		ctx,
		query,
		totalRuns,
		extras,
		wicketCount,
		legalBallIncrement,
		inningID,
	)

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

func UpdateBowlingScorecard(
	ctx context.Context,
	tx pgx.Tx,
	inningID string,
	bowlerID string,
	totalRuns int,
	isLegalDelivery bool,
	isWicket bool,
	wicketType *string,
) error {

	query := `
INSERT INTO bowling_scorecards (
id,
inning_id,
player_id,
runs_conceded,
balls_bowled,
wickets
)
VALUES ($1, $2, $3, $4, $5, $6)

ON CONFLICT (inning_id, player_id)
DO UPDATE SET
runs_conceded =
bowling_scorecards.runs_conceded +
EXCLUDED.runs_conceded,

balls_bowled =
bowling_scorecards.balls_bowled +
EXCLUDED.balls_bowled,

wickets =
bowling_scorecards.wickets +
EXCLUDED.wickets
`

	wicketCount := 0

	if isWicket &&
		wicketType != nil &&
		*wicketType != "run_out" &&
		*wicketType != "retired_hurt" {

		wicketCount = 1
	}

	legalBalls := 0
	if isLegalDelivery {
		legalBalls = 1
	}

	_, err := tx.Exec(
		ctx,
		query,
		uuid.New(),
		inningID,
		bowlerID,
		totalRuns,
		legalBalls,
		wicketCount,
	)

	return err
}

func RotateStrike(
	ctx context.Context,
	tx pgx.Tx,
	inningID string,
) error {

	query := `
UPDATE innings
SET striker_id = non_striker_id,
non_striker_id = striker_id
WHERE id = $1
`

	_, err := tx.Exec(
		ctx,
		query,
		inningID,
	)

	return err
}

func CheckAndConcludeMatch(ctx context.Context, tx pgx.Tx, inningID string) (bool, error) {
	query := `
		SELECT 
			i.match_id, i.inning_number, i.total_runs, i.wickets, i.completed_overs, i.balls_in_current_over,
			m.overs AS match_total_overs, m.players_per_team, m.batting_team_id, m.bowling_team_id,
			COALESCE((SELECT total_runs FROM innings WHERE match_id = i.match_id AND inning_number = 1), 0) AS inning1_runs
		FROM innings i
		JOIN matches m ON i.match_id = m.id
		WHERE i.id = $1
	`

	var matchID, battingTeamID, bowlingTeamID string
	var inningNumber, totalRuns, wickets, completedOvers, ballsInCurrentOver, matchTotalOvers, playersPerTeam, inning1Runs int

	err := tx.QueryRow(ctx, query, inningID).Scan(
		&matchID, &inningNumber, &totalRuns, &wickets, &completedOvers, &ballsInCurrentOver,
		&matchTotalOvers, &playersPerTeam, &battingTeamID, &bowlingTeamID, &inning1Runs,
	)
	if err != nil {
		return false, err
	}
	if inningNumber != 2 {
		return false, nil
	}

	target := inning1Runs + 1
	isAllOut := wickets >= (playersPerTeam - 1)
	isOversFinished := completedOvers >= matchTotalOvers && ballsInCurrentOver == 0

	var winnerTeamID *string
	matchEnded := false

	if totalRuns >= target {
		winnerTeamID = &battingTeamID
		matchEnded = true
	} else if isAllOut || isOversFinished {
		matchEnded = true
		if totalRuns < inning1Runs {
			winnerTeamID = &bowlingTeamID
		} else {
			winnerTeamID = nil
		}
	}

	if matchEnded {
		updateQuery := `
			UPDATE matches
			SET status = 'completed',
			    winner_team_id = $1,
			    ended_at = NOW()
			WHERE id = $2
		`
		_, err = tx.Exec(ctx, updateQuery, winnerTeamID, matchID)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
