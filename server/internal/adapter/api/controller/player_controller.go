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
	player_service  player.InteractorI
	session_service session.InteractorI
}

func NewPlayerController(playerService player.InteractorI, sessionService session.InteractorI) PlayerController {
	return PlayerController{
		player_service:  playerService,
		session_service: sessionService,
	}
}

// make unauthenticated user use different cookie and store relevant data in the cookie so the user can continue with that cookie if they return
func (pc *PlayerController) HandleCreateTempPlayer(c echo.Context) error {
	player_id := pc.player_service.Create()

	session_expiry := pc.session_service.GetTempSessionExpiry()
	session := pc.session_service.Create(player_id, session_expiry)

	c.SetCookie(CreateCookie("session_token", session, session_expiry, true))
	c.SetCookie(CreateCookie("temp_account", player_id, time.Now().Add(730*time.Hour), false))

	return c.JSON(http.StatusOK, player_id)
}

// --------------- OAUTH -----------//

func (pc *PlayerController) HandleOauthLogin(c echo.Context) error {
	url := pc.player_service.OauthLogin()

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (pc *PlayerController) HandleGoogleCallback(c echo.Context) error {
	player_id, err := pc.player_service.GoogleCallback(c.FormValue("state"), c.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	session_expiry := pc.session_service.GetTempSessionExpiry()
	session := pc.session_service.Create(player_id, session_expiry)

	c.SetCookie(CreateCookie("session_token", session, session_expiry, true))
	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/")
}
