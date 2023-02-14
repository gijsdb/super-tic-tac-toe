package entity

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
