package server

import (
	"bigsmoke-unsure/player/playerpb"
	"context"
)

var _ player.PlayerServer = (*Server)(nil)

type Server struct {
}

func New() *Server {
	return &Server{
	}
}

func (Server) CollectRank(context.Context, *player.CollectRankReq) (*player.CollectRankRes, error) {
	return &player.CollectRankRes{
		Rank:1,
		Part:2,
	}, nil
}
