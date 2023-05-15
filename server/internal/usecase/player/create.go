package player

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *PlayerService) Create() string {

	return s.repo.Create(&entity.Player{
		ID: "",
	})
}
