package route

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/controller"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"
	"github.com/labstack/echo/v4"
)

func NewPlayerRouter(e *echo.Echo, pService player.InteractorI, sService session.InteractorI) {
	var player_controller = controller.NewPlayerController(pService, sService)
	e.GET("/player/create", player_controller.HandleCreatePlayer)
	e.GET("/player/:id/setinactive", player_controller.HandleSetInactive, sService.AuthenticateCookie)
	e.GET("/login", player_controller.HandleOauthLogin)
	e.GET("/callback", player_controller.HandleGoogleCallback)
	e.GET("/session/refresh", player_controller.HandleRefreshSession, sService.AuthenticateCookie)
}
