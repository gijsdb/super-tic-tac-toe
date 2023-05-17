package controller

import (
	"fmt"

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

func (sc *SessionController) HandleCreateTempSession(c echo.Context) error {
	player_id := c.QueryParam("player")

	player := sc.player_service.Get(player_id)
	if player == nil {
		return fmt.Errorf("no player found for provided id")
	}

	session_expiry := sc.session_service.GetTempSessionExpiry()
	token := sc.session_service.Create(player_id, session_expiry)
	c.SetCookie(CreateCookie("session_token", token, session_expiry, true))
	return nil
}

// func (sc *SessionController) HandleRefreshSession(c echo.Context) error {
// 	// can be lax with error checking since we know the cookie exists because of the middleware?
// 	cookie, _ := c.Cookie("session_token")

// 	session, _ := sc.session_service.Get(cookie.Value)
// 	session_expiry := sc.session_service.GetTempSessionExpiry()
// 	new_token := sc.session_service.Create(session.PlayerID, session_expiry)

// 	sc.session_service.Delete(cookie.Value)

// 	c.SetCookie(CreateCookie("session_token", new_token, session_expiry, true))
// 	return nil
// }
