package game

import "github.com/inconshreveable/log15"

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

func (gb *GameBoard) Update(player, square, circle int) {
	gb.Squares[square].Circles[circle].SelectedBy = player
	updatedSquare := gb.checkCirclesCondition(gb.Squares[square])
	gb.Squares[square] = updatedSquare
	gb.checkSquareCondition()

	return
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
		s.CapturedBy = -1
		return s
	default:
		return s
	}
}

// Check for three in a row for squares
func (gb *GameBoard) checkSquareCondition() {

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
