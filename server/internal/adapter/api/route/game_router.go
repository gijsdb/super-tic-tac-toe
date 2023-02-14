package route

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/controller"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game"
	"github.com/labstack/echo/v4"
)

func NewGameRouter(e *echo.Echo, service game.InteractorI) {
	var gc = controller.NewGameController(service)
	e.GET("/games", gc.Index)
	e.GET("/game/create", gc.Create)
	e.GET("/game/join", gc.Join)
	e.GET("/game/:id", gc.Get)
}
