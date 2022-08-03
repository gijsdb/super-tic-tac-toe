// UNUSED BUT MAYBE USEFUL LATER
package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Ins *gorm.DB
}

var ins *DB

/*
// db := db.Get()
	// result := db.Ins.Create(&game)
	// if result.Error != nil {
	// 	log15.Debug("Error saving new game CreateNewGame()::game.go", "err", result.Error)
	// 	return nil, result.Error
	// }
*/

// db.AutoMigrate(Game{})
func Get() *DB {
	if ins == nil {
		db, err := gorm.Open(sqlite.Open("super-tic-tac-toe.db"), &gorm.Config{})
		ins = &DB{
			Ins: db,
		}

		if err != nil {
			panic("failed to connect database")
		}

		return ins
	}
	return ins
}

func (db *DB) Clear() {
	db.Ins.Exec("DELETE FROM games")
}
