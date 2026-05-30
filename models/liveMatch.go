package models

import "github.com/google/uuid"

type LiveMatchResponse struct {
	ID        string `json:"id"`
	Team1ID   string `json:"team1_id"`
	Team2ID   string `json:"team2_id"`
	Team1Name string `json:"team_1_name"`
	Team2Name string `json:"team_2_name"`

	BattingTeamName string `json:"batting_team_name"`
	BowlingTeamName string `json:"bowling_team_name"`

	CurrentInnings int `json:"currentInnings"`

	TotalRuns int `json:"total_runs"`
	Wickets   int `json:"wickets"`

	CompletedOvers     int `json:"completed_overs"`
	BallsInCurrentOver int `json:"balls_in_current_over"`

	StrikerID       *string `json:"striker_id"`
	NonStrikerID    *string `json:"non_striker_id"`
	CurrentBowlerID *string `json:"current_bowler_id"`

	StrikerName    *string `json:"striker_name"`
	NonStrikerName *string `json:"non_striker_name"`
	BowlerName     *string `json:"bowler_name"`

	Innings []InningsResponse `json:"innings"`
}

type InningsResponse struct {
	ID           string `json:"id"`
	InningNumber int    `json:"inning_number"`
	BattingTeam  string `json:"battingTeam"`

	Runs    int `json:"runs"`
	Wickets int `json:"wickets"`
	Balls   int `json:"balls"`

	Batsmen []BatsmanScore `json:"batsmen"`
	Bowlers []BowlerScore  `json:"bowlers"`
}

type BatsmanScore struct {
	PlayerID string `json:"playerId"`
	Runs     int    `json:"runs"`
	Balls    int    `json:"balls"`
	IsOut    bool   `json:"isOut"`
}

type BowlerScore struct {
	PlayerID string `json:"playerId"`
	Overs    int    `json:"overs"`
	Balls    int    `json:"balls"`
	Maidens  int    `json:"maidens"`
	Runs     int    `json:"runs"`
	Wickets  int    `json:"wickets"`
}

type MatchListResponse struct {
	ID              uuid.UUID `json:"id"`
	Status          string    `json:"status"`
	Team1ID         uuid.UUID `json:"team1_id"`
	Team2ID         uuid.UUID `json:"team2_id"`
	Team1Name       string    `json:"team_1_name"`
	Team2Name       string    `json:"team_2_name"`
	Venue           string    `json:"venue"`
	Overs           int       `json:"overs"`
	CurrentInnings  int       `json:"currentInnings"`
	TotalRuns       int       `json:"total_runs"`
	Wickets         int       `json:"wickets"`
	BattingTeamID   uuid.UUID `json:"batting_team_id"`
	BattingTeamName string    `json:"batting_team_name"`
}
