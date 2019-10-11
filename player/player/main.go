package main

import (
	"github.com/corverroos/unsure"
	"github.com/luno/jettison/errors"

	"github.com/corverroos/play/ops"
	"github.com/corverroos/play/state"
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
