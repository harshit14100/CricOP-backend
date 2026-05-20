package models

type CreateTeamRequest struct {
	Name string `json:"name" binding:"required"`
}

type AddPlayersToTeamRequest struct {
	PlayerIDs []string `json:"player_ids"`
}
