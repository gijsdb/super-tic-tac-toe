package main

import (
	"log"
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/router"
	"github.com/gijsdb/super-tic-tac-toe/state"

	"github.com/inconshreveable/log15"
)

func main() {

	state := state.InitGame()

	r := router.New(state)
	log15.Debug("Running server on localhost:2020")
	log.Fatal(http.ListenAndServe("localhost:2020", r))
}
