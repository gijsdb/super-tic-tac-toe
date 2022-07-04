package main

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/api"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/db"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"
)

func main() {

	manager := &game.Manager{
		DB:    db.Init(),
		Games: []game.Game{},
	}

	api.Run(manager)
}
