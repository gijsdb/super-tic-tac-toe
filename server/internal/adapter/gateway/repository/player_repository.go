package repository

import (
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
	"github.com/gijsdb/super-tic-tac-toe/pkg/sqlite"
)

type PlayerRepositoryI interface {
	Create(player *entity.Player) string
	DBCreatePlayer(player *entity.Player) string
	Get(playerId string) (*entity.Player, error)
	DBGetWhere(value, column string) (*entity.Player, error)
	DBGetAll() ([]*entity.Player, error)
	GetByEmail(email string) (*entity.Player, error)
	Update(player *entity.Player) *entity.Player
	DBUpdate(player *entity.Player)
}

type PlayerRepository struct {
	store *memorystore.PlayerStore
	db    *sqlite.SQLiteStore
}

func NewPlayerRepository(db_path string) PlayerRepositoryI {
	return &PlayerRepository{
		store: memorystore.NewPlayerMemoryStore(),
		db:    sqlite.NewSQLiteStore(db_path),
	}
}

func (p *PlayerRepository) Create(player *entity.Player) string {
	return p.store.Create(player)
}

func (p *PlayerRepository) Get(playerId string) (*entity.Player, error) {
	player := p.store.Get(playerId)
	if player == nil {
		return nil, fmt.Errorf("could not find player for id: %s", playerId)
	}

	return player, nil
}

func (p *PlayerRepository) GetByEmail(email string) (*entity.Player, error) {
	return p.store.GetByEmail(email)
}

func (p *PlayerRepository) Update(player *entity.Player) *entity.Player {
	return p.store.Update(player)
}

func (p *PlayerRepository) DBCreatePlayer(player *entity.Player) string {
	return p.db.Create(&sqlite.Player{ID: player.ID, GoogleID: player.GoogleID, Email: player.Email, Picture: player.Picture})
}

func (p *PlayerRepository) DBGetWhere(value, column string) (*entity.Player, error) {
	db_player, err := p.db.GetWhere(value, column)
	if err != nil {
		return nil, err
	}

	return &entity.Player{
		ID:       db_player.ID,
		GoogleID: db_player.GoogleID,
		Email:    db_player.Email,
		Picture:  db_player.Picture,
		IsTemp:   false,
	}, nil
}

func (p *PlayerRepository) DBGetAll() ([]*entity.Player, error) {
	players, err := p.db.GetAllPlayers()
	if err != nil {
		return nil, err
	}

	var result []*entity.Player
	for _, player := range players {
		result = append(result, &entity.Player{
			ID:       player.ID,
			GoogleID: player.GoogleID,
			Email:    player.Email,
			Picture:  player.Picture,
			IsTemp:   false,
		})
	}

	return result, nil
}

func (p *PlayerRepository) DBUpdate(player *entity.Player) {
	p.db.UpdatePlayer(sqlite.Player{
		ID:       player.ID,
		GoogleID: player.GoogleID,
		Email:    player.Email,
		Picture:  player.Picture,
		Wins:     player.Wins,
		Losses:   player.Losses,
	})
}
