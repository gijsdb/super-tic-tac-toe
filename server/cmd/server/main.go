package main

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/route"
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game_instance"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	game_store := memorystore.NewGameMemoryStore()

	game_interactor := game.NewService(repository.NewGameRepository(game_store))
	route.NewGameRouter(e, game_interactor)

	player_interactor := player.NewService(repository.NewPlayerRepository())
	route.NewPlayerRouter(e, player_interactor)

	game_instance_interactor := game_instance.NewService(repository.NewGameInstanceRepository(game_store))
	route.NewGameInstanceRouter(e, game_instance_interactor)

	e.Logger.Fatal(e.Start(":1323"))
}
