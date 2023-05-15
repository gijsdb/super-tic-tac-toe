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
	AuthenticateCookie(next echo.HandlerFunc) echo.HandlerFunc
	Refresh(token string)
}

func NewService(sessionRepo repository.SessionRepositoryI) InteractorI {
	return &Service{
		sessionRepo: sessionRepo,
	}
}

type Service struct {
	sessionRepo repository.SessionRepositoryI
}
