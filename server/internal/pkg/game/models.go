package game

import (
	"gorm.io/gorm"
)

// Manager is an overarching struct for accessing DB and Games.
type Manager struct {
	DB      *gorm.DB
	Games   []*Game      // List of games currently created (happening)
	Players map[int]bool // List of clients AKA players connected. int is id, bool is active
	// Lock    sync.Mutex // May need this for players
}

// State holds the game details including state of the board
type Game struct {
	gorm.Model
	ID         int        `gorm:"primary_key"`
	GameBoard  *GameBoard `json:"game_board" gorm:"-:all"` // JSON for gameboard
	GameID     int        `json:"game_id"`
	PlayerTurn int        `json:"player_turn"`
	GameOver   bool       `json:"game_over"`
	Winner     int        `json:"winner"`
	Players    string     `json:"players"`
}
