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
	return c.JSON(http.StatusOK, gc.service.Index())
}

func (gc *GameController) Get(c echo.Context) error {
	id := c.Param("id")
	gameId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gc.service.Get(gameId))
}

func (gc *GameController) Create(c echo.Context) error {
	id := c.QueryParam("player")
	playerId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gc.service.CreateGame(playerId))
}

func (gc *GameController) Join(c echo.Context) error {
	gId := c.QueryParam("id")
	gameId, err := strconv.ParseInt(gId, 10, 64)
	if err != nil {
		return err
	}
	pId := c.QueryParam("player")
	playerId, err := strconv.ParseInt(pId, 10, 64)
	if err != nil {
		return err
	}

	game, err := gc.service.Join(gameId, playerId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, game)
}
