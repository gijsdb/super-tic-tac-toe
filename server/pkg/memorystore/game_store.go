package memorystore

import (
	"sync"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

type GameStore struct {
	idCounterMutex sync.Mutex
	idCounter      int64
	mutex          sync.Mutex
	games          map[int64]*entity.Game
}

func NewGameMemoryStore() *GameStore {
	return &GameStore{
		idCounter: 1,
		games:     make(map[int64]*entity.Game),
	}
}

func (gs *GameStore) NewID() int64 {
	gs.idCounterMutex.Lock()
	defer gs.idCounterMutex.Unlock()
	id := gs.idCounter
	gs.idCounter++
	return id
}

func (gs *GameStore) IndexGames() []*entity.Game {
	games := []*entity.Game{}

	for _, g := range gs.games {
		games = append(games, g)
	}

	return games
}

func (gs *GameStore) GetGame(id int64) *entity.Game {
	return gs.games[id]
}

func (gs *GameStore) CreateGame(game *entity.Game) *entity.Game {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	game.ID = gs.NewID()
	gs.games[game.ID] = game

	return game
}

func (gs *GameStore) UpdateGame(game *entity.Game) *entity.Game {
	gs.games[game.ID] = game
	return gs.games[game.ID]
}
