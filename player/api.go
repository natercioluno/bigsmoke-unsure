package player

import "context"

type Client interface {
	CollectRank(ctx context.Context, roundID int64, player string) (*CollectRankRes, error)
}
