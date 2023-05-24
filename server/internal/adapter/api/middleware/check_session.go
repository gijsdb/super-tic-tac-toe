package middleware

import (
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"

	"github.com/labstack/echo/v4"
)

type SessionMiddlewareI interface {
	CheckSession(next echo.HandlerFunc) echo.HandlerFunc
	RenewSession(next echo.HandlerFunc) echo.HandlerFunc
}

type SessionMiddleware struct {
	session_service session.ServiceI
	player_service  player.ServiceI
}

func NewSessionMiddleware(s session.ServiceI, p player.ServiceI) SessionMiddlewareI {
	return &SessionMiddleware{
		session_service: s,
		player_service:  p,
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

		new_session, new_token, err := sm.session_service.Refresh(cookie.Value)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "unauthorized: session does not exist during refresh")
		}

		c.SetCookie(&http.Cookie{Name: "session_token", Value: new_token, Expires: new_session.Expiry, SameSite: http.SameSiteStrictMode, Path: "/", Secure: true, HttpOnly: true})

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

		session, err := sm.session_service.Get(cookie.Value)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "unauthorized: session does not exist")
		}

		expired := sm.session_service.IsSessionExpired(cookie.Value)
		if expired {
			// Check if the player is authenticated with google, if so set to false
			// return a redirect to login
			player, err := sm.player_service.Get(session.PlayerID)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "unauthorized: player not found for session")
			}

			if !player.IsTemp {
				sm.session_service.Delete(cookie.Value)
				return c.Redirect(302, "http://localhost:1323/login")
			}

			sm.session_service.Delete(cookie.Value)
			return c.JSON(http.StatusUnauthorized, "unauthorized: session expired")
		}

		return next(c)
	}
}
