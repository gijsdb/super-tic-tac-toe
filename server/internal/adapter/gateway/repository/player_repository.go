package repository

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type PlayerRepositoryI interface {
	Create(player *entity.Player) string
	Get(playerId string) *entity.Player
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

func (p *PlayerMemoryRepository) Get(playerId string) *entity.Player {
	return p.store.Get(playerId)
}

func (p *PlayerMemoryRepository) Update(player *entity.Player) *entity.Player {
	return p.store.Update(player)
}
