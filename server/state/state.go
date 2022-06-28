package state

import "github.com/inconshreveable/log15"

const (
	Player1    = 0
	Player2    = 1
	Unselected = 2
)

type State struct {
	GameBoard  GameBoard `json:"game_board"`
	PlayerTurn int       `json:"player_turn"`
	GameOver   bool      `json:"game_over"`
	Winner     int       `json:"winner"`
}

// The whole board
type GameBoard struct {
	Squares []Square `json:"squares"`
}

// The 9 squares on the board
type Square struct {
	Circles    []Circle `json:"circles"`
	CapturedBy int      `json:"captured_by"` // 0 = player 1, 1 = player 2
}

// The circles in each square
type Circle struct {
	SelectedBy int `json:"selected_by"`
	Index      int
}

func InitGame() State {
	state := State{
		GameBoard:  initGameBoard(),
		PlayerTurn: 0,
		GameOver:   false,
	}

	return state
}

func initGameBoard() GameBoard {
	var board GameBoard
	var squares []Square
	for i := 0; i < 9; i++ {
		var square Square
		square = Square{
			Circles:    initCircles(),
			CapturedBy: Unselected,
		}
		squares = append(squares, square)
	}
	board.Squares = squares
	return board
}

func initCircles() []Circle {
	var circles []Circle

	for i := 0; i < 9; i++ {
		var circle Circle
		circle.Index = i
		circle.SelectedBy = Unselected

		circles = append(circles, circle)
	}

	return circles
}

// Check for three in a row for square
func (s *Square) CheckCirclesCondition() error {

	var (
		x = (s.Circles[0].SelectedBy == 0 && s.Circles[1].SelectedBy == 0 && s.Circles[2].SelectedBy == 0) || // Check all rows.
			(s.Circles[3].SelectedBy == 0 && s.Circles[4].SelectedBy == 0 && s.Circles[5].SelectedBy == 0) ||
			(s.Circles[6].SelectedBy == 0 && s.Circles[7].SelectedBy == 0 && s.Circles[8].SelectedBy == 0) ||

			(s.Circles[0].SelectedBy == 0 && s.Circles[3].SelectedBy == 0 && s.Circles[6].SelectedBy == 0) || // Check all columns.
			(s.Circles[1].SelectedBy == 0 && s.Circles[4].SelectedBy == 0 && s.Circles[7].SelectedBy == 0) ||
			(s.Circles[2].SelectedBy == 0 && s.Circles[5].SelectedBy == 0 && s.Circles[8].SelectedBy == 0) ||

			(s.Circles[0].SelectedBy == 0 && s.Circles[4].SelectedBy == 0 && s.Circles[8].SelectedBy == 0) || // Check all diagonals.
			(s.Circles[2].SelectedBy == 0 && s.Circles[4].SelectedBy == 0 && s.Circles[6].SelectedBy == 0)

		o = (s.Circles[0].SelectedBy == 1 && s.Circles[1].SelectedBy == 1 && s.Circles[2].SelectedBy == 1) || // Check all r1ws.
			(s.Circles[3].SelectedBy == 1 && s.Circles[4].SelectedBy == 1 && s.Circles[5].SelectedBy == 1) ||
			(s.Circles[6].SelectedBy == 1 && s.Circles[7].SelectedBy == 1 && s.Circles[8].SelectedBy == 1) ||

			(s.Circles[0].SelectedBy == 1 && s.Circles[3].SelectedBy == 1 && s.Circles[6].SelectedBy == 1) || // Check all c1lumns.
			(s.Circles[1].SelectedBy == 1 && s.Circles[4].SelectedBy == 1 && s.Circles[7].SelectedBy == 1) ||
			(s.Circles[2].SelectedBy == 1 && s.Circles[5].SelectedBy == 1 && s.Circles[8].SelectedBy == 1) ||

			(s.Circles[0].SelectedBy == 1 && s.Circles[4].SelectedBy == 1 && s.Circles[8].SelectedBy == 1) || // Check all diag1nals.
			(s.Circles[2].SelectedBy == 1 && s.Circles[4].SelectedBy == 1 && s.Circles[6].SelectedBy == 1)

		freeCellsLeft = s.Circles[0].SelectedBy == 2 || s.Circles[1].SelectedBy == 2 || s.Circles[2].SelectedBy == 2 ||
			s.Circles[3].SelectedBy == 2 || s.Circles[4].SelectedBy == 2 || s.Circles[5].SelectedBy == 2 ||
			s.Circles[6].SelectedBy == 2 || s.Circles[7].SelectedBy == 2 || s.Circles[8].SelectedBy == 2
	)

	switch {
	case x && !o:
		s.CapturedBy = 1
	case o && !x:
		s.CapturedBy = 1
	case !freeCellsLeft:
		return nil
	default:
		return nil
	}

	return nil
}

// Check for three in a row for the whole board

func (gb *GameBoard) Update(player, square, circle int) error {
	gb.Squares[square].Circles[circle].SelectedBy = player
	row := gb.Squares[square].CheckCirclesCondition()
	log15.Debug("ROW", "row", row)
	return nil
}
