package route

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/controller"
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/middleware"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"
	"github.com/labstack/echo/v4"
)

func NewPlayerRouter(e *echo.Echo, player_service player.InteractorI, session_service session.InteractorI) {
	var pc = controller.NewPlayerController(player_service, session_service)
	var sm = middleware.NewSessionMiddleware(session_service, player_service)

	e.GET("/player/create", pc.HandleCreateTempPlayer)
	e.GET("/player/:id", pc.HandleGetPlayer, sm.CheckSession, sm.RenewSession)
	e.GET("/login", pc.HandleOauthLogin)
	e.GET("/callback", pc.HandleGoogleCallback)
}
