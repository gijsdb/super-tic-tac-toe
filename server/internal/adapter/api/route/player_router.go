package route

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/controller"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/labstack/echo/v4"
)

func NewPlayerRouter(e *echo.Echo, service player.InteractorI) {
	var player_controller = controller.NewPlayerController(service)
	e.GET("/player/create", player_controller.HandleCreatePlayer)
	e.GET("/player/:id/setinactive", player_controller.HandleSetInactive)
	e.GET("/login", player_controller.HandleOauthLogin)
	e.GET("/callback", player_controller.HandleGoogleCallback)
}
