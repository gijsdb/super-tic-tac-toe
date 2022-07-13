package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"

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
	gameId := r.URL.Query().Get("id")
	player := r.URL.Query().Get("player")
	idInt, err := strconv.Atoi(gameId)
	if err != nil {
		errorResponse(err, "Error converting gameid to int", w)
	}

	m.JoinGame(idInt, player)
}

// Update the game board with the selected circle for the player
func UpdateGameBoard(w http.ResponseWriter, r *http.Request, m *game.Manager) {
	gameId := r.URL.Query().Get("gameid")
	player := r.URL.Query().Get("player")
	square := r.URL.Query().Get("square")
	circle := r.URL.Query().Get("circle")

	if player == "" || square == "" || circle == "" {
		err := errors.New("missing parameters")
		errorResponse(err, "Missing parameters in UpdateGameBoard", w)
	}

	idInt, err := strconv.Atoi(gameId)
	if err != nil {
		errorResponse(err, "Error converting square to int", w)
	}

	squareInt, err := strconv.Atoi(square)
	if err != nil {
		errorResponse(err, "Error converting square to int", w)
	}
	playerInt, err := strconv.Atoi(player)
	if err != nil {
		errorResponse(err, "Error converting square to int", w)
	}
	circleInt, err := strconv.Atoi(circle)
	if err != nil {
		errorResponse(err, "Error converting square to int", w)
	}
	if err != nil {
		errorResponse(err, "Error getting gameboard from JSON", w)
	}

	m.Games[idInt].State.GameBoard.Update(playerInt, squareInt, circleInt)

	if err != nil {
		errorResponse(err, "Error getting game", w)
	}
	jsonGame, err := json.Marshal(m.Games[idInt])
	if err != nil {
		errorResponse(err, "Error marshalling state", w)
	}

	genericResponse(w, jsonGame, nil)
}

func GetGameBoard(w http.ResponseWriter, r *http.Request, m *game.Manager) {
	id := r.URL.Query().Get("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		errorResponse(err, "Error converting string to int", w)
	}
	game, err := m.GetGame(gameId)
	if err != nil {
		errorResponse(err, "Error getting game", w)
	}
	jsonGame, err := json.Marshal(game)
	if err != nil {
		errorResponse(err, "Error marshalling state", w)
	}

	genericResponse(w, jsonGame, nil)
}
