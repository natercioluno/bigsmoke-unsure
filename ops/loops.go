package ops

import (
	"context"
	"flag"
	"time"

	"github.com/corverroos/unsure"
	"github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rpatterns"

	"github.com/corverroos/play/db/cursors"
	"github.com/corverroos/play/state"
)

var (
	team    = flag.String("team", "bigsmoke", "team name")
	player  = flag.String("player", "smoker", "player name")
	players = flag.Int("players", 4, "number of players in the team")
)

func StartLoops(s *state.S) {
	go startMatchForever(s)
	go consumeEngineForever(s)
	go logHeadForever(s)
}

func logHeadForever(s *state.S) {
	var delay time.Duration
	for {
		time.Sleep(delay)
		delay = time.Second

		ctx := unsure.FatedContext()
		cl, err := s.EngineClient().Stream(ctx, "", reflex.WithStreamFromHead())
		if err != nil {
			log.Error(ctx, errors.Wrap(err, "log head stream error"))
			continue
		}
		for {
			e, err := cl.Recv()
			if err != nil {
				log.Error(ctx, errors.Wrap(err, "log head recv error"))
				break
			}
			typ := engine.EventType(e.Type.ReflexType())
			log.Info(ctx, "log head consumed event",
				j.MKV{"id": e.ForeignIDInt(), "type": typ, "latency": time.Since(e.Timestamp)})
		}
	}
}

func startMatchForever(s *state.S) {
	for {
		ctx := unsure.ContextWithFate(context.Background(), unsure.DefaultFateP())

		err := s.EngineClient().StartMatch(ctx, *team, *players)

		if errors.Is(err, engine.ErrActiveMatch) {
			// Match active, just ignore
			return
		} else if err != nil {
			log.Error(ctx, errors.Wrap(err, "start match error"))
		} else {
			log.Info(ctx, "match started")
			return
		}

		time.Sleep(time.Minute)
	}
}

func consumeEngineForever(s *state.S)  {
	cs := cursors.ToStore(s.SmokeDB().DB)
	c := reflex.NewConsumable(s.EngineClient().Stream, cs)
	rpatterns.ConsumeForever(unsure.FatedContext, c.Consume, reflex.NewConsumer("first", EngineConsume))
}

func EngineConsume(ctx context.Context, fate fate.Fate, event *reflex.Event) error {
	return fate.Tempt()
}
