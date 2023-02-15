package repository

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type PlayerRepositoryI interface {
	Create(player *entity.Player) int64
	Get(playerId int64) *entity.Player
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

func (p *PlayerMemoryRepository) Create(player *entity.Player) int64 {
	return p.store.Create(player)
}

func (p *PlayerMemoryRepository) Get(playerId int64) *entity.Player {
	return p.store.Get(playerId)
}

func (p *PlayerMemoryRepository) Update(player *entity.Player) *entity.Player {
	return p.store.Update(player)
}
