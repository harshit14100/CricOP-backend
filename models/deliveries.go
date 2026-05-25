package models

type RecordDeliveryRequest struct {
	OverNumber   int     `json:"over_number" binding:"required"`
	BallNumber   int     `json:"ball_number" binding:"required"`
	StrikerID    string  `json:"striker_id" binding:"required"`
	NonStrikerID string  `json:"non_striker_id" binding:"required"`
	BowlerID     string  `json:"bowler_id" binding:"required"`
	RunsBat      int     `json:"runs_bat"`
	Extras       int     `json:"extras"`
	ExtraType    *string `json:"extra_type"` // wide, no_ball etc
	Wicket       bool    `json:"wicket"`
	WicketType   *string `json:"wicket_type"`
	FielderID    *string `json:"fielder_id"`
	PlayerOutID  *string `json:"player_out_id"`
	IsFreeHit    bool    `json:"is_free_hit"`
}
