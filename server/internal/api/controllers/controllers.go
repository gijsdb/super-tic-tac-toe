package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"
)

func CreateGame(w http.ResponseWriter, r *http.Request, m *game.Manager) {
	player := r.URL.Query().Get("player")

	id, err := m.CreateGame(player)
	if err != nil {
		errorResponse(err, "Error creating game in controllers.go::CreateGame", w)
	}
	bb, err := json.Marshal(id)
	if err != nil {
		errorResponse(err, "Error marshalling to JSON game in controllers.go::CreateGame", w)
	}
	genericResponse(w, bb, nil)
}

func ListGames(w http.ResponseWriter, r *http.Request, m *game.Manager) {
	games, err := m.ListGames()
	if err != nil {
		errorResponse(err, "Error listing games in controllers.go::ListGames", w)
	}

	genericResponse(w, games, nil)
}

func CreatePlayer(w http.ResponseWriter, r *http.Request, m *game.Manager) {
	id := m.CreatePlayer()
	bb, err := json.Marshal(id)
	if err != nil {
		errorResponse(err, "Error marshalling to JSON game in controllers.go::CreateClient", w)
	}
	genericResponse(w, bb, nil)
}
