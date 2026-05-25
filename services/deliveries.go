package services

import (
	"backend/database"
	"backend/database/dbHelper"
	"backend/models"
	"context"
)

func RecordDelivery(inningID string, req models.RecordDeliveryRequest) error {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	totalRuns := req.RunsBat + req.Extras
	isLegalDelivery := true
	if req.ExtraType != nil && (*req.ExtraType == "wide" || *req.ExtraType == "no_ball") {
		isLegalDelivery = false
	}

	if err := dbHelper.InsertDelivery(ctx, tx, inningID, req, totalRuns); err != nil {
		return err
	}

	if err := dbHelper.UpdateInningStats(ctx, tx, inningID, totalRuns, req.Extras, isLegalDelivery, req.Wicket); err != nil {
		return err
	}

	isLegalForBatter := isLegalDelivery
	if req.ExtraType != nil && *req.ExtraType == "no_ball" {
		isLegalForBatter = true
	}

	var dismissal *string
	if req.Wicket && req.PlayerOutID != nil && *req.PlayerOutID == req.StrikerID {
		dismissal = req.WicketType
	}

	if err := dbHelper.UpdateBattingScorecard(ctx, tx, inningID, req.StrikerID, req.RunsBat, isLegalForBatter, req.Wicket, dismissal); err != nil {
		return err
	}

	if err := dbHelper.UpdateBowlingScorecard(ctx, tx, inningID, req.BowlerID, totalRuns, isLegalDelivery, req.Wicket); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
