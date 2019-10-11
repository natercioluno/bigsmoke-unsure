package main

import (
	"github.com/corverroos/unsure"
	"github.com/luno/jettison/errors"

	"bigsmoke-unsure/player/ops"
	"bigsmoke-unsure/player/state"
)

func main() {
	unsure.Bootstrap()

	s, err := state.New()
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new state error"))
	}

	ops.StartLoops(s)

	unsure.WaitForShutdown()
}
