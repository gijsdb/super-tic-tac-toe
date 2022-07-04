package router

import (
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/internal/api/controllers"
	"github.com/gijsdb/super-tic-tac-toe/internal/api/middlewares"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"

	"github.com/gorilla/mux"
)

func New(state game.State) *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.CORS)
	// r.Use(middlewares.PrepContext)
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetGameBoard(w, r, state)
	}).Methods("GET")

	api.HandleFunc("/updateboard", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateGameBoard(w, r, state)
	}).Methods("GET")

	return r
}
