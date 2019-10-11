package client

import (
	"context"
	"flag"
	"github.com/corverroos/unsure"
	p "bigsmoke-unsure/player"
	pb "bigsmoke-unsure/player/playerpb"
	"google.golang.org/grpc"
)

var addr = flag.String("player_address", "", "host:port of player gRPC service")

var _ p.Client = (*client)(nil)

type client struct {
	address   string
	rpcConn   *grpc.ClientConn
	rpcClient pb.PlayerClient
}


type option func(*client)

func WithAddress(address string) option {
	return func(c *client) {
		c.address = address
	}
}

func New(opts ...option) (*client, error) {
	c := client{
		address: *addr,
	}
	for _, o := range opts {
		o(&c)
	}

	var err error
	c.rpcConn, err = unsure.NewClient(c.address)
	if err != nil {
		return nil, err
	}

	c.rpcClient = pb.NewPlayerClient(c.rpcConn)

	return &c, nil
}

func (c *client) CollectRank(ctx context.Context, roundID int64, player string) (*p.CollectRankRes, error) {
	res, err := c.rpcClient.CollectRank(ctx, &pb.CollectRankReq {
		RoundId: roundID,
		Player:  player,
	})
	if err != nil {
		return nil, err
	}


	return &p.CollectRankRes{
		Rank: res.Rank,
		Part: res.Part,
	}, nil
}

