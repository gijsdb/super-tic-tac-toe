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

func (gc *PlayerController) CreatePlayer(c echo.Context) error {
	id := c.QueryParam("id")

	res := gc.service.CreatePlayer(id)

	return c.JSON(http.StatusOK, res)
}

func (gc *PlayerController) SetInactive(c echo.Context) error {
	id := c.Param("id")

	gc.service.SetInactive(id)

	return nil
}
