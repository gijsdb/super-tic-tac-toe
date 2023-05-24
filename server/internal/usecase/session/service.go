package session

import (
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

type ServiceI interface {
	IsSessionExpired(token string) bool
	GetTempSessionExpiry() time.Time
	Create(playerId string, expiry time.Time) string
	Get(token string) (*entity.Session, error)
	Delete(token string)
	Refresh(token string) (*entity.Session, string, error)
}

type SessionService struct {
	memorystore           repository.SessionMemoryRepoI
	temp_session_duration time.Duration
}

func NewSessionService(mem_store repository.SessionMemoryRepoI) ServiceI {
	return &SessionService{
		memorystore:           mem_store,
		temp_session_duration: 5 * time.Minute, // TODO make configurable?
	}
}
