package models

type CreateMatchRequest struct {
	Team1ID          string `json:"team1_id" binding:"required"`
	Team2ID          string `json:"team2_id" binding:"required"`
	Venue            string `json:"venue"`
	Overs            int    `json:"overs" binding:"required"`
	Players_per_team int    `json:"players_per_team" binding:"required"`
}
