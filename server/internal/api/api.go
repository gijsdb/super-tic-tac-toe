package api

import (
	"log"
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/internal/api/router"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"
	"github.com/inconshreveable/log15"
)

func Run(state game.State) {

	r := router.New(state)
	log15.Debug("Running server on localhost:2020")
	log.Fatal(http.ListenAndServe("localhost:2020", r))
}
