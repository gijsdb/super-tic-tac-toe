package player

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

// Need a better system for this. Currently a player can get over ridden by another quite easily if the server restarts
func (s *Service) CreatePlayer(playerId string) string {

	if playerId == "" {
		player := &entity.Player{
			ID:     "",
			Active: true,
		}

		return s.repo.Create(player)
	} else {

		// returning player, set to active
		player := s.repo.Get(playerId)
		if player == nil {
			return s.repo.Create(&entity.Player{
				ID:     playerId,
				Active: true,
			})
		}
		player.Active = true

		return s.repo.Update(player).ID
	}
}
