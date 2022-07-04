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

// Update the game board with the selected circle for the player
func UpdateGameBoard(w http.ResponseWriter, r *http.Request, manager *game.Manager) {
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

	gb, err := manager.Games[0].State.JsonToGameboard(manager.Games[0].State.GameBoard)
	if err != nil {
		errorResponse(err, "Error getting gameboard from JSON", w)
	}
	gb.Update(playerInt, squareInt, circleInt)

	GetGameBoard(w, r, manager)
}

// Return the state of the game board at the end of each turn

// type State struct {
// 	GameBoard  game.GameBoard `json:"game_board"`
// 	GameID     int            `json:"game_id"`
// 	PlayerTurn int            `json:"player_turn"`
// 	GameOver   bool           `json:"game_over"`
// 	Winner     int            `json:"winner"`
// }

func GetGameBoard(w http.ResponseWriter, r *http.Request, m *game.Manager) {
	// id := r.Context().Value("reqId")
	// log15.Debug("Request received on /: ", "id", id)
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
	// time.Sleep(6 * time.Second)

	// go func() {
	// 	select {
	// 	case <-r.Context().Done():
	// 		fmt.Println(r.Context().Err()) // prints "context deadline exceeded"
	// 		errorResponse(fmt.Errorf("too slow"), "too slow", w)
	// 		return
	// 	}
	// }()

	genericResponse(w, jsonGame, nil)
}
