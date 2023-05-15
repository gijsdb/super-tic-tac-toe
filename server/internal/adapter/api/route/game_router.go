package route

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/controller"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"
	"github.com/labstack/echo/v4"
)

func NewGameRouter(e *echo.Echo, service game.InteractorI, sService session.InteractorI) {
	var gc = controller.NewGameController(service)
	e.GET("/games", gc.HandleIndex, sService.AuthenticateCookie)
	e.GET("/game/create", gc.HandleCreate, sService.AuthenticateCookie)
	e.GET("/game/join", gc.HandleJoin, sService.AuthenticateCookie)
	e.GET("/game/:id/leave", gc.HandleLeave, sService.AuthenticateCookie)
	e.GET("/game/:id", gc.HandleGet, sService.AuthenticateCookie)
	e.GET("/game/:id/board/update", gc.HandleUpdateBoard, sService.AuthenticateCookie)
	e.GET("/game/:id/board/circle/:c/remove", gc.HandleRemoveCircle, sService.AuthenticateCookie)
	e.GET("/game/:id/roll", gc.HandleRollDice, sService.AuthenticateCookie)
}
