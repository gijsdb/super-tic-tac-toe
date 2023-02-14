package controller

import (
	"net/http"
	"strconv"

	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/labstack/echo/v4"
)

type PlayerController struct {
	service player.InteractorI
	// TODO: Implement Logger
}

func NewPlayerController(service player.InteractorI) PlayerController {
	return PlayerController{
		service: service,
	}
}

func (gc *PlayerController) CreatePlayer(c echo.Context) error {
	id := c.QueryParam("id")

	isNewPlayer, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	res := gc.service.CreatePlayer(isNewPlayer)

	return c.JSON(http.StatusOK, res)
}
