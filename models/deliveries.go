package models

type RecordDeliveryRequest struct {
	InningID string `json:"inning_id"`

	StrikerID string `json:"striker_id"`

	NonStrikerID *string `json:"non_striker_id,omitempty"`

	BowlerID string `json:"bowler_id"`

	RunsBat int `json:"runs_bat"`

	Extras int `json:"extras"`

	ExtraType *string `json:"extra_type,omitempty"`

	Wicket bool `json:"wicket"`

	WicketType *string `json:"wicket_type,omitempty"`

	FielderID *string `json:"fielder_id,omitempty"`

	PlayerOutID *string `json:"player_out_id,omitempty"`

	IsFreeHit bool `json:"is_free_hit"`
}

type DeliveryRecord struct {
	InningID        string
	OverNumber      int
	BallNumber      int
	StrikerID       string
	NonStrikerID    *string
	BowlerID        string
	RunsBat         int
	Extras          int
	ExtraType       *string
	TotalRuns       int
	Wicket          bool
	WicketType      *string
	FielderID       *string
	PlayerOutID     *string
	IsFreeHit       bool
	IsLegalDelivery bool
}
