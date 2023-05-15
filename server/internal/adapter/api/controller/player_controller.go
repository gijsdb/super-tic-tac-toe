package controller

import (
	"net/http"

	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/labstack/echo/v4"
)

type PlayerController struct {
	service player.InteractorI
}

func NewPlayerController(service player.InteractorI) PlayerController {
	return PlayerController{
		service: service,
	}
}

func (gc *PlayerController) HandleCreatePlayer(c echo.Context) error {
	id := c.QueryParam("id")

	res := gc.service.CreatePlayer(id)

	return c.JSON(http.StatusOK, res)
}

func (gc *PlayerController) HandleSetInactive(c echo.Context) error {
	id := c.Param("id")

	gc.service.SetInactive(id)

	return nil
}

func (gc *PlayerController) HandleOauthLogin(c echo.Context) error {
	url := gc.service.OauthLogin()

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (gc *PlayerController) HandleGoogleCallback(c echo.Context) error {
	gc.service.GoogleCallback(c.FormValue("state"), c.FormValue("code"))

	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/")
}
