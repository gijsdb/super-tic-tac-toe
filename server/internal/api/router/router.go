package router

import (
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/internal/api/controllers"
	"github.com/gijsdb/super-tic-tac-toe/internal/api/middlewares"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/manager"

	"github.com/gorilla/mux"
)

func New(m *manager.Manager) *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.CORS)
	// r.Use(middlewares.PrepContext)
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		controllers.ListGames(w, r, m)
	}).Methods("GET")

	api.HandleFunc("/creategame", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateGame(w, r, m)
	}).Methods("GET")

	api.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetGame(w, r, m)
	}).Methods("GET")

	api.HandleFunc("/joingame", func(w http.ResponseWriter, r *http.Request) {
		controllers.JoinGame(w, r, m)
	}).Methods("GET")

	api.HandleFunc("/leavegame", func(w http.ResponseWriter, r *http.Request) {
		controllers.LeaveGame(w, r, m)
	}).Methods("GET")

	api.HandleFunc("/updateboard", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateGameBoard(w, r, m)
	}).Methods("GET")

	api.HandleFunc("/rolldice", func(w http.ResponseWriter, r *http.Request) {
		controllers.RollDice(w, r, m)
	}).Methods("GET")

	api.HandleFunc("/createplayer", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreatePlayer(w, r, m)
	}).Methods("GET")

	api.HandleFunc("/removeplayer", func(w http.ResponseWriter, r *http.Request) {
		controllers.RemovePlayer(w, r, m)
	}).Methods("GET")

	return r
}
