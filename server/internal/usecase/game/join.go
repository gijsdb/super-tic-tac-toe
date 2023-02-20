package game

import (
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *Service) Join(gameId int64, joiningPlayer string) (*entity.Game, error) {
	game := s.repo.Get(gameId)
	noOfPlayers := len(game.Players)
	if noOfPlayers >= 3 {
		return nil, fmt.Errorf("game is full")
	}
	game.Players = append(game.Players, joiningPlayer)
	game.GameBoard.Player2 = joiningPlayer
	game.Full = true

	return s.repo.Update(game), nil
}
