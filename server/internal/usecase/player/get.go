package player

import "github.com/gijsdb/super-tic-tac-toe/internal/entity"

func (s *PlayerService) Get(id string) *entity.Player {
	return s.repo.Get(id)
}
