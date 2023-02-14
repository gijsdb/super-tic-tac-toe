package repository

import (
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type GameRepositoryI interface {
	// Create(*entity.Game) error
	// Get() (*entity.Game, error)
	// Index() ([]*entity.Game, error)
	// Changeturn()
	// RollDice(int, int)
	// RemovePlayer(int)
	// EndGame(string)
}

type GameMemoryRepository struct {
	store *memorystore.GameStore
}

func NewGameRepository(store *memorystore.GameStore) GameRepositoryI {
	return &GameMemoryRepository{
		store: store,
	}
}
