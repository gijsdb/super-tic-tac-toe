package controller

import (
	"net/http"
	"strconv"

	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game"
	"github.com/labstack/echo/v4"
)

type GameController struct {
	service game.InteractorI
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
	return c.JSON(http.StatusOK, gc.service.CreateGame(c.QueryParam("player")))
}

func (gc *GameController) Join(c echo.Context) error {
	gId := c.QueryParam("id")
	gameId, err := strconv.ParseInt(gId, 10, 64)
	if err != nil {
		return err
	}

	game, err := gc.service.Join(gameId, c.QueryParam("player"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, game)
}

func (gc *GameController) Leave(c echo.Context) error {
	g := c.Param("id")
	game, err := strconv.ParseInt(g, 10, 64)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gc.service.Leave(game, c.QueryParam("player")))
}

func (gc *GameController) UpdateBoard(c echo.Context) error {
	g := c.Param("id")
	game, err := strconv.ParseInt(g, 10, 64)
	if err != nil {
		return err
	}
	crcle := c.QueryParam("circle")
	circle, err := strconv.ParseInt(crcle, 10, 64)
	if err != nil {
		return err
	}

	s := c.QueryParam("square")
	square, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gc.service.UpdateBoard(game, square, circle, c.QueryParam("player")))
}

func (gc *GameController) RemoveCircle(c echo.Context) error {
	g := c.Param("id")
	game, err := strconv.ParseInt(g, 10, 64)
	if err != nil {
		return err
	}

	s := c.QueryParam("square")
	square, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	crcle := c.Param("c")
	circle, err := strconv.ParseInt(crcle, 10, 64)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gc.service.RemoveCircle(game, square, circle, c.QueryParam("player")))
}

func (gc *GameController) RollDice(c echo.Context) error {
	g := c.Param("id")
	game, err := strconv.ParseInt(g, 10, 64)
	if err != nil {
		return err
	}
	d1 := c.QueryParam("dice1")
	dice1, err := strconv.Atoi(d1)
	if err != nil {
		return err
	}
	d2 := c.QueryParam("dice2")
	dice2, err := strconv.Atoi(d2)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gc.service.RollDice(dice1, dice2, game))
}
