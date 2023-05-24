package controller

import (
	"net/http"
	"strconv"

	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game"

	"github.com/labstack/echo/v4"
)

type GameController struct {
	service game.ServiceI
}

func NewGameController(service game.ServiceI) GameController {
	return GameController{
		service: service,
	}
}

func (gc *GameController) HandleIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, gc.service.Index())
}

func (gc *GameController) HandleGet(c echo.Context) error {
	id := c.Param("id")
	game_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gc.service.Get(game_id))
}

func (gc *GameController) HandleCreate(c echo.Context) error {
	return c.JSON(http.StatusOK, gc.service.CreateGame(c.QueryParam("player")))
}

func (gc *GameController) HandleJoin(c echo.Context) error {
	g := c.QueryParam("id")
	game_id, err := strconv.ParseInt(g, 10, 64)
	if err != nil {
		return err
	}

	game, err := gc.service.Join(game_id, c.QueryParam("player"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, game)
}

func (gc *GameController) HandleLeave(c echo.Context) error {
	g := c.Param("id")
	game, err := strconv.ParseInt(g, 10, 64)
	if err != nil {
		return err
	}

	err = gc.service.Leave(game, c.QueryParam("player"))
	if err != nil {
		return err
	}

	return nil
}

func (gc *GameController) HandleUpdateBoard(c echo.Context) error {
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

func (gc *GameController) HandleRemoveCircle(c echo.Context) error {
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

func (gc *GameController) HandleRollDice(c echo.Context) error {
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
