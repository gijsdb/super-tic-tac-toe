package sqlite

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	OriginalID string // Represents the ID set in memory by the game when the account was created
	GoogleID   string
	Email      string
	Picture    string
}
