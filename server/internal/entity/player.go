package entity

type Player struct {
	ID string `json:"ID"`
	//LastSeen int64  `json:"last_seen"`
	Email   string `json:"email,omitempty"`
	Picture string `json:"picture,omitempty"`
}
