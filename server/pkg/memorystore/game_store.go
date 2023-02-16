package memorystore

import (
	"sync"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/go-co-op/gocron"
	"github.com/inconshreveable/log15"
)

type GameStore struct {
	idCounterMutex sync.Mutex
	idCounter      int64
	mutex          sync.Mutex
	games          map[int64]*entity.Game
}

func NewGameMemoryStore() *GameStore {
	gs := &GameStore{
		idCounter: 1,
		games:     make(map[int64]*entity.Game),
	}

	go gs.RemoveEndedGames()

	return gs
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
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	gs.games[game.ID] = game
	return gs.games[game.ID]
}

func (gs *GameStore) RemoveEndedGames() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Minute().Do(func() {
		log15.Info("Checking for games to remove")
		for _, game := range gs.games {
			if game.GameOver.Over {
				now := time.Now().Unix()
				// check game has been over for more than 1 minute
				if now-game.GameOver.EndTime > 60 {
					log15.Info("Removing game", "ID", game.ID)
					gs.mutex.Lock()
					delete(gs.games, game.ID)
					gs.mutex.Unlock()
				}
			}
		}
	})

	s.StartBlocking()
}
