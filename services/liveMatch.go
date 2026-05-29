package services

import (
	"backend/database/dbHelper"
	"context"
)

func GetLiveMatch(matchID string) (interface{}, error) {

	ctx := context.Background()

	return dbHelper.GetLiveMatch(
		ctx,
		matchID,
	)
}
