package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gijsdb/super-tic-tac-toe/internal/api/helpers"
	h "github.com/gijsdb/super-tic-tac-toe/internal/api/helpers"
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"
	"github.com/inconshreveable/log15"
)

func genericResponse(w http.ResponseWriter, data []byte, err error) {
	if err != nil {
		errorResponse(err, "generic error", w)
	} else if data != nil {
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			// Should be API logger
			log15.Error("Error sending successful response: ", err.Error())
		}
	}
}

func errorResponse(err error, msg string, w http.ResponseWriter) {
	// Should be API logger
	fmt.Printf("%s: %s\n", msg, err.Error())
	res, herr := h.GenerateResponseError(msg, err, string(debug.Stack()))
	if herr != nil {
		fmt.Printf("Error creating response error: %s", herr)
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(res)
}

func JoinGame(w http.ResponseWriter, r *http.Request, m *game.Manager) {
	g := r.URL.Query().Get("id")
	p := r.URL.Query().Get("player")

	playerId := helpers.StringToInt(p)
	gameId := helpers.StringToInt(g)

	err := m.JoinGame(gameId, playerId)
	if err != nil {
		errorResponse(err, "game full", w)
	}
	bb, err := json.Marshal("true")
	genericResponse(w, bb, nil)
}

// Update the game board with the selected circle for the player
func UpdateGameBoard(w http.ResponseWriter, r *http.Request, m *game.Manager) {
	g := r.URL.Query().Get("gameid")
	p := r.URL.Query().Get("player")
	s := r.URL.Query().Get("square")
	c := r.URL.Query().Get("circle")

	if p == "" || s == "" || c == "" {
		err := errors.New("missing parameters")
		errorResponse(err, "Missing parameters in UpdateGameBoard", w)
	}

	// -1 because m.Games is an index
	gameIdx := helpers.StringToInt(g) - 1
	playerId := helpers.StringToInt(p)
	squareIdx := helpers.StringToInt(s)
	circleIdx := helpers.StringToInt(c)

	err := m.Games[gameIdx].GameBoard.Update(playerId, squareIdx, circleIdx)
	if err != nil {
		errorResponse(err, "Error updating game", w)
	}
	m.Games[gameIdx].Changeturn(p)

	jsonGame, err := json.Marshal(m.Games[gameIdx])
	if err != nil {
		errorResponse(err, "Error marshalling state", w)
	}

	genericResponse(w, jsonGame, nil)
}

func GetGameBoard(w http.ResponseWriter, r *http.Request, m *game.Manager) {
	id := r.URL.Query().Get("id")
	gameIdx := helpers.StringToInt(id) - 1
	game, err := m.GetGame(gameIdx)
	if err != nil {
		errorResponse(err, "Error getting game", w)
	}
	jsonGame, err := json.Marshal(game)
	if err != nil {
		errorResponse(err, "Error marshalling state", w)
	}

	genericResponse(w, jsonGame, nil)
}
