package models

type CreateMatchRequest struct {
	Team1Id          string `json:"team1_id" :"team1Id"`
	Team2Id          string `json:"team2_id" :"team2Id"`
	Venue            string `json:"venue" :"venue"`
	Overs            string `json:"overs" :"overs"`
	Players_per_team int    `json:"players_per_team" :"players_Per_Team"`
}
