package repository

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type PlayerRepositoryI interface {
	Create(*entity.Player)
}

type PlayerMemoryRepository struct {
	store *memorystore.PlayerStore
}

func NewPlayerRepository() PlayerRepositoryI {
	return &PlayerMemoryRepository{
		store: memorystore.NewPlayerMemoryStore(),
	}
}

func (p *PlayerMemoryRepository) Create(player *entity.Player) {
	p.store.Create(player)
}
