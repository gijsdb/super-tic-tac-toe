package middleware

import (
	"net/http"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"
	"github.com/labstack/echo/v4"
)

type SessionMiddlewareI interface {
	CheckSession(next echo.HandlerFunc) echo.HandlerFunc
	RenewSession(next echo.HandlerFunc) echo.HandlerFunc
}

type SessionMiddleware struct {
	session_service session.InteractorI
	player_service  player.InteractorI
}

func NewSessionMiddleware(session_service session.InteractorI, player_service player.InteractorI) SessionMiddlewareI {
	return &SessionMiddleware{
		session_service: session_service,
		player_service:  player_service,
	}
}

func (sm *SessionMiddleware) RenewSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				return c.JSON(http.StatusUnauthorized, "unauthorized: no session cookie found")
			}
			// For any other type of error, return a bad request status
			return c.JSON(http.StatusBadRequest, "bad request")
		}

		// renew session
		new_token, err := sm.session_service.Refresh(cookie.Value)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "unauthorized: session does not exist during refresh")
		}

		playerId, err := sm.session_service.GetPlayerIdFromSession(cookie.Value)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "unauthorized: session does not exist")
		}

		player := sm.player_service.Get(playerId)
		var session_expiry time.Time
		if player.Email == "" {
			session_expiry = sm.session_service.GetTempSessionExpiry()
		} else {
			// TODO session_expiry = sm.session_service.GetSessionExpiry()
		}

		c.SetCookie(&http.Cookie{Name: "session_token", Value: new_token, Expires: session_expiry, SameSite: http.SameSiteStrictMode, Path: "/", Secure: true, HttpOnly: true})

		return next(c)
	}

}

func (sm *SessionMiddleware) CheckSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				return c.JSON(http.StatusUnauthorized, "unauthorized: no session cookie found")
			}
			// For any other type of error, return a bad request status
			return c.JSON(http.StatusBadRequest, "bad request")
		}

		_, err = sm.session_service.Get(cookie.Value)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "unauthorized: session does not exist")
		}

		expired := sm.session_service.IsSessionExpired(cookie.Value)
		if expired {
			sm.session_service.Delete(cookie.Value)
			return c.JSON(http.StatusUnauthorized, "unauthorized: session expired")
		}

		return next(c)
	}
}
