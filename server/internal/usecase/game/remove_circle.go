package game

import "github.com/gijsdb/super-tic-tac-toe/internal/entity"

func (s *Service) RemoveCircle(gameId, playerId, square, circle int64) *entity.Game {
	game := s.repo.Get(gameId)

	game.GameBoard.Squares[square].Circles[circle].SelectedBy = -1

	if int(playerId) == game.Players[0] {
		game.PlayerTurn = game.Players[1]
	} else {
		game.PlayerTurn = game.Players[0]
	}

	return s.repo.Update(game)
}
