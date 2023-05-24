package game

import (
	"fmt"
	"time"
)

func (s *GameService) Leave(gameId int64, leavingPlayer string) error {
	game := s.game_repo.Get(gameId)
	if game == nil {
		return fmt.Errorf("game being left does not exist")
	}

	for i, player := range game.Players {
		if player == leavingPlayer {
			game.Players = append(game.Players[:i], game.Players[i+1:]...)
			// Increment losses if !IsTemp
		}
	}

	game.GameOver.Over = true
	game.GameOver.Reason = "Player left the game"
	game.GameOver.EndTime = time.Now().Unix()
	s.game_repo.Update(game)

	return nil
}
