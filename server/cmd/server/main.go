package main

import (
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/route"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println(err.Error())
		panic(1)
	}

	game_interactor := game.NewService(repository.NewGameRepository())
	player_interactor := player.NewService(repository.NewPlayerRepository())
	session_interactor := session.NewService(repository.NewSessionRepository())

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowCredentials},
		ExposeHeaders:    []string{echo.HeaderSetCookie, echo.HeaderAccessControlAllowCredentials},
		AllowCredentials: true,
	}))

	route.NewPlayerRouter(e, player_interactor, session_interactor)
	route.NewGameRouter(e, game_interactor, session_interactor, player_interactor)
	route.NewSessionRouter(e, session_interactor, player_interactor)

	// Load players from database into memorystore
	// Not sure if this is a good idea?
	player_interactor.LoadDBPlayersIntoMemory()

	e.Logger.Fatal(e.Start(":1323"))
}
