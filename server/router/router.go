package router

import (
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/controllers"
	"github.com/gijsdb/super-tic-tac-toe/middlewares"
	"github.com/gijsdb/super-tic-tac-toe/state"

	"github.com/gorilla/mux"
)

func New(state state.State) *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.CORS)

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetGameBoard(w, r, state)
	}).Methods("GET")

	api.HandleFunc("/updateboard", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateGameBoard(w, r, state)
	}).Methods("GET")

	return r
}
