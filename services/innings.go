package services

import (
	"backend/database"
	"backend/database/dbHelper"
	"backend/models"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func StartInning(matchID string, req models.StartInningRequest) (string, error) {
	ctx := context.Background()

	if req.InningNumber == 2 {
		var completedOvers, wickets, totalOvers, playersPerTeam int

		validationQuery := `
			SELECT i.completed_overs, i.wickets, m.overs, m.players_per_team
			FROM innings i
			JOIN matches m ON i.match_id = m.id
			WHERE i.match_id = $1 AND i.inning_number = 1
		`
		err := database.DB.QueryRow(ctx, validationQuery, matchID).Scan(&completedOvers, &wickets, &totalOvers, &playersPerTeam)
		if err != nil {
			return "", errors.New("cannot start second inning: first inning records not found")
		}

		isOversQuotaMet := completedOvers >= totalOvers
		isAllOut := wickets >= (playersPerTeam - 1)

		if !isOversQuotaMet && !isAllOut {
			return "", fmt.Errorf("cannot start second inning: first inning is still active (%d/%d overs, %d wickets down)", completedOvers, totalOvers, wickets)
		}
	} else if req.InningNumber == 1 {
		return "", errors.New("inning 1 is automatically initialized when creating a match; cannot create manually")
	}

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)
	inningID := uuid.New().String()

	if err := dbHelper.UpdateMatchStatusLive(ctx, tx, matchID); err != nil {
		return "", err
	}
	if err := dbHelper.CreateInning(ctx, tx, inningID, matchID, req); err != nil {
		return "", err
	}
	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return inningID, nil
}
