package repository

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type GameRepositoryI interface {
	Create(game *entity.Game) *entity.Game
	// Get() (*entity.Game, error)
	Index() []*entity.Game
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

func (g GameMemoryRepository) Index() []*entity.Game {
	return g.store.IndexGames()
}

func (g GameMemoryRepository) Create(game *entity.Game) *entity.Game {
	return g.store.CreateGame(game)
}
