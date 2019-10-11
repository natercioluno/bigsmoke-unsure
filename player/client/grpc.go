package client

import (
	"bigsmoke-unsure/player/config"
	"context"
	"github.com/corverroos/unsure"
	p "bigsmoke-unsure/player"
	pb "bigsmoke-unsure/player/playerpb"
	"google.golang.org/grpc"
)

var _ p.Client = (*Client)(nil)

type Client struct {
	address   string
	rpcConn   *grpc.ClientConn
	rpcClient pb.PlayerClient
}


type option func(*Client)

func WithAddress(address string) option {
	return func(c *Client) {
		c.address = address
	}
}

func New(opts ...option) (*Client, error) {
	c := Client{
		address: config.Address(),
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

func (c *Client) CollectRank(ctx context.Context, roundID int64, player string) (*p.CollectRankRes, error) {
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
