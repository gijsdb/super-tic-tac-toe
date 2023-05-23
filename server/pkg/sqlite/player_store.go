package sqlite

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteStore struct {
	db *gorm.DB
}

func NewSQLiteStore(db_path string) *SQLiteStore {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
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

// Can make more generic by passing the entity we want to retrieve
func (ss *SQLiteStore) GetWhere(value string, column string) (Player, error) {
	player := &Player{}
	if err := ss.db.Where(fmt.Sprintf("%s = ?", column), value).First(player).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Player{}, fmt.Errorf("no player found for value: %s in column %s", value, column)
		} else {
			return Player{}, fmt.Errorf("error retrieving player: %s", err.Error())
		}
	}

	return *player, nil
}

func (ss *SQLiteStore) GetAllPlayers() ([]Player, error) {
	var players []Player

	result := ss.db.Find(&players)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no records found in player table")
	}

	return players, nil
}
