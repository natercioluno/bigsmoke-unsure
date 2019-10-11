package state

import (
	"bigsmoke-unsure/player/client"
	"bigsmoke-unsure/player/config"
	"fmt"
	"github.com/corverroos/unsure/engine"
	ec "github.com/corverroos/unsure/engine/client"
	"strconv"

	"bigsmoke-unsure/player/db"
)

type S struct {
	engineClient engine.Client
	db *db.SmokeDB
	playerClients map[string]*client.Client
}

func (s *S) EngineClient() engine.Client {
	return s.engineClient
}

func (s *S) SmokeDB() *db.SmokeDB {
	return s.db
}

func (s *S) PlayerClients() map[string]*client.Client {
	return s.playerClients
}

// New returns a new engine state.
func New() (*S, error) {
	var (
		s   S
		err error
	)

	s.engineClient, err = ec.New()
	if err != nil {
		return nil, err
	}

	s.db, err = db.Connect()
	if err != nil {
		return nil, err
	}

	s.playerClients = make(map[string]*client.Client)

	for i := 1; i <= config.Players(); i++ {
		port := 23048 + i - 1
		player:= strconv.Itoa(i)
		client, err:=client.New(client.WithAddress(fmt.Sprintf("127.0.0.1:%v", port)))
		if err != nil {
			return nil, err
		}
		s.playerClients[player] = client
	}

	return &s, nil
}
