package models

type ScoreBallRequest struct {
	MatchID       string `json:"matchId" binding:"required"`
	Runs          int    `json:"runs"`
	IsWicket      bool   `json:"isWicket"`
	IsWide        bool   `json:"isWide"`
	IsNoBall      bool   `json:"isNoBall"`
	IsBye         bool   `json:"isBye"`
	IsLegBye      bool   `json:"isLegBye"`
	DismissalType string `json:"dismissalType"`
	FielderID     string `json:"fielderId"`
	StrikerID     string `json:"strikerId" binding:"required"`
	NonStrikerID  string `json:"nonStrikerId"`
	BowlerID      string `json:"bowlerId" binding:"required"`
	NewBatsmanID  string `json:"newBatsmanId"`
	NewBowlerID   string `json:"newBowlerId"`
}
