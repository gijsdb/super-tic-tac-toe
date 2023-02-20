package entity

type Player struct {
	ID       string `json:"ID"`
	Active   bool   `json:"active"`
	LastSeen int64  `json:"last_seen"`
}
