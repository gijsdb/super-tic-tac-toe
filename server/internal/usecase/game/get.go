package game

import "github.com/gijsdb/super-tic-tac-toe/internal/entity"

func (s *Service) Get(id int64) *entity.Game {
	return s.repo.Get(id)
}
