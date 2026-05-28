package dbHelper

import (
	"backend/database"
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
JOIN matches m ON m.id = i.match_id
WHERE i.id = $1
`

	err := database.DB.GetContext(
		c,
		&innings,
		query,
		inningsID,
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

	err := database.DB.GetContext(
		c,
		&count,
		query,
		matchID,
		teamID,
	)

	return count, err
}
