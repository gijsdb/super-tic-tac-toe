package game

import (
	"strconv"
	"strings"

	"github.com/inconshreveable/log15"
)

func (g *Game) Changeturn(currentPlayerTurn string) {
	// get the current players
	currentPlayers := strings.Split(g.Players, ",")

	for _, player := range currentPlayers {
		player = strings.Trim(player, " ")
	}
	// log15.Debug("Changing turn", "currentPlayerTurn", currentPlayerTurn, "game player 1", currentPlayers[0], "game player 2", currentPlayers[1])
	if currentPlayerTurn == currentPlayers[0] {
		g.PlayerTurn = currentPlayers[1]
	} else if currentPlayerTurn == currentPlayers[1] {
		g.PlayerTurn = currentPlayers[0]
	} else {
		log15.Crit("Player not found in ChangeTurn()::game.go", "requesting player", currentPlayerTurn)
	}
}

// GetPlayers Returns csv string as list of ints
func (g *Game) GetPlayersAsListInt() []int {
	var players []int

	playersSplit := strings.Split(g.Players, ",")
	for _, player := range playersSplit {
		player = strings.Trim(player, " ")
		playerInt, _ := strconv.Atoi(player)
		players = append(players, playerInt)
	}

	return players
}
