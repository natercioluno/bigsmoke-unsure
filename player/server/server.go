package server

import (
	"bigsmoke-unsure/player/db/parts"
	"bigsmoke-unsure/player/playerpb"
	"bigsmoke-unsure/player/state"
	"context"
)

var _ player.PlayerServer = (*Server)(nil)

type Server struct {
	state *state.S
}

func New(s *state.S) *Server {
	return &Server{
		state: s,
	}
}

func (ser Server) CollectRank(ctx context.Context, req *player.CollectRankReq) (*player.CollectRankRes, error) {
	db := ser.state.SmokeDB().DB

	part, err := parts.Part(ctx, db, req.RoundId, req.Player)
	if err != nil {
		return nil, err
	}

	return &player.CollectRankRes{
		Rank:1,
		Part:part,
	}, nil
}
