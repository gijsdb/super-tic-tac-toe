package route

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/controller"
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/middleware"

	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"
	"github.com/labstack/echo/v4"
)

func NewGameRouter(e *echo.Echo, service game.InteractorI, session_service session.InteractorI) {
	var gc = controller.NewGameController(service)
	var sm = middleware.NewSessionMiddleware(session_service)

	e.GET("/games", gc.HandleIndex, sm.CheckSession)
	e.GET("/game/create", gc.HandleCreate, sm.CheckSession)
	e.GET("/game/join", gc.HandleJoin, sm.CheckSession)
	e.GET("/game/:id/leave", gc.HandleLeave, sm.CheckSession)
	e.GET("/game/:id", gc.HandleGet, sm.CheckSession)
	e.GET("/game/:id/board/update", gc.HandleUpdateBoard, sm.CheckSession)
	e.GET("/game/:id/board/circle/:c/remove", gc.HandleRemoveCircle, sm.CheckSession)
	e.GET("/game/:id/roll", gc.HandleRollDice, sm.CheckSession)
}
