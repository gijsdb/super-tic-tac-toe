package router

import (
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/internal/api/controllers"
	"github.com/gijsdb/super-tic-tac-toe/internal/api/middlewares"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"

	"github.com/gorilla/mux"
)

func New(manager *game.Manager) *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.CORS)
	// r.Use(middlewares.PrepContext)
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		controllers.ListGames(w, r, manager)
	}).Methods("GET")

	api.HandleFunc("/creategame", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateGame(w, r, manager)
	}).Methods("GET")

	api.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetGame(w, r, manager)
	}).Methods("GET")

	api.HandleFunc("/joingame", func(w http.ResponseWriter, r *http.Request) {
		controllers.JoinGame(w, r, manager)
	}).Methods("GET")

	api.HandleFunc("/leavegame", func(w http.ResponseWriter, r *http.Request) {
		controllers.LeaveGame(w, r, manager)
	}).Methods("GET")

	api.HandleFunc("/updateboard", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateGameBoard(w, r, manager)
	}).Methods("GET")

	api.HandleFunc("/rolldice", func(w http.ResponseWriter, r *http.Request) {
		controllers.RollDice(w, r, manager)
	}).Methods("GET")

	api.HandleFunc("/createplayer", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreatePlayer(w, r, manager)
	}).Methods("GET")

	api.HandleFunc("/removeplayer", func(w http.ResponseWriter, r *http.Request) {
		controllers.RemovePlayer(w, r, manager)
	}).Methods("GET")

	return r
}
