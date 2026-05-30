package dbHelper

import (
	"context"

	"backend/database"
	"backend/models"

	"github.com/google/uuid"
)

func CreateMatch(
	req models.CreateMatchRequest,
	hostUserID string,
) (string, error) {

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
		req.Team1ID,
		req.Team2ID,
		req.Venue,
		req.Overs,
		req.Players_per_team,
	)
	return "", err
}

func StartMatchToss(
	matchID uuid.UUID,
	tossWinnerID uuid.UUID,
	tossDecision string,
) error {

	query := `
	UPDATE matches
	SET
		toss_winner_id = $1,
		toss_decision = $2,
		status = 'live',
		started_at = NOW()
	WHERE id = $3
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		tossWinnerID,
		tossDecision,
		matchID,
	)

	return err
}

func StartMatch(
	req models.StartMatchRequest,
	hostUserID string,
) (string, error) {

	matchID := uuid.New()

	query := `
	INSERT INTO matches (
		id,
		host_user_id,
		team1_id,
		team2_id,
		venue,
		overs,
		toss_winner_id,
		toss_decision,
		batting_team_id,
		bowling_team_id,
		striker_id,
		non_striker_id,
		current_bowler_id,
		status,
		started_at
	)
	VALUES (
		$1,$2,$3,$4,$5,$6,
		$7,$8,$9,$10,$11,$12,$13,
		'live',
		NOW()
	)
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		matchID,
		hostUserID,
		req.Team1ID,
		req.Team2ID,
		req.Venue,
		req.Overs,
		req.TossWinnerID,
		req.TossDecision,
		req.BattingTeamID,
		req.BowlingTeamID,
		req.StrikerID,
		req.NonStrikerID,
		req.CurrentBowlerID,
	)

	if err != nil {
		return "", err
	}

	return matchID.String(), nil
}

func SuperSetupMatch(
	req models.SuperStartMatchRequest,
	hostUserID string,
) (string, error) {
	tx, err := database.DB.Begin(context.Background())
	if err != nil {
		return "", err
	}
	defer tx.Rollback(context.Background())

	matchID := uuid.New()

	matchQuery := `
	INSERT INTO matches (
		id, host_user_id, team1_id, team2_id, venue, overs,
		toss_winner_id, toss_decision, batting_team_id, bowling_team_id,
		striker_id, non_striker_id, current_bowler_id, status, started_at
	)
	VALUES (
		$1,$2,$3,$4,$5,$6,
		$7,$8,$9,$10,$11,$12,$13,
		'live', NOW()
	)
	`
	_, err = tx.Exec(context.Background(), matchQuery,
		matchID, hostUserID, req.Team1ID, req.Team2ID, req.Venue, req.Overs,
		req.TossWinnerID, req.TossDecision, req.BattingTeamID, req.BowlingTeamID,
		req.StrikerID, req.NonStrikerID, req.CurrentBowlerID,
	)
	if err != nil {
		return "", err
	}

	for _, playerID := range req.Team1Players {
		_, err = tx.Exec(context.Background(),
			`INSERT INTO team_players (team_id, player_id) VALUES ($1,$2) ON CONFLICT (team_id, player_id) DO NOTHING`,
			req.Team1ID, playerID,
		)
		if err != nil {
			return "", err
		}
	}

	for _, playerID := range req.Team2Players {
		_, err = tx.Exec(context.Background(),
			`INSERT INTO team_players (team_id, player_id) VALUES ($1,$2) ON CONFLICT (team_id, player_id) DO NOTHING`,
			req.Team2ID, playerID,
		)
		if err != nil {
			return "", err
		}
	}

	inningID := uuid.New().String()
	inningQuery := `
		INSERT INTO innings (id, match_id, inning_number, batting_team_id, bowling_team_id)
		VALUES ($1, $2, 1, $3, $4)
	`
	_, err = tx.Exec(context.Background(), inningQuery, inningID, matchID, req.BattingTeamID, req.BowlingTeamID)
	if err != nil {
		return "", err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return "", err
	}
	return matchID.String(), nil
}

func GetLiveMatchState(ctx context.Context, matchID string) (*models.LiveMatchStateResponse, error) {
	query := `
		SELECT 
			m.id AS match_id,
			COALESCE(m.venue, '') AS venue,
			m.overs,
			m.status,
			t1.name AS batting_team_name,
			t2.name AS bowling_team_name,
			COALESCE(u1.name, 'Yet to Bat') AS striker_name,
			COALESCE(u2.name, 'Yet to Bat') AS non_striker_name,
			COALESCE(u3.name, 'Yet to Bowl') AS current_bowler_name,
			COALESCE(i.inning_number, 1) AS inning_number,
			COALESCE(i.total_runs, 0) AS total_runs,
			COALESCE(i.wickets, 0) AS wickets,
			COALESCE(i.completed_overs, 0) AS completed_overs,
			COALESCE(i.balls_in_current_over, 0) AS balls_in_current_over,
			COALESCE(i.overs, 0.0) AS display_overs
		FROM matches m
		LEFT JOIN teams t1 ON m.batting_team_id = t1.team_id
		LEFT JOIN teams t2 ON m.bowling_team_id = t2.team_id
		LEFT JOIN users u1 ON m.striker_id = u1.id
		LEFT JOIN users u2 ON m.non_striker_id = u2.id
		LEFT JOIN users u3 ON m.current_bowler_id = u3.id
		LEFT JOIN innings i ON m.id = i.match_id AND i.inning_number = (
			SELECT MAX(inning_number) FROM innings WHERE match_id = m.id
		)
		WHERE m.id = $1
	`

	var resp models.LiveMatchStateResponse
	err := database.DB.QueryRow(ctx, query, matchID).Scan( // fix heading
		&resp.MatchID,
		&resp.Venue,
		&resp.Overs,
		&resp.Status,
		&resp.BattingTeamName,
		&resp.BowlingTeamName,
		&resp.StrikerName,
		&resp.NonStrikerName,
		&resp.CurrentBowlerName,
		&resp.InningNumber,
		&resp.TotalRuns,
		&resp.Wickets,
		&resp.CompletedOvers,
		&resp.BallsInCurrentOver,
		&resp.DisplayOvers,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetMatches() ([]models.MatchListResponse, error) {

	query := `
	SELECT
		m.id AS match_id,
		m.status,
		m.team1_id,
		m.team2_id,
		t1.name,
		t2.name,

		COALESCE(m.venue, ''),
		COALESCE(m.overs, 0),
		COALESCE(i.inning_number, 1),
		COALESCE(i.total_runs, 0),
		COALESCE(i.wickets, 0),
		i.batting_team_id,
		COALESCE(bt.name, '')

	FROM matches m

	JOIN teams t1
		ON m.team1_id = t1.team_id

	JOIN teams t2
		ON m.team2_id = t2.team_id

	LEFT JOIN innings i
		ON i.match_id = m.id

	LEFT JOIN teams bt
		ON bt.team_id = i.batting_team_id

	ORDER BY m.created_at DESC
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []models.MatchListResponse
	for rows.Next() {
		var match models.MatchListResponse
		err := rows.Scan(
			&match.ID,
			&match.Status,
			&match.Team1ID,
			&match.Team2ID,
			&match.Team1Name,
			&match.Team2Name,
			&match.Venue,
			&match.Overs,
			&match.CurrentInnings,
			&match.TotalRuns,
			&match.Wickets,
			&match.BattingTeamID,
			&match.BattingTeamName,
		)

		if err != nil {
			return nil, err
		}

		matches = append(matches, match)
	}
	return matches, nil
}
