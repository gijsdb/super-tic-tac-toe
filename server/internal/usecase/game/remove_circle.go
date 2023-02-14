package game

// TODO: change front-end store to receive game or error instead of string
func (s *Service) RemoveCircle(gameId, playerId, square, circle int64) string {
	game := s.repo.Get(gameId)

	if game.GameBoard.Squares[square].Circles[circle].SelectedBy == playerId {
		return "you must select a circle of the other player, not your own"
	}

	if game.GameBoard.Squares[square].Circles[circle].SelectedBy == -1 {
		return "you must select a circle of the other player, this one has not been captured"
	}

	game.GameBoard.Squares[square].Circles[circle].SelectedBy = -1

	s.repo.Update(game)

	return "success"
}
