package game

import "github.com/gijsdb/super-tic-tac-toe/internal/entity"

func (s *Service) Index() []*entity.Game {
	return s.repo.Index()
}
