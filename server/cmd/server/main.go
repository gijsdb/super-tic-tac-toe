package main

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/api"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"
)

func main() {
	var m *game.Manager
	if m == nil {
		m = game.NewManager()
	}

	api.Run(m)
}
