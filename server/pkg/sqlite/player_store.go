package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteStore struct {
	db *gorm.DB
}

func NewSQLiteStore() *SQLiteStore {
	db, err := gorm.Open(sqlite.Open("../../super_tic_tac_toe.db"), &gorm.Config{})
	if err != nil {
		panic("failed to create SQLite database")
	}

	err = db.AutoMigrate(&Player{})
	if err != nil {
		panic("failed to auto migrate SQLite database")
	}

	return &SQLiteStore{
		db: db,
	}
}

func (ss *SQLiteStore) Create(player *Player) uint {
	ss.db.Create(player)

	return player.ID
}
