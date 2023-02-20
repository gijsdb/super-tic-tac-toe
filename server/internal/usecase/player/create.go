package player

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/inconshreveable/log15"
)

// Need a better system for this. Currently a player can get over ridden by another quite easily if the server restarts
func (s *Service) CreatePlayer(playerId int64) int64 {

	if playerId == 0 {
		player := &entity.Player{
			ID:     -1,
			Active: true,
		}

		return s.repo.Create(player)
	} else {

		// returning player, set to active
		player := s.repo.Get(playerId)
		if player == nil {
			log15.Debug("Player returning after restart")
			return s.repo.Create(&entity.Player{
				ID:     playerId,
				Active: true,
			})
		}
		player.Active = true

		return s.repo.Update(player).ID
	}
}
