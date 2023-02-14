package api

import (
	"log"
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/internal/api/router"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/manager"
	"github.com/inconshreveable/log15"
)

func Run(manager *manager.Manager) {
	r := router.New(manager)
	log15.Info("Running server on localhost:2020")
	log.Fatal(http.ListenAndServe("localhost:2020", r))
}
