package entity

type Game struct {
	ID         int64      `json:"ID"`
	GameBoard  *GameBoard `json:"game_board" ` // JSON for gameboard
	PlayerTurn int        `json:"player_turn"`
	GameOver   *GameOver  `json:"game_over"`
	Winner     int64      `json:"winner"`
	Players    []int      `json:"players"` // List of string for the moment, needs changing
	Full       bool       `json:"full"`
	LastRoll   []int      `json:"last_roll"` // Last roll of the game, e.g. 5,5
}

type GameOver struct {
	Reason string `json:"reason"`
	Over   bool   `json:"over"`
}
