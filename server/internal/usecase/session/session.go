package session

import (
	"fmt"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *SessionService) IsSessionExpired(token string) bool {
	session, err := s.memorystore.Get(token)
	if err != nil {
		return true
	}
	return session.Expiry.Before(time.Now())
}

func (s *SessionService) Create(playerId string, expiry time.Time) string {
	return s.memorystore.Create(&entity.Session{
		PlayerID: playerId,
		Expiry:   expiry,
	})
}

func (s *SessionService) Get(token string) (*entity.Session, error) {
	session, err := s.memorystore.Get(token)
	if err != nil {
		return nil, fmt.Errorf("session does not exist - Get()::session.go")
	}

	return session, nil
}

func (s *SessionService) Delete(token string) {
	s.memorystore.Delete(token)
}

func (s *SessionService) Refresh(token string) (*entity.Session, string, error) {
	current_session, err := s.memorystore.Get(token)
	if err != nil {
		return nil, "", fmt.Errorf("session does not exist - Refresh()::session.go")
	}

	// TODO check if player is authenticated or temporary and set expiry accordingly
	current_session.Expiry = s.GetTempSessionExpiry()

	s.memorystore.Delete(token)

	new_token := s.memorystore.Create(current_session)
	return &entity.Session{PlayerID: current_session.PlayerID, Expiry: current_session.Expiry, CSRF: current_session.CSRF}, new_token, nil
}

func (s *SessionService) GetTempSessionExpiry() time.Time {
	return time.Now().Add(s.temp_session_duration)
}
