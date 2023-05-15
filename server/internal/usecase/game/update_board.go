package game

import (
	"fmt"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *GameService) UpdateBoard(gameId, square, circle int64, playerId string) *entity.Game {
	game := s.repo.Get(gameId)

	game.GameBoard.Squares[square].Circles[circle].SelectedBy = playerId
	updatedSquare := game.GameBoard.CheckCirclesCondition(game.GameBoard.Squares[square])
	game.GameBoard.Squares[square] = updatedSquare
	game.GameBoard.CheckSquareCondition()

	// Check for winner
	if game.GameBoard.Winner != "" {
		game.Winner = game.GameBoard.Winner
		game.GameOver.Reason = fmt.Sprintf("Player %s wins!", game.Winner)
		game.GameOver.Over = true
		game.GameOver.EndTime = time.Now().Unix()
	}

	// Change turn
	if playerId == game.Players[0] {
		game.PlayerTurn = game.Players[1]
	} else {
		game.PlayerTurn = game.Players[0]
	}

	return s.repo.Update(game)
}
