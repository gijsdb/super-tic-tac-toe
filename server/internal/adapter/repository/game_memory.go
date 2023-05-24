package repository

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type GameMemoryRepoI interface {
	Create(game *entity.Game) *entity.Game
	Get(id int64) *entity.Game
	Index() []*entity.Game
	Update(game *entity.Game) *entity.Game
}

type GameMemoryRepo struct {
	store *memorystore.GameStore
}

func NewGameMemoryRepo() GameMemoryRepoI {
	return &GameMemoryRepo{
		store: memorystore.NewGameMemoryStore(),
	}
}

func (g *GameMemoryRepo) Index() []*entity.Game {
	return g.store.IndexGames()
}

func (g *GameMemoryRepo) Get(id int64) *entity.Game {
	return g.store.GetGame(id)
}

func (g *GameMemoryRepo) Create(game *entity.Game) *entity.Game {
	return g.store.CreateGame(game)
}

func (g *GameMemoryRepo) Update(game *entity.Game) *entity.Game {
	return g.store.UpdateGame(game)
}
