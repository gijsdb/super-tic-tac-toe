package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"
	"github.com/labstack/echo/v4"
)

type SessionController struct {
	session_service session.InteractorI
	player_service  player.InteractorI
}

func NewSessionController(session_service session.InteractorI, player_service player.InteractorI) SessionController {
	return SessionController{
		session_service: session_service,
		player_service:  player_service,
	}
}

func (sc *SessionController) HandleCreateSession(c echo.Context) error {
	player_id := c.QueryParam("player")

	// Todo check the player actually exists
	player := sc.player_service.Get(player_id)
	if player == nil {
		return fmt.Errorf("no player found for provided id")
	}

	sessionExpiry := sc.getSessionExpiry()
	token := sc.session_service.Create(player_id, sessionExpiry)

	c.SetCookie(&http.Cookie{Name: "session_token", Value: token, Expires: sessionExpiry, SameSite: http.SameSiteStrictMode, Path: "/"})
	return nil
}

func (sc *SessionController) HandleRefreshSession(c echo.Context) error {
	// can be lax with error checking since we know the cookie exists because of the middleware?
	cookie, _ := c.Cookie("session_token")

	session, _ := sc.session_service.Get(cookie.Value)

	sessionExpiry := sc.getSessionExpiry()
	newToken := sc.session_service.Create(session.PlayerID, sessionExpiry)

	sc.session_service.Delete(cookie.Value)
	c.SetCookie(&http.Cookie{Name: "session_token", Value: newToken, Expires: sessionExpiry, SameSite: http.SameSiteStrictMode, Path: "/"})
	return nil
}

// duplicated for now
func (sc *SessionController) getSessionExpiry() time.Time {
	return time.Now().Add(sc.session_service.GetExpiry())
}
