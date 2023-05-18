package entity

type Player struct {
	ID       string `json:"ID"`
	GoogleID string `json:"GoogleID,omitempty"`
	//LastSeen int64  `json:"last_seen"`
	Email   string `json:"email,omitempty"`
	Picture string `json:"picture,omitempty"`
}
