package repository

import (
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type PlayerMemoryRepoI interface {
	Create(player *entity.Player) string
	Get(playerId string) (*entity.Player, error)
	GetByEmail(email string) (*entity.Player, error)
	Update(player *entity.Player) *entity.Player
}

type PlayerMemoryRepo struct {
	store *memorystore.PlayerStore
}

func NewPlayerMemoryRepo() PlayerMemoryRepoI {
	return &PlayerMemoryRepo{
		store: memorystore.NewPlayerMemoryStore(),
	}
}

func (p *PlayerMemoryRepo) Create(player *entity.Player) string {
	return p.store.Create(player)
}

func (p *PlayerMemoryRepo) Get(playerId string) (*entity.Player, error) {
	player := p.store.Get(playerId)
	if player == nil {
		return nil, fmt.Errorf("could not find player for id: %s", playerId)
	}

	return player, nil
}

func (p *PlayerMemoryRepo) GetByEmail(email string) (*entity.Player, error) {
	return p.store.GetByEmail(email)
}

func (p *PlayerMemoryRepo) Update(player *entity.Player) *entity.Player {
	return p.store.Update(player)
}
