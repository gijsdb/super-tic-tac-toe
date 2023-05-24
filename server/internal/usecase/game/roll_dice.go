package game

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *GameService) RollDice(dice1, dice2 int, gameId int64) *entity.Game {
	game := s.g_mem_store.Get(gameId)
	game.LastRoll = []int{dice1, dice2}

	return s.g_mem_store.Update(game)
}
