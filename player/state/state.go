package state

import (
	"github.com/corverroos/unsure/engine"
	ec "github.com/corverroos/unsure/engine/client"

	"bigsmoke-unsure/player/db"
)

type S struct {
	engineClient engine.Client
	db *db.SmokeDB
}

func (s *S) EngineClient() engine.Client {
	return s.engineClient
}

func (s *S) SmokeDB() *db.SmokeDB {
	return s.db
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

	return &s, nil
}
