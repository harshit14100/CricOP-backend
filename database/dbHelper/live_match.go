package dbHelper

import (
	"backend/database"
	"backend/models"
	"context"
)

func GetLiveMatch(
	ctx context.Context,
	matchID string,
) (*models.LiveMatchResponse, error) {

	query := `
SELECT
m.id,
m.team1_id,
m.team2_id,
t1.name,
t2.name,
i.inning_number,
i.total_runs,
i.wickets,
i.completed_overs,
i.balls_in_current_over,
lms.striker_id,
lms.non_striker_id,
lms.bowler_id
FROM matches m

JOIN innings i
ON i.match_id = m.id
    
LEFT JOIN live_match_stats lms
ON lms.innings_id = i.id

JOIN teams t1
ON t1.team_id = m.team1_id

JOIN teams t2
ON t2.team_id = m.team2_id

WHERE m.id = $1
ORDER BY i.inning_number DESC
LIMIT 1
`

	var live models.LiveMatchResponse

	err := database.DB.QueryRow(
		ctx,
		query,
		matchID,
	).Scan(
		&live.ID,
		&live.Team1ID,
		&live.Team2ID,

		&live.Team1Name,
		&live.Team2Name,

		&live.CurrentInnings,

		&live.TotalRuns,
		&live.Wickets,

		&live.CompletedOvers,
		&live.BallsInCurrentOver,

		&live.StrikerID,
		&live.NonStrikerID,
		&live.CurrentBowlerID,
	)

	if err != nil {
		return nil, err
	}

	live.BattingTeamName = live.Team1Name
	live.BowlingTeamName = live.Team2Name

	return &live, nil
}
