package game

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/inconshreveable/log15"
)

func (g *Game) Changeturn(currentPlayerTurn string) {
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

func (g *Game) RollDice(dice1 string, dice2 string) {
	g.LastRoll = fmt.Sprintf("%s,%s", dice1, dice2)
}

func (g *Game) LeaveGame(leavingPlayer int) (*Game, error) {

	// Remove player from Players on game
	playersSplit := strings.Split(g.Players, ",")
	for i, player := range playersSplit {
		playerInt, err := strconv.Atoi(player)
		if err != nil {
			log15.Crit("Error converting player to int while leaving game", "err", err)
		}
		if playerInt == leavingPlayer {
			playersSplit = append(playersSplit[:i], playersSplit[i+1:]...)
			g.Players = strings.Join(playersSplit, ",")
			break
		}
	}

	g.GameOver = true
	return g, nil
}
