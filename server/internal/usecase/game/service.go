package game

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

// TODO
// Seperate game board related stuff
type InteractorI interface {
	Index() []*entity.Game
	Get(id int64) *entity.Game
	CreateGame(creatingPlayer string) *entity.Game
	Join(gameId int64, playerId string) (*entity.Game, error)
	Leave(gameId int64, playerId string) error
	UpdateBoard(gameId, square, circle int64, playerId string) *entity.Game
	RemoveCircle(gameId, square, circle int64, playerId string) *entity.Game
	RollDice(dice1, dice2 int, gameId int64) *entity.Game
}

func NewService(game_repo repository.GameRepositoryI, player_repo repository.PlayerRepositoryI) InteractorI {
	return &GameService{
		game_repo:   game_repo,
		player_repo: player_repo,
	}
}

type GameService struct {
	game_repo   repository.GameRepositoryI
	player_repo repository.PlayerRepositoryI
}
