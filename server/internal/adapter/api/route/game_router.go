package route

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/controller"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game"
	"github.com/labstack/echo/v4"
)

func NewGameRouter(e *echo.Echo, service game.InteractorI) {
	var gc = controller.NewGameController(service)
	e.GET("/games", gc.HandleIndex)
	e.GET("/game/create", gc.HandleCreate)
	e.GET("/game/join", gc.HandleJoin)
	e.GET("/game/:id/leave", gc.HandleLeave)
	e.GET("/game/:id", gc.HandleGet)
	e.GET("/game/:id/board/update", gc.HandleUpdateBoard)
	e.GET("/game/:id/board/circle/:c/remove", gc.HandleRemoveCircle)
	e.GET("/game/:id/roll", gc.HandleRollDice)
}
