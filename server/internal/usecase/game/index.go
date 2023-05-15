package game

import "github.com/gijsdb/super-tic-tac-toe/internal/entity"

func (s *GameService) Index() []*entity.Game {
	return s.repo.Index()
}
