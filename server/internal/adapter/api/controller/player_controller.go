package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"
	"github.com/labstack/echo/v4"
)

type PlayerController struct {
	playerService  player.InteractorI
	sessionService session.InteractorI
}

func NewPlayerController(playerService player.InteractorI, sessionService session.InteractorI) PlayerController {
	return PlayerController{
		playerService:  playerService,
		sessionService: sessionService,
	}
}

// make unauthenticated user use different cookie and store relevant data in the cookie so the user can continue with that cookie if they return
func (pc *PlayerController) HandleCreatePlayer(c echo.Context) error {
	player_id := pc.playerService.Create()

	session_expiry := pc.getSessionExpiry()
	session := pc.sessionService.Create(player_id, session_expiry)
	c.SetCookie(&http.Cookie{Name: "session_token", Value: session, Expires: session_expiry, SameSite: http.SameSiteStrictMode, Path: "/"})
	c.SetCookie(&http.Cookie{Name: "temp_account", Value: player_id, Expires: time.Now().Add(730 * time.Hour), SameSite: http.SameSiteStrictMode, Path: "/"})
	return c.JSON(http.StatusOK, player_id)
}

func (pc *PlayerController) HandleOauthLogin(c echo.Context) error {
	url := pc.playerService.OauthLogin()

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (pc *PlayerController) HandleGoogleCallback(c echo.Context) error {
	player_id, err := pc.playerService.GoogleCallback(c.FormValue("state"), c.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	session_expiry := pc.getSessionExpiry()
	session := pc.sessionService.Create(player_id, session_expiry)

	c.SetCookie(&http.Cookie{Name: "session_token", Value: session, Expires: session_expiry, SameSite: http.SameSiteStrictMode, Path: "/"})
	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/")
}

func (pc *PlayerController) getSessionExpiry() time.Time {
	return time.Now().Add(pc.sessionService.GetExpiry())
}
