package entity

type Player struct {
	ID       string `json:"ID"`
	GoogleID string `json:"google_id,omitempty"`
	Email    string `json:"email,omitempty"`
	Picture  string `json:"picture,omitempty"`
	IsTemp   bool   `json:"is_temp"`
}
