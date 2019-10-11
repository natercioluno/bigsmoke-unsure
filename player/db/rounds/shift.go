package rounds

import (
	"github.com/luno/reflex/rsql"
	"github.com/luno/shift"

	"bigsmoke-unsure/player/internal"
)

//go:generate shiftgen -inserter=join -updaters=collectReq -table=rounds

var fsm = shift.NewFSM(rsql.NewEventsTableInt("smoke_events")).
	Insert(internal.RoundStatusJoin, join{}, internal.RoundStatusIncluded, internal.RoundStatusExcluded).
	Update(internal.RoundStatusIncluded, collectReq{}, internal.RoundStatusCollect).
	Build()

type join struct {
	Round    int64
	Included bool
}

type collectReq struct {
	ID int64
}
