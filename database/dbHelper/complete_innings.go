package dbHelper

import (
	"backend/database"
	"context"
)

func CompleteInnings(
	ctx context.Context,
	inningID string,
) error {

	query := `
UPDATE innings
SET status = 'completed',
updated_at = NOW()
WHERE id = $1
`

	_, err := database.DB.Exec(
		ctx,
		query,
		inningID,
	)

	return err
}
