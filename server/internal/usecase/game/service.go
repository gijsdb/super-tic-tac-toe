package game

import (
	repo "github.com/gijsdb/super-tic-tac-toe/internal/adapter/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

type ServiceI interface {
	Index() []*entity.Game
	Get(id int64) *entity.Game
	CreateGame(creatingPlayer string) *entity.Game
	Join(gameId int64, playerId string) (*entity.Game, error)
	Leave(gameId int64, playerId string) error
	UpdateBoard(gameId, square, circle int64, playerId string) *entity.Game
	RemoveCircle(gameId, square, circle int64, playerId string) *entity.Game
	RollDice(dice1, dice2 int, gameId int64) *entity.Game
}

type GameService struct {
	g_mem_store repo.GameMemoryRepoI
	p_mem_store repo.PlayerMemoryRepoI
	p_db        repo.PlayerDatabaseRepoI
}

func NewGameService(g_mem_store repo.GameMemoryRepoI, p_mem_store repo.PlayerMemoryRepoI, p_db repo.PlayerDatabaseRepoI) ServiceI {
	return &GameService{
		g_mem_store: g_mem_store,
		p_mem_store: p_mem_store,
		p_db:        p_db,
	}
}
