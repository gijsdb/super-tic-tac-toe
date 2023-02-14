package player

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *Service) CreatePlayer(isNewPlayer int64) int64 {

	if isNewPlayer == 0 {
		player := &entity.Player{
			ID:     -1,
			Active: true,
		}

		return s.repo.Create(player)
	} else {
		// returning player, set to active
		s.repo.SetPlayerActive(isNewPlayer)
		return isNewPlayer
	}
}
