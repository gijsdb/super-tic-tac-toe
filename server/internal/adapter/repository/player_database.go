package repository

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/sqlite"
)

type PlayerDatabaseRepoI interface {
	Create(player *entity.Player) string
	Update(player *entity.Player)
	GetWhere(value, column string) (*entity.Player, error)
	GetAll() ([]*entity.Player, error)
}

type PlayerDatabaseRepo struct {
	db *sqlite.SQLiteStore
}

func NewPlayerDatabaseRepo(db_path string) PlayerDatabaseRepoI {
	return &PlayerDatabaseRepo{
		db: sqlite.NewSQLiteStore(db_path),
	}
}

func (p *PlayerDatabaseRepo) Create(player *entity.Player) string {
	return p.db.Create(&sqlite.Player{ID: player.ID, GoogleID: player.GoogleID, Email: player.Email, Picture: player.Picture, Wins: player.Wins, Losses: player.Losses})
}

func (p *PlayerDatabaseRepo) Update(player *entity.Player) {
	p.db.UpdatePlayer(sqlite.Player{
		ID:       player.ID,
		GoogleID: player.GoogleID,
		Email:    player.Email,
		Picture:  player.Picture,
		Wins:     player.Wins,
		Losses:   player.Losses,
	})
}

func (p *PlayerDatabaseRepo) GetWhere(value, column string) (*entity.Player, error) {
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
		Wins:     db_player.Wins,
		Losses:   db_player.Losses,
	}, nil
}

func (p *PlayerDatabaseRepo) GetAll() ([]*entity.Player, error) {
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
			Wins:     player.Wins,
			Losses:   player.Losses,
		})
	}

	return result, nil
}
