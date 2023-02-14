package controller

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game_instance"
	"github.com/labstack/echo/v4"
)

type GameInstanceController struct {
	service game_instance.InteractorI
	// TODO: Implement Logger
}

func NewGameInstanceController(service game_instance.InteractorI) GameInstanceController {
	return GameInstanceController{
		service: service,
	}
}

func (gc *GameInstanceController) Get(c echo.Context) error {
	return nil
}
