package entity

// Gameboard represents the game board
type GameBoard struct {
	Player1 int64    `json:"player_1"`
	Player2 int64    `json:"player_2"`
	Winner  int64    `json:"winner"`
	Squares []Square `json:"squares"`
}

// The 9 squares on the board
type Square struct {
	Circles    []Circle `json:"circles"`
	CapturedBy int64    `json:"captured_by"` // 0 = player 1, 1 = player 2
	Index      int
}

// The 9 circles in each square
type Circle struct {
	SelectedBy int64 `json:"selected_by"`
	Index      int
}
