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

func (pc *PlayerController) HandleCreatePlayer(c echo.Context) error {
	id := c.QueryParam("id")

	pId := pc.playerService.CreatePlayer(id)

	sessionExpiry := time.Now().Add(120 * time.Second)
	session := pc.sessionService.Create(pId, sessionExpiry)

	c.SetCookie(&http.Cookie{Name: "session_token", Value: session, Expires: sessionExpiry, SameSite: http.SameSiteStrictMode, Path: "/"})
	return c.JSON(http.StatusOK, pId)
}

func (pc *PlayerController) HandleSetInactive(c echo.Context) error {
	id := c.Param("id")

	pc.playerService.SetInactive(id)

	return nil
}

func (pc *PlayerController) HandleOauthLogin(c echo.Context) error {
	url := pc.playerService.OauthLogin()

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (pc *PlayerController) HandleGoogleCallback(c echo.Context) error {
	playerId, err := pc.playerService.GoogleCallback(c.FormValue("state"), c.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	sessionExpiry := time.Now().Add(120 * time.Second)
	session := pc.sessionService.Create(playerId, sessionExpiry)

	c.SetCookie(&http.Cookie{Name: "session_token", Value: session, Expires: sessionExpiry, SameSite: http.SameSiteStrictMode, Path: "/"})
	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/")
}

func (pc *PlayerController) HandleRefreshSession(c echo.Context) error {
	cookie, _ := c.Cookie("session_token")

	session, _ := pc.sessionService.Get(cookie.Value)

	sessionExpiry := time.Now().Add(120 * time.Second)
	newToken := pc.sessionService.Create(session.PlayerID, sessionExpiry)

	pc.sessionService.Delete(cookie.Value)
	c.SetCookie(&http.Cookie{Name: "session_token", Value: newToken, Expires: sessionExpiry, SameSite: http.SameSiteStrictMode, Path: "/"})
	return nil
}
