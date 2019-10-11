package server

import "time"

type CollectRank struct {
	RoundID        int64
	Player      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
