package route

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/controller"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game_instance"
	"github.com/labstack/echo/v4"
)

func NewGameInstanceRouter(e *echo.Echo, service game_instance.InteractorI) {
	var game_instance_controller = controller.NewGameInstanceController(service)
	e.GET("/instances", game_instance_controller.Get)
}
