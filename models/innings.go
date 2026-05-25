package models

import "github.com/google/uuid"

type StartInningRequest struct {
	InningNumber  int       `json:"inning_number" binding:"required"`
	BattingTeamID uuid.UUID `json:"batting_team_id" binding:"required"`
	BowlingTeamID uuid.UUID `json:"bowling_team_id" binding:"required"`
}
