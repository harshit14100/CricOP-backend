package services

import (
	"backend/database"
	"backend/database/dbHelper"
	"backend/models"
	"context"

	"github.com/google/uuid"
)

func StartInning(matchID string, req models.StartInningRequest) (string, error) {
	ctx := context.Background()

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
