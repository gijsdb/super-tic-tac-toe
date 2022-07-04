package main

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/api"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"
)

func main() {

	state := game.InitGame()

	api.Run(state)
}
