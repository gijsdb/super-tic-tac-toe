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
		return nil, fmt.Errorf("session does not exist")
	}

	return session, nil
}

func (s *SessionService) Delete(token string) {
	s.repo.Delete(token)
}

func (s *SessionService) Refresh(token string) {

}

func (s *SessionService) GetExpiry() time.Duration {
	return s.sessionExpiry
}
