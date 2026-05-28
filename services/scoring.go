package services

import (
	"backend/database/dbHelper"
	"backend/models"
	"context"
)

func RecordBall(c context.Context, input models.RecordDeliveryRequest) error {

	inningsData, err := dbHelper.GetInningsState(c, input.InningID)
	if err != nil {
		return err
	}

	if inningsData.Status != "ongoing" {
		return err
	}

	activePlayers, err := dbHelper.GetActivePlayersCount(
		c,
		inningsData.MatchID,
		inningsData.BattingTeamID,
	)

	if err != nil {
		return err
	}

	lastPlayerRemaining :=
		inningsData.TotalWickets >= activePlayers-1

	if lastPlayerRemaining {

		if !inningsData.AllowSoloBatting {
			return dbHelper.CompleteInnings(
				c,
				input.InningID,
			)
		}

		if input.NonStrikerID != nil {
			return err
		}

	} else {

		if input.NonStrikerID == nil {
			return err
		}
	}

	return nil
}
