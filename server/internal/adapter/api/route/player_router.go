package route

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/controller"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"
	"github.com/labstack/echo/v4"
)

func NewPlayerRouter(e *echo.Echo, pService player.InteractorI, sService session.InteractorI) {
	var pc = controller.NewPlayerController(pService, sService)
	e.GET("/player/create", pc.HandleCreatePlayer)
	e.GET("/login", pc.HandleOauthLogin)
	e.GET("/callback", pc.HandleGoogleCallback)
}
