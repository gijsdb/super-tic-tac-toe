package route

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/controller"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"
	"github.com/labstack/echo/v4"
)

func NewSessionRouter(e *echo.Echo, session_service session.InteractorI, player_service player.InteractorI) {
	var sc = controller.NewSessionController(session_service, player_service)
	e.GET("/session/create", sc.HandleCreateTempSession)
	// e.GET("/session/refresh", sc.HandleRefreshSession, session_service.AuthenticateCookie)
}
