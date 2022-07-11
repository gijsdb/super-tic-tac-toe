package game

import (
	"errors"

	"github.com/inconshreveable/log15"
)

func (g *Game) createNewState() (State, error) {
	gb := CreateGameBoard()
	state := State{
		PlayerTurn: 0,
		GameOver:   false,
		Winner:     2,
		GameID:     g.ID,
		GameBoard:  &gb,
	}

	return state, nil
}

func (g *Game) JoinGame() error {
	if g.Players == 2 {
		log15.Debug("Game is full", "ID", g.ID)
		return errors.New("Game is full")
	}

	g.Players++

	return nil
}
