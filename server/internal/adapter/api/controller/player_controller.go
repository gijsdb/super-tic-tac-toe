package controller

import (
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
	gc.service.CreatePlayer()

	return nil
}
