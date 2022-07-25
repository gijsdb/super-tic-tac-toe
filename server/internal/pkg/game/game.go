package game

import (
	"strconv"
	"strings"

	"github.com/inconshreveable/log15"
)

func (g *Game) Changeturn(player int) {
	// get the current players
	currentPlayers := strings.Split(g.Players, ",")
	player1, err := strconv.Atoi(currentPlayers[0])
	if err != nil {
		log15.Debug("Failed to convert player to int changing turn", "err", err)
	}
	player2, err := strconv.Atoi(currentPlayers[1])
	if err != nil {
		log15.Debug("Failed to convert player to int changing turn", "err", err)
	}

	if player == player1 {
		log15.Debug("PLAYER 2 turn")
		g.PlayerTurn = strconv.Itoa(player2)
	} else {
		log15.Debug("PLAYER 1 turn", "player", player)

		g.PlayerTurn = strconv.Itoa(player1)
	}
}
