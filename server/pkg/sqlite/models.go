package sqlite

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	ID        string `gorm:"primaryKey"` // Represents the ID set in memory by the game when the account was created
	GoogleID  string
	Email     string
	Picture   string
	Wins      int
	Losses    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
