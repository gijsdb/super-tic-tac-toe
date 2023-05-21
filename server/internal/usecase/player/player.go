package player

import "github.com/gijsdb/super-tic-tac-toe/internal/entity"

func (s *PlayerService) Get(id string) (*entity.Player, error) {
	return s.repo.Get(id)
}

func (s *PlayerService) GetByEmail(email string) (*entity.Player, error) {
	return s.repo.GetByEmail(email)
}

func (s *PlayerService) Create() string {

	return s.repo.Create(&entity.Player{
		ID:     "",
		IsTemp: true,
	})
}

func (s *PlayerService) Update(player *entity.Player) *entity.Player {
	return s.repo.Update(player)
}
