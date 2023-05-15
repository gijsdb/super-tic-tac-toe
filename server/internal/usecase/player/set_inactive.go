package player

func (s *Service) SetInactive(playerId string) {
	player := s.playerRepo.Get(playerId)
	if player == nil {
		return
	}
	player.Active = false
	s.playerRepo.Update(player)
}
