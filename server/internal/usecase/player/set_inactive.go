package player

func (s *Service) SetInactive(playerId string) {
	player := s.repo.Get(playerId)
	if player == nil {
		return
	}
	player.Active = false
	s.repo.Update(player)
	return
}
