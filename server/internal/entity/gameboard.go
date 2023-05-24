package entity

type GameBoard struct {
	Player1 string   `json:"player_1"`
	Player2 string   `json:"player_2"`
	Winner  string   `json:"winner"`
	Squares []Square `json:"squares"`
}

// The 9 squares on the board
type Square struct {
	Circles    []Circle `json:"circles"`
	CapturedBy string   `json:"captured_by"` // 0 = player 1, 1 = player 2
	Index      int
}

// The 9 circles in each square
type Circle struct {
	SelectedBy string `json:"selected_by"`
	Index      int
}

// Check for three in a row for circles in a square
func (gb *GameBoard) CheckCirclesCondition(s Square) Square {
	var (
		x = (s.Circles[0].SelectedBy == gb.Player1 && s.Circles[1].SelectedBy == gb.Player1 && s.Circles[2].SelectedBy == gb.Player1) || // Check all rows.
			(s.Circles[3].SelectedBy == gb.Player1 && s.Circles[4].SelectedBy == gb.Player1 && s.Circles[5].SelectedBy == gb.Player1) ||
			(s.Circles[6].SelectedBy == gb.Player1 && s.Circles[7].SelectedBy == gb.Player1 && s.Circles[8].SelectedBy == gb.Player1) ||

			(s.Circles[0].SelectedBy == gb.Player1 && s.Circles[3].SelectedBy == gb.Player1 && s.Circles[6].SelectedBy == gb.Player1) || // Check all columns.
			(s.Circles[1].SelectedBy == gb.Player1 && s.Circles[4].SelectedBy == gb.Player1 && s.Circles[7].SelectedBy == gb.Player1) ||
			(s.Circles[2].SelectedBy == gb.Player1 && s.Circles[5].SelectedBy == gb.Player1 && s.Circles[8].SelectedBy == gb.Player1) ||

			(s.Circles[0].SelectedBy == gb.Player1 && s.Circles[4].SelectedBy == gb.Player1 && s.Circles[8].SelectedBy == gb.Player1) || // Check all diagonals.
			(s.Circles[2].SelectedBy == gb.Player1 && s.Circles[4].SelectedBy == gb.Player1 && s.Circles[6].SelectedBy == gb.Player1)

		o = (s.Circles[0].SelectedBy == gb.Player2 && s.Circles[1].SelectedBy == gb.Player2 && s.Circles[2].SelectedBy == gb.Player2) || // Check all rows.
			(s.Circles[3].SelectedBy == gb.Player2 && s.Circles[4].SelectedBy == gb.Player2 && s.Circles[5].SelectedBy == gb.Player2) ||
			(s.Circles[6].SelectedBy == gb.Player2 && s.Circles[7].SelectedBy == gb.Player2 && s.Circles[8].SelectedBy == gb.Player2) ||

			(s.Circles[0].SelectedBy == gb.Player2 && s.Circles[3].SelectedBy == gb.Player2 && s.Circles[6].SelectedBy == gb.Player2) || // Check all collumns.
			(s.Circles[1].SelectedBy == gb.Player2 && s.Circles[4].SelectedBy == gb.Player2 && s.Circles[7].SelectedBy == gb.Player2) ||
			(s.Circles[2].SelectedBy == gb.Player2 && s.Circles[5].SelectedBy == gb.Player2 && s.Circles[8].SelectedBy == gb.Player2) ||

			(s.Circles[0].SelectedBy == gb.Player2 && s.Circles[4].SelectedBy == gb.Player2 && s.Circles[8].SelectedBy == gb.Player2) || // Check all diagonals.
			(s.Circles[2].SelectedBy == gb.Player2 && s.Circles[4].SelectedBy == gb.Player2 && s.Circles[6].SelectedBy == gb.Player2)

		freeCellsLeft = s.Circles[0].SelectedBy == "" || s.Circles[1].SelectedBy == "" || s.Circles[2].SelectedBy == "" ||
			s.Circles[3].SelectedBy == "" || s.Circles[4].SelectedBy == "" || s.Circles[5].SelectedBy == "" ||
			s.Circles[6].SelectedBy == "" || s.Circles[7].SelectedBy == "" || s.Circles[8].SelectedBy == ""
	)
	switch {
	case x && !o:
		s.CapturedBy = gb.Player1
		return s
	case o && !x:
		s.CapturedBy = gb.Player2
		return s
	case !freeCellsLeft:
		s.CapturedBy = ""
		return s
	default:
		return s
	}
}

