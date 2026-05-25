package models

import "github.com/google/uuid"

type Toss struct {
	TossWinnerID uuid.UUID `json:"toss_winner_id" binding:"required"`
	TossDecision string    `json:"toss_decision" binding:"required"`
}
