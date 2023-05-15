package repository

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type GameRepositoryI interface {
	Create(game *entity.Game) *entity.Game
	Get(id int64) *entity.Game
	Index() []*entity.Game
	Update(game *entity.Game) *entity.Game
}

type GameMemoryRepository struct {
	store *memorystore.GameStore
}

func NewGameRepository() GameRepositoryI {
	return &GameMemoryRepository{
		store: memorystore.NewGameMemoryStore(),
	}
}

func (g *GameMemoryRepository) Index() []*entity.Game {
	return g.store.IndexGames()
}

func (g *GameMemoryRepository) Get(id int64) *entity.Game {
	return g.store.GetGame(id)
}

func (g *GameMemoryRepository) Create(game *entity.Game) *entity.Game {
	return g.store.CreateGame(game)
}

func (g *GameMemoryRepository) Update(game *entity.Game) *entity.Game {
	return g.store.UpdateGame(game)
}
