package game

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Overarching struct for accessing DB and Games.
type Manager struct {
	DB    *gorm.DB
	Games []Game
}

type Game struct {
	gorm.Model
	ID      int `gorm:"primary_key" json:"id"`
	Players int `json:"players"`
	// Private bool
	State State `json:"state"`
}

type State struct {
	gorm.Model
	GameBoard  datatypes.JSON `json:"game_board"`
	GameID     int            `json:"game_id"`
	PlayerTurn int            `json:"player_turn"`
	GameOver   bool           `json:"game_over"`
	Winner     int            `json:"winner"`
}
