package player

import "github.com/gijsdb/super-tic-tac-toe/internal/entity"

func (s *PlayerService) Get(id string) *entity.Player {
	return s.repo.Get(id)
}

func (s *PlayerService) Create() string {

	return s.repo.Create(&entity.Player{
		ID: "",
	})
}
