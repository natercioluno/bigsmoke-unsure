package ops

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"bigsmoke-unsure/player/config"
	"bigsmoke-unsure/player/db/rounds"

	"github.com/corverroos/unsure"
	"github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rpatterns"

	"bigsmoke-unsure/player/db/cursors"
	"bigsmoke-unsure/player/state"
)

func StartLoops(s *state.S) {
	go startMatchForever(s)
	go consumeEngineForever(s)
	go logHeadForever(s)
	go logPlayersForever(s)
}

func logPlayersForever(s *state.S) {

	ctx := unsure.FatedContext()

	for {
		for key, value := range s.PlayerClients() {
			if key == config.Player() {
				continue
			}

			log.Info(ctx, "player" + key)
			res, err := value.CollectRank(ctx, 0, config.Player())
			if err != nil {
				log.Error(ctx, err)
			} else {
				log.Info(ctx, fmt.Sprintf("rank %d part %d", res.Rank, res.Part))
			}
		}
		time.Sleep(10 * time.Second)
	}
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

		err := s.EngineClient().StartMatch(ctx, config.Team(), config.Players())

		if errors.Is(err, engine.ErrActiveMatch) {
			// Match active, just ignore
			return
		} else if err != nil {
			log.Error(ctx, errors.Wrap(err, "start match error"))
		} else {
			log.Info(ctx, "match started")
			return
		}

		time.Sleep(time.Second)
	}
}

func consumeEngineForever(s *state.S) {
	cs := cursors.ToStore(s.SmokeDB().DB)
	c := reflex.NewConsumable(s.EngineClient().Stream, cs)

	f := func(ctx context.Context, fate fate.Fate, event *reflex.Event) error {
		log.Info(nil, "==== EVENT ====")
		if reflex.IsAnyType(event.Type, engine.EventTypeRoundJoin) {
			//create round "join"
			e := joinRound(event, s, ctx)

			//included  OR excluded
			if e != nil {
				return e
			}
		}

		return fate.Tempt()
	}

	rpatterns.ConsumeForever(unsure.FatedContext, c.Consume, reflex.NewConsumer("first", f))
}

func joinRound(event *reflex.Event, s *state.S, ctx context.Context) error {
	round, err := strconv.ParseInt(event.ForeignID, 10, 64)
	if err != nil {
		return err
	}

	err = rounds.Create(ctx, s.SmokeDB().DB, round)
	if err != nil {
		return err
	}

	included, err := s.EngineClient().JoinRound(ctx, config.Team(), config.Player(), round)
	if errors.Is(err, engine.ErrAlreadyJoined) {
		included = true
	}
	if errors.Is(err, engine.ErrAlreadyExcluded) {
		included = false
	}
	if err != nil {
		return err
	}
	if included {
		//Update to included
		err = rounds.Included(ctx, s.SmokeDB().DB, round)
		if err != nil {
			return err
		}

		return nil
	}

	err = rounds.Excluded(ctx, s.SmokeDB().DB, round)
	if err != nil {
		return err
	}

	return nil
}
