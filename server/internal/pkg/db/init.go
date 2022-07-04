package db

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("super-tic-tac-toe.db"), &gorm.Config{})
	db.AutoMigrate(game.Game{}, game.State{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
