package main

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/api"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/manager"
)

func main() {
	var m *manager.Manager
	if m == nil {
		m = manager.NewManager()
	}

	api.Run(m)
}
