package models

import "github.com/google/uuid"

type StartInningRequest struct {
	InningNumber  int       `json:"inning_number" binding:"required"`
	BattingTeamID uuid.UUID `json:"batting_team_id" binding:"required"`
	BowlingTeamID uuid.UUID `json:"bowling_team_id" binding:"required"`
}

type InningsState struct {
	Status           string `json:"status"`
	MatchID          string `json:"match_id" binding:"required"`
	BattingTeamID    string `json:"batting_team_id" binding:"required"`
	BowlingTeamID    string `json:"bowling_team_id" binding:"required"`
	TotalWickets     int    `json:"total_wickets"`
	AllowSoloBatting bool   `json:"allow_solo_batting"`
}
