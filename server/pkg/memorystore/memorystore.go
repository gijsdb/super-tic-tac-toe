package memorystore

import (
	"sync"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

type PlayerStore struct {
	idCounterMutex sync.Mutex
	idCounter      int64
	mutex          sync.Mutex
	players        map[int64]*entity.Player
}

type GameStore struct {
	idCounterMutex sync.Mutex
	idCounter      int64
	mutex          sync.Mutex
	games          map[int64]*entity.Game
}

func NewPlayerMemoryStore() *PlayerStore {
	return &PlayerStore{
		idCounter: 1,
		players:   make(map[int64]*entity.Player),
	}
}

func NewGameMemoryStore() *GameStore {
	return &GameStore{
		idCounter: 1,
		games:     make(map[int64]*entity.Game),
	}
}

func (gs *GameStore) newID() int64 {
	gs.idCounterMutex.Lock()
	defer gs.idCounterMutex.Unlock()
	id := gs.idCounter
	gs.idCounter++
	return id
}

func (ps *PlayerStore) newID() int64 {
	ps.idCounterMutex.Lock()
	defer ps.idCounterMutex.Unlock()
	id := ps.idCounter
	ps.idCounter++
	return id
}

func (ps *PlayerStore) Create(player *entity.Player) int64 {
	player.ID = ps.newID()
	ps.players[player.ID] = player

}
