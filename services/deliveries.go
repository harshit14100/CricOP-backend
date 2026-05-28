package services

import (
	"backend/database"
	"backend/database/dbHelper"
	"backend/models"
	"context"
)

func RecordDelivery(
	inningID string,
	req models.RecordDeliveryRequest,
) error {

	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	legalBalls, err := dbHelper.GetLegalBalls(
		ctx,
		tx,
		inningID,
	)
	if err != nil {
		return err
	}

	overNumber := legalBalls / 6
	ballNumber := (legalBalls % 6) + 1
	totalRuns := req.RunsBat + req.Extras
	isLegalDelivery := true

	if req.ExtraType != nil &&
		(*req.ExtraType == "wide" ||
			*req.ExtraType == "no_ball") {

		isLegalDelivery = false
	}

	delivery := models.DeliveryRecord{
		InningID:        inningID,
		OverNumber:      overNumber,
		BallNumber:      ballNumber,
		StrikerID:       req.StrikerID,
		NonStrikerID:    req.NonStrikerID,
		BowlerID:        req.BowlerID,
		RunsBat:         req.RunsBat,
		Extras:          req.Extras,
		ExtraType:       req.ExtraType,
		TotalRuns:       totalRuns,
		Wicket:          req.Wicket,
		WicketType:      req.WicketType,
		FielderID:       req.FielderID,
		PlayerOutID:     req.PlayerOutID,
		IsFreeHit:       req.IsFreeHit,
		IsLegalDelivery: isLegalDelivery,
	}

	err = dbHelper.InsertDelivery(
		ctx,
		tx,
		delivery,
	)
	if err != nil {
		return err
	}

	err = dbHelper.UpdateInningStats(
		ctx,
		tx,
		inningID,
		totalRuns,
		req.Extras,
		isLegalDelivery,
		req.Wicket,
	)
	if err != nil {
		return err
	}

	isLegalForBatter := isLegalDelivery
	if req.ExtraType != nil &&
		*req.ExtraType == "no_ball" {
		isLegalForBatter = true
	}

	var dismissal *string
	if req.Wicket &&
		req.PlayerOutID != nil &&
		*req.PlayerOutID == req.StrikerID {
		dismissal = req.WicketType
	}

	err = dbHelper.UpdateBattingScorecard(
		ctx,
		tx,
		inningID,
		req.StrikerID,
		req.RunsBat,
		isLegalForBatter,
		req.Wicket,
		dismissal,
	)
	if err != nil {
		return err
	}

	err = dbHelper.UpdateBowlingScorecard(
		ctx,
		tx,
		inningID,
		req.BowlerID,
		totalRuns,
		isLegalDelivery,
		req.Wicket,
		req.WicketType,
	)
	if err != nil {
		return err
	}

	rotateDueToRuns :=
		req.RunsBat == 1 ||
			req.RunsBat == 3

	nextLegalBalls := legalBalls

	if isLegalDelivery {
		nextLegalBalls++
	}

	rotateDueToOverEnd :=
		isLegalDelivery &&
			nextLegalBalls%6 == 0

	if rotateDueToRuns != rotateDueToOverEnd {

		err = dbHelper.RotateStrike(
			ctx,
			tx,
			inningID,
		)
		if err != nil {
			return err
		}
	}

	_, err = dbHelper.CheckAndConcludeMatch(
		ctx,
		tx,
		inningID,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
