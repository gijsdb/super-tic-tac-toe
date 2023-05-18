package entity

type Player struct {
	ID       string `json:"ID"`
	GoogleID string `json:"GoogleID,omitempty"`
	Email    string `json:"email,omitempty"`
	Picture  string `json:"picture,omitempty"`
}
