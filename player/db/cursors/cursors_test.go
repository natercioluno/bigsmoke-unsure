package cursors

import (
	"bigsmoke-unsure/player"
	"testing"

	"github.com/corverroos/unsure"
	"github.com/luno/reflex/rsql"

	"github.com/corverroos/play/db"
)

func TestCursorsTable(t *testing.T) {
	defer unsure.CheatFateForTesting(t)()
	dbc := db.ConnectForTesting(t)
	defer dbc.Close()

	rsql.TestCursorsTable(t, dbc, player.cursors)
}
