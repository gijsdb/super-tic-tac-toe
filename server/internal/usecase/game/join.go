package game

import (
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *Service) Join(gameId, joiningPlayer int64) (*entity.Game, error) {
	// Find game thats being joined
	game := s.repo.Get(gameId)
	// Assign player joining to that game and gameboard
	noOfPlayers := len(game.Players)
	if noOfPlayers >= 3 {
		return nil, fmt.Errorf("game is full")
	}
	game.Players = append(game.Players, int(joiningPlayer))
	game.GameBoard.Player2 = joiningPlayer
	game.Full = true

	return s.repo.Update(game), nil
}
