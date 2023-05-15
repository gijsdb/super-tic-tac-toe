package session

import (
	"fmt"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *Service) IsSessionExpired(token string) bool {
	session, err := s.sessionRepo.Get(token)
	if err != nil {
		return true
	}
	return session.Expiry.Before(time.Now())
}

func (s *Service) Create(playerId string, expiry time.Time) string {
	return s.sessionRepo.Create(&entity.Session{
		PlayerID: playerId,
		Expiry:   expiry,
	})
}

func (s *Service) Get(token string) (*entity.Session, error) {
	session, err := s.sessionRepo.Get(token)
	if err != nil {
		return nil, fmt.Errorf("session does not exist")
	}

	return session, nil
}

func (s *Service) Delete(token string) {
	s.sessionRepo.Delete(token)
}

func (s *Service) Refresh(token string) {

}
