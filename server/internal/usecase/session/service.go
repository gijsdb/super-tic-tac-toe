package session

import (
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

type InteractorI interface {
	IsSessionExpired(token string) bool
	GetTempSessionExpiry() time.Time
	Create(playerId string, expiry time.Time) string
	Get(token string) (*entity.Session, error)
	Delete(token string)
	Refresh(token string) (string, error)
}

func NewService(repo repository.SessionRepositoryI) InteractorI {
	return &SessionService{
		repo:                  repo,
		temp_session_duration: 1 * time.Minute, // TODO change to 5 minutes
	}
}

type SessionService struct {
	repo                  repository.SessionRepositoryI
	temp_session_duration time.Duration
}
