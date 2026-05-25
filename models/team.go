package models

type CreateTeamRequest struct {
	Name string `json:"name" binding:"required"`
}

type AddPlayersToTeamRequest struct {
	PlayerIDs []string `json:"player_ids"`
}

type Team struct {
	ID        string `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	CreatedBy string `db:"created_by" json:"created_by"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}
