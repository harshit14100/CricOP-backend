package dbHelper

import (
	"backend/database"
	"backend/models"
	"context"
)

func GetInningsState(
	c context.Context,
	inningsID string,
) (*models.InningsState, error) {

	var innings models.InningsState

	query := `
SELECT
i.status,
i.match_id,
i.batting_team_id,
i.bowling_team_id,
i.total_wickets,
m.allow_solo_batting
FROM innings i
JOIN matches m
ON m.id = i.match_id
WHERE i.id = $1
`

	err := database.DB.QueryRow(
		c,
		query,
		inningsID,
	).Scan(
		&innings.Status,
		&innings.MatchID,
		&innings.BattingTeamID,
		&innings.BowlingTeamID,
		&innings.TotalWickets,
		&innings.AllowSoloBatting,
	)

	if err != nil {
		return nil, err
	}

	return &innings, nil
}

func GetActivePlayersCount(
	c context.Context,
	matchID string,
	teamID string,
) (int, error) {

	var count int

	query := `
SELECT COUNT(*)
FROM match_players
WHERE match_id = $1
AND team_id = $2
AND is_retired = FALSE
`

	err := database.DB.QueryRow( // use get select and excs
		c,
		query,
		matchID,
		teamID,
	).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
