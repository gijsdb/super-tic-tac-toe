package player

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *Service) CreatePlayer() {
	player := &entity.Player{
		ID:     -1,
		Active: true,
	}

	s.repo.Create(player)
}
