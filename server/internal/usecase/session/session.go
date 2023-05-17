package session

import (
	"fmt"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *SessionService) IsSessionExpired(token string) bool {
	session, err := s.repo.Get(token)
	if err != nil {
		return true
	}
	return session.Expiry.Before(time.Now())
}

func (s *SessionService) Create(playerId string, expiry time.Time) string {
	return s.repo.Create(&entity.Session{
		PlayerID: playerId,
		Expiry:   expiry,
	})
}

func (s *SessionService) Get(token string) (*entity.Session, error) {
	session, err := s.repo.Get(token)
	if err != nil {
		return nil, fmt.Errorf("session does not exist - Get()::session.go")
	}

	return session, nil
}

func (s *SessionService) Delete(token string) {
	s.repo.Delete(token)
}

func (s *SessionService) Refresh(token string) (string, error) {
	current_session, err := s.repo.Get(token)
	if err != nil {
		return "", fmt.Errorf("session does not exist - Refresh()::session.go")
	}

	current_session.Expiry = s.GetTempSessionExpiry()

	s.repo.Delete(token)

	new_token := s.repo.Create(current_session)
	return new_token, nil
}

func (s *SessionService) GetTempSessionExpiry() time.Time {
	return time.Now().Add(s.temp_session_duration)
}
