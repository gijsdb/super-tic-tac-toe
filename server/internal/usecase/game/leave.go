package game

import (
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *Service) Leave(gameId, leavingPlayer int64) *entity.Game {
	game := s.repo.Get(gameId)

	for i, player := range game.Players {
		if player == int(leavingPlayer) {
			game.Players = append(game.Players[:i], game.Players[i+1:]...)
		}
	}

	game.GameOver.Over = true
	game.GameOver.Reason = "Player left the game"
	game.GameOver.EndTime = time.Now().Unix()

	return s.repo.Update(game)
}
