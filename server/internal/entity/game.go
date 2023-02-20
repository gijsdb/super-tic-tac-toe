package entity

type Game struct {
	ID         int64      `json:"ID"`
	GameBoard  *GameBoard `json:"game_board" ` // JSON for gameboard
	PlayerTurn string     `json:"player_turn"`
	GameOver   *GameOver  `json:"game_over"`
	Winner     string     `json:"winner"`
	Players    []string   `json:"players"` // List of string for the moment, needs changing
	Full       bool       `json:"full"`
	LastRoll   []int      `json:"last_roll"` // Last roll of the game, e.g. 5,5
}

type GameOver struct {
	Reason  string `json:"reason"`
	Over    bool   `json:"over"`
	EndTime int64  `json:"endtime"`
}
