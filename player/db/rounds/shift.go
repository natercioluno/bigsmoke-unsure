package rounds

import (
	"github.com/luno/reflex/rsql"
	"github.com/luno/shift"

	"bigsmoke-unsure/player/internal"
)

//go:generate shiftgen -inserter=join -updaters=playersReady -table=rounds

var fsm = shift.NewFSM(rsql.NewEventsTableInt("smoke_events")).
	Insert(internal.RoundStatusJoin, join{}, internal.RoundStatusIncluded, internal.RoundStatusExcluded).
	Build()

type join struct {
	roundID  int64
	included bool
}

type playersReady struct {
	ID      int64
	MatchID int64
}
