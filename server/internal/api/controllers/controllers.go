package controllers

import (
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"
)

func CreateGame(w http.ResponseWriter, r *http.Request, m *game.Manager) {
	id, err := m.CreateNewGame()
	if err != nil {
		errorResponse(err, "Error creating game in controllers.go::CreateGame", w)
	}

	genericResponse(w, []byte{byte(id)}, nil)
}

func ListGames(w http.ResponseWriter, r *http.Request, m *game.Manager) {
	games, err := m.JSONGames()
	if err != nil {
		errorResponse(err, "Error listing games in controllers.go::ListGames", w)
	}

	genericResponse(w, games, nil)
}
