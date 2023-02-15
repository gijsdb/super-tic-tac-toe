package player

func (s *Service) SetInactive(playerId int64) {
	player := s.repo.Get(playerId)
	if player == nil {
		return
	}
	player.Active = false
	s.repo.Update(player)
	return
}
