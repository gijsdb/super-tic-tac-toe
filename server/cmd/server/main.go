package main

import (
	"flag"

	"github.com/gijsdb/super-tic-tac-toe/internal/api"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"
)

func main() {
	clearDB := flag.Bool("cleardb", true, "Clear the database")

	flag.Parse()
	var m *game.Manager

	if m == nil {
		m = game.NewManager()
	}

	if *clearDB {
		m.ClearDB()
	}

	api.Run(m)
}
