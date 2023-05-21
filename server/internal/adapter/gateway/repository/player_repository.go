package repository

import (
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type PlayerRepositoryI interface {
	Create(player *entity.Player) string
	Get(playerId string) (*entity.Player, error)
	GetByEmail(email string) (*entity.Player, error)
	Update(player *entity.Player) *entity.Player
}

type PlayerMemoryRepository struct {
	store *memorystore.PlayerStore
}

func NewPlayerRepository() PlayerRepositoryI {
	return &PlayerMemoryRepository{
		store: memorystore.NewPlayerMemoryStore(),
	}
}

func (p *PlayerMemoryRepository) Create(player *entity.Player) string {
	return p.store.Create(player)
}

func (p *PlayerMemoryRepository) Get(playerId string) (*entity.Player, error) {
	player := p.store.Get(playerId)
	if player == nil {
		return nil, fmt.Errorf("could not find player for id: %s", playerId)
	}

	return player, nil
}

func (p *PlayerMemoryRepository) GetByEmail(email string) (*entity.Player, error) {
	return p.store.GetByEmail(email)
}

func (p *PlayerMemoryRepository) Update(player *entity.Player) *entity.Player {
	return p.store.Update(player)
}
