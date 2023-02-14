package repository

import "github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"

type GameInstanceRepositoryI interface {
	// Create(int)
	// Update(int, int, int)
	// RemoveCircle(int, int, int)
}

type GameInstanceMemoryStore struct {
	store *memorystore.GameStore
}

func NewGameInstanceRepository(store *memorystore.GameStore) GameRepositoryI {
	return &GameMemoryRepository{
		store: store,
	}
}
