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

func NewPlayerMemoryStore() *PlayerStore {
	return &PlayerStore{
		idCounter: 1,
		players:   make(map[int64]*entity.Player),
	}
}

func (ps *PlayerStore) NewID() int64 {
	ps.idCounterMutex.Lock()
	defer ps.idCounterMutex.Unlock()
	id := ps.idCounter
	ps.idCounter++
	return id
}

func (ps *PlayerStore) Create(player *entity.Player) int64 {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	player.ID = ps.NewID()
	ps.players[player.ID] = player

	return player.ID
}

func (ps *PlayerStore) Get(playerId int64) *entity.Player {
	return ps.players[playerId]
}

func (ps *PlayerStore) Update(player *entity.Player) *entity.Player {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	ps.players[player.ID] = player
	return ps.players[player.ID]
}
