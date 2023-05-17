package entity

import "time"

type Session struct {
	PlayerID string
	Expiry   time.Time
	CSRF     string
}
