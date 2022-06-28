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
	index      int
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
			index:      i,
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

// Check for three in a row for circles in a square
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
		s.CapturedBy = 0
	case o && !x:
		s.CapturedBy = 2
	case !freeCellsLeft:
		return nil
	default:
		return nil
	}

	return nil
}

// Check for three in a row for squares
func (gb *GameBoard) CheckSquareCondition() string {

	var (
		x = (gb.Squares[0].CapturedBy == 0 && gb.Squares[1].CapturedBy == 0 && gb.Squares[2].CapturedBy == 0) || // Check all rows.
			(gb.Squares[3].CapturedBy == 0 && gb.Squares[4].CapturedBy == 0 && gb.Squares[5].CapturedBy == 0) ||
			(gb.Squares[6].CapturedBy == 0 && gb.Squares[7].CapturedBy == 0 && gb.Squares[8].CapturedBy == 0) ||

			(gb.Squares[0].CapturedBy == 0 && gb.Squares[3].CapturedBy == 0 && gb.Squares[6].CapturedBy == 0) || // Check all columns.
			(gb.Squares[1].CapturedBy == 0 && gb.Squares[4].CapturedBy == 0 && gb.Squares[7].CapturedBy == 0) ||
			(gb.Squares[2].CapturedBy == 0 && gb.Squares[5].CapturedBy == 0 && gb.Squares[8].CapturedBy == 0) ||

			(gb.Squares[0].CapturedBy == 0 && gb.Squares[4].CapturedBy == 0 && gb.Squares[8].CapturedBy == 0) || // Check all diagonals.
			(gb.Squares[2].CapturedBy == 0 && gb.Squares[4].CapturedBy == 0 && gb.Squares[6].CapturedBy == 0)

		o = (gb.Squares[0].CapturedBy == 1 && gb.Squares[1].CapturedBy == 1 && gb.Squares[2].CapturedBy == 1) || // Check all r1ws.
			(gb.Squares[3].CapturedBy == 1 && gb.Squares[4].CapturedBy == 1 && gb.Squares[5].CapturedBy == 1) ||
			(gb.Squares[6].CapturedBy == 1 && gb.Squares[7].CapturedBy == 1 && gb.Squares[8].CapturedBy == 1) ||

			(gb.Squares[0].CapturedBy == 1 && gb.Squares[3].CapturedBy == 1 && gb.Squares[6].CapturedBy == 1) || // Check all c1lumns.
			(gb.Squares[1].CapturedBy == 1 && gb.Squares[4].CapturedBy == 1 && gb.Squares[7].CapturedBy == 1) ||
			(gb.Squares[2].CapturedBy == 1 && gb.Squares[5].CapturedBy == 1 && gb.Squares[8].CapturedBy == 1) ||

			(gb.Squares[0].CapturedBy == 1 && gb.Squares[4].CapturedBy == 1 && gb.Squares[8].CapturedBy == 1) || // Check all diag1nals.
			(gb.Squares[2].CapturedBy == 1 && gb.Squares[4].CapturedBy == 1 && gb.Squares[6].CapturedBy == 1)

		freeCellsLeft = gb.Squares[0].CapturedBy == 2 || gb.Squares[1].CapturedBy == 2 || gb.Squares[2].CapturedBy == 2 ||
			gb.Squares[3].CapturedBy == 2 || gb.Squares[4].CapturedBy == 2 || gb.Squares[5].CapturedBy == 2 ||
			gb.Squares[6].CapturedBy == 2 || gb.Squares[7].CapturedBy == 2 || gb.Squares[8].CapturedBy == 2
	)

	switch {
	case x && !o:
		return "player1"
	case o && !x:
		return "player2"
	case !freeCellsLeft:
		return "CONTINUE"
	default:
		return "default"
	}

}

func (gb *GameBoard) Update(player, square, circle int) error {
	gb.Squares[square].Circles[circle].SelectedBy = player
	circleWin := gb.Squares[square].CheckCirclesCondition()
	squareWin := gb.CheckSquareCondition()
	log15.Debug("circle", "circle", circleWin)
	log15.Debug("square", "square", squareWin)
	return nil
}
