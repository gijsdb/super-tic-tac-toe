package game

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *Service) RollDice(dice1, dice2 int, gameId int64) *entity.Game {
	game := s.repo.Get(gameId)
	game.LastRoll = []int{dice1, dice2}

	return s.repo.Update(game)
}
