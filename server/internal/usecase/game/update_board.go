package game

import (
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *Service) UpdateBoard(gameId, playerId, square, circle int64) *entity.Game {
	game := s.repo.Get(gameId)
	game.GameBoard.Squares[square].Circles[circle].SelectedBy = playerId
	updatedSquare := game.GameBoard.CheckCirclesCondition(game.GameBoard.Squares[square])
	game.GameBoard.Squares[square] = updatedSquare
	game.GameBoard.CheckSquareCondition()

	// Check for winner
	if game.GameBoard.Winner != -1 {
		game.Winner = game.GameBoard.Winner
		game.GameOver.Reason = fmt.Sprintf("Player %d wins!", game.Winner)
		game.GameOver.Over = true
	}

	// Change turn
	if int(playerId) == game.Players[0] {
		game.PlayerTurn = game.Players[1]
	} else {
		game.PlayerTurn = game.Players[0]
	}

	return s.repo.Update(game)
}
