package main

import (
	"bigsmoke-unsure/player/config"
	"bigsmoke-unsure/player/playerpb"
	"bigsmoke-unsure/player/server"
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

	go serveGRPCForever(s)

	ops.StartLoops(s)

	unsure.WaitForShutdown()
}

func serveGRPCForever(s *state.S) {
	grpcServer, err := unsure.NewServer(config.Address())
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new grpctls server"))
	}

	server := server.New()
	player.RegisterPlayerServer(grpcServer.GRPCServer(), server)

	unsure.RegisterNoErr(func() {
		grpcServer.Stop()
	})

	unsure.Fatal(grpcServer.ServeForever())
}
