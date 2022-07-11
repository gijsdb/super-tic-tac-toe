package game

import (
	"gorm.io/gorm"
)

// Manager is an overarching struct for accessing DB and Games.
type Manager struct {
	DB    *gorm.DB
	Games map[int]Game // List of games currently created (happening)
}

// A game is shown on the UI and is "joinable"
type Game struct {
	gorm.Model
	ID      int `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Players int `json:"players"`
	// Private bool
	State *State `json:"state" gorm:"foreign_key:StateID"`
}

// State holds the game details including state of the board
type State struct {
	gorm.Model
	GameBoard  *GameBoard `json:"game_board" gorm:"-:all"` // JSON for gameboard
	GameID     int        `json:"game_id"`
	PlayerTurn int        `json:"player_turn"`
	GameOver   bool       `json:"game_over"`
	Winner     int        `json:"winner"`
}
