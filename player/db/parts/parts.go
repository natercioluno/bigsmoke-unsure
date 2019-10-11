package parts

import (
	"context"
	"database/sql"
)

func Part(ctx context.Context, tx *sql.DB, roundID int64, player string) (int64, error) {
	row := tx.QueryRowContext(ctx, "select `part` from player_parts where round_id=? and player_id=?", roundID, player)

	var part int64
	err := row.Scan(part)

	if err !=  nil {
		return 0, err
	}


	return part, nil
}
