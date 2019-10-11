package cursors

import (
	"database/sql"

	"github.com/luno/reflex"
	"github.com/luno/reflex/rsql"
)

var cursors = rsql.NewCursorsTable("smoke_cursors",
	rsql.WithCursorCursorField("`cursor`"))

func ToStore(dbc *sql.DB) reflex.CursorStore {
	return cursors.ToStore(dbc, rsql.WithCursorAsyncDisabled()) // Have to disable async since it doesn't use fated context.
}
