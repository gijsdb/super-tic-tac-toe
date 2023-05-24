package game

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *GameService) CreateGame(creatingPlayer string) *entity.Game {
	gb := newGameboard(creatingPlayer)
	game := entity.Game{
		ID:         -1,
		GameBoard:  &gb,
		PlayerTurn: creatingPlayer,
		GameOver: &entity.GameOver{
			Over:   false,
			Reason: "",
		},
		Winner:   "",
		Players:  []string{creatingPlayer},
		Full:     false,
		LastRoll: []int{0, 0},
	}
	return s.game_repo.Create(&game)
}

func newGameboard(creatingPlayer string) entity.GameBoard {
	var board entity.GameBoard
	board.Player1 = creatingPlayer
	board.Winner = ""
	var squares []entity.Square
	for i := 0; i < 9; i++ {
		square := entity.Square{
			Circles:    initCircles(),
			CapturedBy: "",
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
		circle.SelectedBy = ""

		circles = append(circles, circle)
	}

	return circles
}