// Check for three in a row for squares
func (gb *GameBoard) CheckSquareCondition() {

	var (
		x = (gb.Squares[0].CapturedBy == gb.Player1 && gb.Squares[1].CapturedBy == gb.Player1 && gb.Squares[2].CapturedBy == gb.Player1) || // Check all rows.
			(gb.Squares[3].CapturedBy == gb.Player1 && gb.Squares[4].CapturedBy == gb.Player1 && gb.Squares[5].CapturedBy == gb.Player1) ||
			(gb.Squares[6].CapturedBy == gb.Player1 && gb.Squares[7].CapturedBy == gb.Player1 && gb.Squares[8].CapturedBy == gb.Player1) ||

			(gb.Squares[0].CapturedBy == gb.Player1 && gb.Squares[3].CapturedBy == gb.Player1 && gb.Squares[6].CapturedBy == gb.Player1) || // Check all columns.
			(gb.Squares[1].CapturedBy == gb.Player1 && gb.Squares[4].CapturedBy == gb.Player1 && gb.Squares[7].CapturedBy == gb.Player1) ||
			(gb.Squares[2].CapturedBy == gb.Player1 && gb.Squares[5].CapturedBy == gb.Player1 && gb.Squares[8].CapturedBy == gb.Player1) ||

			(gb.Squares[0].CapturedBy == gb.Player1 && gb.Squares[4].CapturedBy == gb.Player1 && gb.Squares[8].CapturedBy == gb.Player1) || // Check all diagonals.
			(gb.Squares[2].CapturedBy == gb.Player1 && gb.Squares[4].CapturedBy == gb.Player1 && gb.Squares[6].CapturedBy == gb.Player1)

		o = (gb.Squares[0].CapturedBy == gb.Player2 && gb.Squares[1].CapturedBy == gb.Player2 && gb.Squares[2].CapturedBy == gb.Player2) || // Check all rgb.Player2ws.
			(gb.Squares[3].CapturedBy == gb.Player2 && gb.Squares[4].CapturedBy == gb.Player2 && gb.Squares[5].CapturedBy == gb.Player2) ||
			(gb.Squares[6].CapturedBy == gb.Player2 && gb.Squares[7].CapturedBy == gb.Player2 && gb.Squares[8].CapturedBy == gb.Player2) ||

			(gb.Squares[0].CapturedBy == gb.Player2 && gb.Squares[3].CapturedBy == gb.Player2 && gb.Squares[6].CapturedBy == gb.Player2) || // Check all cgb.Player2lumns.
			(gb.Squares[1].CapturedBy == gb.Player2 && gb.Squares[4].CapturedBy == gb.Player2 && gb.Squares[7].CapturedBy == gb.Player2) ||
			(gb.Squares[2].CapturedBy == gb.Player2 && gb.Squares[5].CapturedBy == gb.Player2 && gb.Squares[8].CapturedBy == gb.Player2) ||

			(gb.Squares[0].CapturedBy == gb.Player2 && gb.Squares[4].CapturedBy == gb.Player2 && gb.Squares[8].CapturedBy == gb.Player2) || // Check all diaggb.Player2nals.
			(gb.Squares[2].CapturedBy == gb.Player2 && gb.Squares[4].CapturedBy == gb.Player2 && gb.Squares[6].CapturedBy == gb.Player2)

		freeCellsLeft = gb.Squares[0].CapturedBy == "" || gb.Squares[1].CapturedBy == "" || gb.Squares[2].CapturedBy == "" ||
			gb.Squares[3].CapturedBy == "" || gb.Squares[4].CapturedBy == "" || gb.Squares[5].CapturedBy == "" ||
			gb.Squares[6].CapturedBy == "" || gb.Squares[7].CapturedBy == "" || gb.Squares[8].CapturedBy == ""
	)

	switch {
	case x && !o:
		gb.Winner = gb.Player1
		return
	case o && !x:
		gb.Winner = gb.Player1
		return
	case !freeCellsLeft:
		return
	default:
		return
	}
}
