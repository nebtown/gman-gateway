package gameserver

import (
	"fmt"
)

type GameServer struct {
	Game     string
	Instance string
}

func NewGameServer(game string, instance string) *GameServer {
	return &GameServer{
		Game:     game,
		Instance: instance,
	}
}

func (s *GameServer) Start() {
	fmt.Printf("Starting server %s, instance %s\n", s.Game, s.Instance)
}
