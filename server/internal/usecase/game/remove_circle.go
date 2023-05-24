package game

import "github.com/gijsdb/super-tic-tac-toe/internal/entity"

func (s *GameService) RemoveCircle(gameId, square, circle int64, playerId string) *entity.Game {
	game := s.game_repo.Get(gameId)

	game.GameBoard.Squares[square].Circles[circle].SelectedBy = ""

	if playerId == game.Players[0] {
		game.PlayerTurn = game.Players[1]
	} else {
		game.PlayerTurn = game.Players[0]
	}

	return s.game_repo.Update(game)
}
