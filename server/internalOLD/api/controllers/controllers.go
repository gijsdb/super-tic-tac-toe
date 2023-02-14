package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/internal/api/helpers"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/manager"
)

func CreateGame(w http.ResponseWriter, r *http.Request, m *manager.Manager) {
	p := r.URL.Query().Get("player")
	playerId := helpers.StringToInt(p)

	game, err := m.CreateGame(playerId)
	if err != nil {
		errorResponse(err, "Error creating game in controllers.go::CreateGame", w)
	}
	bb, err := json.Marshal(game)
	if err != nil {
		errorResponse(err, "Error marshalling to JSON game in controllers.go::CreateGame", w)
	}
	genericResponse(w, bb, nil)
}

func ListGames(w http.ResponseWriter, r *http.Request, m *manager.Manager) {
	games, err := m.ListGames()
	if err != nil {
		errorResponse(err, "Error listing games in controllers.go::ListGames", w)
	}

	genericResponse(w, games, nil)
}

func CreatePlayer(w http.ResponseWriter, r *http.Request, m *manager.Manager) {
	player := r.URL.Query().Get("id")
	playerId := helpers.StringToInt(player)

	playerId = m.CreatePlayer(playerId)
	bb, err := json.Marshal(playerId)
	if err != nil {
		errorResponse(err, "Error marshalling to JSON game in controllers.go::CreateClient", w)
	}
	genericResponse(w, bb, nil)
}

func RemovePlayer(w http.ResponseWriter, r *http.Request, m *manager.Manager) {
	player := r.URL.Query().Get("id")
	playerId := helpers.StringToInt(player)

	m.SetPlayerInactive(playerId)

	genericResponse(w, nil, nil)
}
