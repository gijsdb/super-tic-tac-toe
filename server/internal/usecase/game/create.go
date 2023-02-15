package game

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *Service) CreateGame(creatingPlayer int64) *entity.Game {
	gb := newGameboard(creatingPlayer)
	game := entity.Game{
		ID:         -1,
		GameBoard:  &gb,
		PlayerTurn: int(creatingPlayer),
		GameOver: &entity.GameOver{
			Over:   false,
			Reason: "",
		},
		Winner:   -1,
		Players:  []int{int(creatingPlayer)},
		Full:     false,
		LastRoll: []int{0, 0},
	}
	return s.repo.Create(&game)
}

func newGameboard(creatingPlayer int64) entity.GameBoard {
	var board entity.GameBoard
	board.Player1 = creatingPlayer
	board.Winner = -1
	var squares []entity.Square
	for i := 0; i < 9; i++ {
		var square entity.Square
		square = entity.Square{
			Circles:    initCircles(),
			CapturedBy: -1,
			Index:      i,
		}
		squares = append(squares, square)
	}
	board.Squares = squares
	return board
}

func initCircles() []entity.Circle {
	var circles []entity.Circle

	for i := 0; i < 9; i++ {
		var circle entity.Circle
		circle.Index = i
		circle.SelectedBy = -1

		circles = append(circles, circle)
	}

	return circles
}
