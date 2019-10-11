package config

import (
	"flag"
)

var (
	team    = flag.String("team", "bigsmoke", "team name")
	player  = flag.String("player", "1", "player name")
	players = flag.Int("players", 4, "number of players in the team")
    addr = flag.String("player_address", "localhost:23048", "host:port of player gRPC service")
)

func Team() string {
	return *team
}

func Player() string {
	return *player
}

func Players() int {
	return *players
}

func Address() string {
	return *addr
}