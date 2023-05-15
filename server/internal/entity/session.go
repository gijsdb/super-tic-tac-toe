package entity

import "time"

type Session struct {
	PlayerID string    `json:"PlayerID"`
	Expiry   time.Time `json:"expiry"`
}
