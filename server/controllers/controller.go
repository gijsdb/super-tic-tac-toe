package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"

	h "github.com/gijsdb/super-tic-tac-toe/helpers"
	"github.com/gijsdb/super-tic-tac-toe/state"
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

// Update the game board with the selected circle for the player
func UpdateGameBoard(w http.ResponseWriter, r *http.Request, state state.State) {
	player := r.URL.Query().Get("player")
	square := r.URL.Query().Get("square")
	circle := r.URL.Query().Get("circle")

	if player == "" || square == "" || circle == "" {
		err := errors.New("missing parameters")
		errorResponse(err, "Missing parameters in UpdateGameBoard", w)
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

	state.GameBoard.Update(playerInt, squareInt, circleInt)

	GetGameBoard(w, r, state)
}

// Return the state of the game board at the end of each turn
func GetGameBoard(w http.ResponseWriter, r *http.Request, state state.State) {

	jsonState, err := json.Marshal(state)

	if err != nil {
		errorResponse(err, "Error marshalling state", w)
	}

	genericResponse(w, jsonState, nil)
}
