package player

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/inconshreveable/log15"
)

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

func (s *PlayerService) LoadDBPlayersIntoMemory() {
	db_players, err := s.repo.DBGetAll()
	if err != nil {
		log15.Info("No players found to load into memory")
		return
	}

	for _, db_player := range db_players {
		player_id := s.repo.Create(db_player)
		log15.Debug("loaded player into memory", "p", player_id)
	}
}
