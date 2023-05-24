package player

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/inconshreveable/log15"
)

func (s *PlayerService) Get(id string) (*entity.Player, error) {
	return s.memorystore.Get(id)
}

func (s *PlayerService) GetByEmail(email string) (*entity.Player, error) {
	return s.memorystore.GetByEmail(email)
}

func (s *PlayerService) Create() string {

	return s.memorystore.Create(&entity.Player{
		ID:     "",
		IsTemp: true,
	})
}

func (s *PlayerService) Update(player *entity.Player) *entity.Player {
	return s.memorystore.Update(player)
}

func (s *PlayerService) LoadDBPlayersIntoMemory() {
	log15.Info("Loading players from database into memory")
	db_players, err := s.database.GetAll()
	if err != nil {
		log15.Info("No players found to load into memory")
		return
	}

	for _, db_player := range db_players {
		s.memorystore.Create(db_player)
	}
}

func (s *PlayerService) GetHighscores() ([]*entity.Player, error) {
	return s.database.GetAll()
}

func (s *PlayerService) SeedPlayers() {
	log15.Info("Seeding database with users")
	for i := 0; i < 15; i++ {
		rand.Seed(time.Now().UnixNano())
		player := &entity.Player{
			ID:       fmt.Sprintf("%d-seed", i),
			GoogleID: fmt.Sprintf("%d-seed", i),
			Email:    fmt.Sprintf("%d-seed@gmail.com", i),
			IsTemp:   false,
			Wins:     rand.Intn(20-1) + 1,
			Losses:   rand.Intn(20-1) + 1,
		}
		s.database.Create(player)
	}
}
