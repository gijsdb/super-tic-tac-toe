package player

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *Service) CreatePlayer(playerId string) string {

	if playerId == "" {
		player := &entity.Player{
			ID:     "",
			Active: true,
		}

		return s.playerRepo.Create(player)
	} else {

		// returning player, set to active
		player := s.playerRepo.Get(playerId)
		if player == nil {
			return s.playerRepo.Create(&entity.Player{
				ID:     playerId,
				Active: true,
			})
		}
		player.Active = true

		return s.playerRepo.Update(player).ID
	}
}
