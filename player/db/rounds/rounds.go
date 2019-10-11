package rounds

import (
	"context"
	"database/sql"

	"bigsmoke-unsure/player/internal"
)

func Create(ctx context.Context, dbc *sql.DB, roundID int64) error {
	_, err := fsm.Insert(ctx, dbc, join{roundID, false})
	if err != nil {
		return err
	}

	return nil
}

func Included(ctx context.Context, dbc *sql.DB, roundID int64) error {
	err := fsm.Update(ctx, dbc, internal.RoundStatusJoin, internal.RoundStatusIncluded,
		collectReq{roundID})
	if err != nil {
		return err
	}

	return nil
}

func Excluded(ctx context.Context, dbc *sql.DB, roundID int64) error {
	err := fsm.Update(ctx, dbc, internal.RoundStatusJoin, internal.RoundStatusExcluded,
		collectReq{roundID})
	if err != nil {
		return err
	}

	return nil
}
