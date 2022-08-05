package game

import (
	"github.com/inconshreveable/log15"
)

func (g *Game) Changeturn(currentPlayerTurn int) {

	// log15.Debug("Changing turn", "currentPlayerTurn", currentPlayerTurn, "game player 1", currentPlayers[0], "game player 2", currentPlayers[1])
	if currentPlayerTurn == g.Players[0] {
		g.PlayerTurn = g.Players[1]
	} else if currentPlayerTurn == g.Players[1] {
		g.PlayerTurn = g.Players[0]
	} else {
		log15.Crit("Player not found in ChangeTurn()::game.go", "requesting player", currentPlayerTurn)
	}
}

func (g *Game) RollDice(dice1 int, dice2 int) {
	g.LastRoll = []int{dice1, dice2}
}

func (g *Game) LeaveGame(leavingPlayer int) (*Game, error) {

	// Remove player from Players on game
	for i, player := range g.Players {
		if player == leavingPlayer {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
		}
	}

	g.SetGameOver("Player left the game")
	return g, nil
}

func (g *Game) SetGameOver(reason string) {
	g.GameOver.Over = true
	g.GameOver.Reason = reason
}
