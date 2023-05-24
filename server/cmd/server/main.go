package main

import (
	"flag"
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/api/route"
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/game"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/player"
	"github.com/gijsdb/super-tic-tac-toe/internal/usecase/session"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	flag_seed_db := flag.Bool("seed", false, "Seeds the database with users")
	flag.Parse()

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println(err.Error())
		panic(1)
	}

	p_mem_store := repository.NewPlayerMemoryRepo()
	p_db := repository.NewPlayerDatabaseRepo("../../super_tic_tac_toe.db")
	g_service := game.NewGameService(repository.NewGameMemoryRepo(), p_mem_store, p_db)
	p_service := player.NewPlayerService(p_mem_store, p_db)
	s_service := session.NewSessionService(repository.NewSessionRepo())

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowCredentials},
		ExposeHeaders:    []string{echo.HeaderSetCookie, echo.HeaderAccessControlAllowCredentials},
		AllowCredentials: true,
	}))

	route.NewPlayerRouter(e, p_service, s_service)
	route.NewGameRouter(e, g_service, s_service, p_service)
	route.NewSessionRouter(e, s_service, p_service)

	if *flag_seed_db {
		p_service.SeedPlayers()
	}

	p_service.LoadDBPlayersIntoMemory()

	e.Logger.Fatal(e.Start(":1323"))
}
