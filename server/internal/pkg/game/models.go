package game

// State holds the game details including state of the board
type Game struct {
	ID         int        `json:"ID"`
	GameBoard  *GameBoard `json:"game_board" ` // JSON for gameboard
	PlayerTurn int        `json:"player_turn"`
	GameOver   *GameOver  `json:"game_over"`
	Winner     int        `json:"winner"`
	Players    []int      `json:"players"` // List of string for the moment, needs changing
	Full       bool       `json:"full"`
	LastRoll   []int      `json:"last_roll"` // Last roll of the game, e.g. 5,5
}

type GameOver struct {
	Reason string `json:"reason"`
	Over   bool   `json:"over"`
}

// Gameboard represents the game board
type GameBoard struct {
	Player1 int      `json:"player_1"`
	Player2 int      `json:"player_2"`
	Winner  int      `json:"winner"`
	Squares []Square `json:"squares"`
}

// The 9 squares on the board
type Square struct {
	Circles    []Circle `json:"circles"`
	CapturedBy int      `json:"captured_by"` // 0 = player 1, 1 = player 2
	index      int
}

// The 9 circles in each square
type Circle struct {
	SelectedBy int `json:"selected_by"`
	Index      int
}
