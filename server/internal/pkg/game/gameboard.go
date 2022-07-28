package game

import "github.com/inconshreveable/log15"

// var (
// 	Player1    = -1
// 	Player2    = -1
// 	Unselected = -1
// )

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

func CreateGameBoard(creatingPlayer int) GameBoard {
	var board GameBoard
	board.Player1 = creatingPlayer
	board.Winner = -1
	var squares []Square
	for i := 0; i < 9; i++ {
		var square Square
		square = Square{
			Circles:    initCircles(),
			CapturedBy: -1,
			index:      i,
		}
		squares = append(squares, square)
	}
	board.Squares = squares
	return board
}

func (gb *GameBoard) Update(player, square, circle int) error {
	// TODO add check here to see if its not selected by another player
	gb.Squares[square].Circles[circle].SelectedBy = player
	updatedSquare := gb.checkCirclesCondition(gb.Squares[square])
	gb.Squares[square] = updatedSquare
	squareWin := gb.checkSquareCondition()

	log15.Debug("square", "square", squareWin)
	return nil
}

func initCircles() []Circle {
	var circles []Circle

	for i := 0; i < 9; i++ {
		var circle Circle
		circle.Index = i
		circle.SelectedBy = -1

		circles = append(circles, circle)
	}

	return circles
}

// Check for three in a row for circles in a square
func (gb *GameBoard) checkCirclesCondition(s Square) Square {

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

		freeCellsLeft = s.Circles[0].SelectedBy == -1 || s.Circles[1].SelectedBy == -1 || s.Circles[2].SelectedBy == -1 ||
			s.Circles[3].SelectedBy == -1 || s.Circles[4].SelectedBy == -1 || s.Circles[5].SelectedBy == -1 ||
			s.Circles[6].SelectedBy == -1 || s.Circles[7].SelectedBy == -1 || s.Circles[8].SelectedBy == -1
	)
	switch {
	case x && !o:
		log15.Debug("Player 1 has captured a square")
		s.CapturedBy = gb.Player1
		return s
	case o && !x:
		log15.Debug("Player 2 has captured a square")
		s.CapturedBy = gb.Player2
		return s
	case !freeCellsLeft:
		log15.Debug("Freecells false in checkCircleCOndition")
		s.CapturedBy = -1
		return s
	default:
		log15.Debug("Defaulted in checkCircleCOndition")

		return s
	}
}

// Check for three in a row for squares
func (gb *GameBoard) checkSquareCondition() int {

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

		freeCellsLeft = gb.Squares[0].CapturedBy == -1 || gb.Squares[1].CapturedBy == -1 || gb.Squares[2].CapturedBy == -1 ||
			gb.Squares[3].CapturedBy == -1 || gb.Squares[4].CapturedBy == -1 || gb.Squares[5].CapturedBy == -1 ||
			gb.Squares[6].CapturedBy == -1 || gb.Squares[7].CapturedBy == -1 || gb.Squares[8].CapturedBy == -1
	)

	switch {
	case x && !o:
		gb.Winner = gb.Player1
		return gb.Player1
	case o && !x:
		gb.Winner = gb.Player1
		return gb.Player1
	case !freeCellsLeft:
		return -1
	default:
		return -1
	}

}
