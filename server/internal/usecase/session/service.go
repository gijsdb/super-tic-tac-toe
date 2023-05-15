package session

import (
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/labstack/echo/v4"
)

type InteractorI interface {
	IsSessionExpired(token string) bool
	Create(playerId string, expiry time.Time) string
	Get(token string) (*entity.Session, error)
	Delete(token string)
	AuthenticateCookie(next echo.HandlerFunc) echo.HandlerFunc // Middleware
	Refresh(token string)
	GetExpiry() time.Duration
}

func NewService(repo repository.SessionRepositoryI) InteractorI {
	return &SessionService{
		repo:          repo,
		sessionExpiry: 120 * time.Second,
	}
}

type SessionService struct {
	repo          repository.SessionRepositoryI
	sessionExpiry time.Duration
}
