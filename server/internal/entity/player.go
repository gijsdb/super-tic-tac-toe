package entity

type Player struct {
	ID     string `json:"ID"`
	Active bool   `json:"active"`
	//LastSeen int64  `json:"last_seen"`
	Email   string `json:"email,omitempty"`
	Picture string `json:"picture,omitempty"`
}

// type Session struct {
// 	PlayerID string    `json:"PlayerID"`
// 	Expiry   time.Time `json:"expiry"`
// }
