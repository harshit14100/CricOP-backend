package dbHelper

import (
	"context"
	"fmt"

	"backend/database"
)

func GetNextBallNumber(
	ctx context.Context,
	inningID string,
) (int, int, error) {

	var legalBalls int

	query := `
SELECT legal_balls
FROM innings
WHERE id = $1
`

	err := database.DB.QueryRow(
		ctx,
		query,
		inningID,
	).Scan(&legalBalls)

	if err != nil {
		return 0, 0, err
	}

	overNumber := legalBalls / 6
	ballNumber := (legalBalls % 6) + 1

	fmt.Println(
		"OVER:",
		overNumber,
		"BALL:",
		ballNumber,
	)

	return overNumber, ballNumber, nil
}
