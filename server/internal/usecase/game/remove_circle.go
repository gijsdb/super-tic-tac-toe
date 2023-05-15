package game

import "github.com/gijsdb/super-tic-tac-toe/internal/entity"

func (s *GameService) RemoveCircle(gameId, square, circle int64, playerId string) *entity.Game {
	game := s.repo.Get(gameId)

	game.GameBoard.Squares[square].Circles[circle].SelectedBy = ""

	if playerId == game.Players[0] {
		game.PlayerTurn = game.Players[1]
	} else {
		game.PlayerTurn = game.Players[0]
	}

	return s.repo.Update(game)
}
