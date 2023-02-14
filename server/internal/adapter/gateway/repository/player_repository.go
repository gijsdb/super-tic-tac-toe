package repository

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type PlayerRepositoryI interface {
	Create(player *entity.Player) int64
	SetPlayerActive(id int64)
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

func (p *PlayerMemoryRepository) SetPlayerActive(id int64) {
	p.store.SetPlayerActive(id)
}
