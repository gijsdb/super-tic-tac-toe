package repository

import (
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
	"github.com/gijsdb/super-tic-tac-toe/pkg/sqlite"
	"github.com/inconshreveable/log15"
)

type PlayerRepositoryI interface {
	Create(player *entity.Player) string
	DBCreatePlayer(player *entity.Player) uint
	Get(playerId string) (*entity.Player, error)
	DBGetWhere(value, column string) (*entity.Player, error)
	DBGetAll() ([]*entity.Player, error)
	GetByEmail(email string) (*entity.Player, error)
	Update(player *entity.Player) *entity.Player
}

// TODO: Rename to PlayerRepository since it deals with the SQLite database too?
type PlayerMemoryRepository struct {
	store *memorystore.PlayerStore
	db    *sqlite.SQLiteStore
}

func NewPlayerRepository() PlayerRepositoryI {
	return &PlayerMemoryRepository{
		store: memorystore.NewPlayerMemoryStore(),
		db:    sqlite.NewSQLiteStore("../../super_tic_tac_toe.db"),
	}
}

func (p *PlayerMemoryRepository) Create(player *entity.Player) string {
	return p.store.Create(player)
}

func (p *PlayerMemoryRepository) Get(playerId string) (*entity.Player, error) {
	player := p.store.Get(playerId)
	if player == nil {
		return nil, fmt.Errorf("could not find player for id: %s", playerId)
	}

	return player, nil
}

func (p *PlayerMemoryRepository) GetByEmail(email string) (*entity.Player, error) {
	return p.store.GetByEmail(email)
}

func (p *PlayerMemoryRepository) Update(player *entity.Player) *entity.Player {
	return p.store.Update(player)
}

func (p *PlayerMemoryRepository) DBCreatePlayer(player *entity.Player) uint {
	return p.db.Create(&sqlite.Player{OriginalID: player.ID, GoogleID: player.GoogleID, Email: player.Email, Picture: player.Picture})
}

func (p *PlayerMemoryRepository) DBGetWhere(value, column string) (*entity.Player, error) {
	db_player, err := p.db.GetWhere(value, column)
	if err != nil {
		return nil, err
	}

	return &entity.Player{
		ID:       db_player.OriginalID,
		GoogleID: db_player.GoogleID,
		Email:    db_player.Email,
		Picture:  db_player.Picture,
		IsTemp:   false,
	}, nil
}

func (p *PlayerMemoryRepository) DBGetAll() ([]*entity.Player, error) {
	players, err := p.db.GetAllPlayers()
	if err != nil {
		return nil, err
	}
	log15.Debug("found players", "ps", players)
	var result []*entity.Player
	for _, player := range players {
		result = append(result, &entity.Player{
			ID:       player.OriginalID,
			GoogleID: player.GoogleID,
			Email:    player.Email,
			Picture:  player.Picture,
			IsTemp:   false,
		})
	}

	return result, nil
}
