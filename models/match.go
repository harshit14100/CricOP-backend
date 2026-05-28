package models

import "time"

type CreateMatchRequest struct {
	Team1ID          string `json:"team1_id" binding:"required"`
	Team2ID          string `json:"team2_id" binding:"required"`
	Venue            string `json:"venue"`
	Overs            int    `json:"overs" binding:"required"`
	Players_per_team int    `json:"players_per_team" binding:"required"`
}

type Match struct {
	ID             string     `json:"id" db:"id"`
	HostUserID     *string    `json:"host_user_id" db:"host_user_id"`
	Team1ID        string     `json:"team1_id" db:"team1_id"`
	Team1Name      string     `db:"team_1_name" json:"team_1_name"`
	Team2ID        string     `json:"team2_id" db:"team2_id"`
	Team2Name      string     `db:"team_2_name" json:"team_2_name"`
	Venue          *string    `json:"venue" db:"venue"`
	Overs          int        `json:"overs" db:"overs"`
	PlayersPerTeam int        `json:"players_per_team" db:"players_per_team"`
	Status         string     `json:"status" db:"status"` // scheduled, live, completed, cancelled
	TossWinnerID   *string    `json:"toss_winner_id" db:"toss_winner_id"`
	TossDecision   *string    `json:"toss_decision" db:"toss_decision"` // bat, bowl
	WinnerTeamID   *string    `json:"winner_team_id" db:"winner_team_id"`
	ManOfMatchID   *string    `json:"man_of_match_id" db:"man_of_match_id"`
	WorstPlayerID  *string    `json:"worst_player_id" db:"worst_player_id"`
	StartedAt      *time.Time `json:"started_at" db:"started_at"`
	EndedAt        *time.Time `json:"ended_at" db:"ended_at"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

type StartMatchRequest struct {
	Team1ID         string `json:"team1_id"`
	Team2ID         string `json:"team2_id"`
	Venue           string `json:"venue"`
	Overs           int    `json:"overs"`
	TossWinnerID    string `json:"toss_winner_id"`
	TossDecision    string `json:"toss_decision"`
	BattingTeamID   string `json:"batting_team_id"`
	BowlingTeamID   string `json:"bowling_team_id"`
	StrikerID       string `json:"striker_id"`
	NonStrikerID    string `json:"non_striker_id"`
	CurrentBowlerID string `json:"current_bowler_id"`
}
type SuperStartMatchRequest struct {
	Team1ID         string   `json:"team1_id"`
	Team2ID         string   `json:"team2_id"`
	Venue           string   `json:"venue"`
	Overs           int      `json:"overs"`
	Team1Players    []string `json:"team1_players"`
	Team2Players    []string `json:"team2_players"`
	TossWinnerID    string   `json:"toss_winner_id"`
	TossDecision    string   `json:"toss_decision"`
	BattingTeamID   string   `json:"batting_team_id"`
	BowlingTeamID   string   `json:"bowling_team_id"`
	StrikerID       string   `json:"striker_id"`
	NonStrikerID    string   `json:"non_striker_id"`
	CurrentBowlerID string   `json:"current_bowler_id"`
}

type LiveMatchStateResponse struct {
	MatchID            string  `json:"match_id"`
	Venue              string  `json:"venue"`
	Overs              int     `json:"overs"`
	Status             string  `json:"status"`
	BattingTeamName    string  `json:"batting_team_name"`
	BowlingTeamName    string  `json:"bowling_team_name"`
	StrikerName        string  `json:"striker_name"`
	NonStrikerName     string  `json:"non_striker_name"`
	CurrentBowlerName  string  `json:"current_bowler_name"`
	InningNumber       int     `json:"inning_number"`
	TotalRuns          int     `json:"total_runs"`
	Wickets            int     `json:"wickets"`
	CompletedOvers     int     `json:"completed_overs"`
	BallsInCurrentOver int     `json:"balls_in_current_over"`
	DisplayOvers       float64 `json:"display_overs"`
}
