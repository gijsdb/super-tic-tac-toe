package controller

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game"
	"github.com/labstack/echo/v4"
)

type GameController struct {
	service game.InteractorI
	// TODO: Implement Logger
}

func NewGameController(service game.InteractorI) GameController {
	return GameController{
		service: service,
	}
}

func (gc *GameController) GetAllGames(c echo.Context) error {
	return nil
}
