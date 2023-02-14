package controller

import (
	"net/http"
	"strconv"

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

func (gc *GameController) Index(c echo.Context) error {

	return nil
}

func (gc *GameController) Create(c echo.Context) error {
	id := c.QueryParam("player")
	playerId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	game := gc.service.CreateGame(playerId)

	return c.JSON(http.StatusOK, game)
}
