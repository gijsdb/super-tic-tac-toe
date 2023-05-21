package memorystore

import (
	"fmt"
	"sync"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

type PlayerStore struct {
	mutex   sync.Mutex
	players map[string]*entity.Player // Indexed by ID
}

func NewPlayerMemoryStore() *PlayerStore {
	return &PlayerStore{
		players: make(map[string]*entity.Player),
	}
}

func (ps *PlayerStore) Create(player *entity.Player) string {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	player.ID = NewUUID()
	ps.players[player.ID] = player

	return player.ID
}

func (ps *PlayerStore) Get(playerId string) *entity.Player {
	return ps.players[playerId]
}

func (ps *PlayerStore) GetByEmail(email string) (*entity.Player, error) {
	for _, player := range ps.players {
		if player.Email == email {
			return player, nil
		}
	}

	return nil, fmt.Errorf("can't find player by provided email %s", email)
}

func (ps *PlayerStore) Update(player *entity.Player) *entity.Player {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	ps.players[player.ID] = player
	return ps.players[player.ID]
}
