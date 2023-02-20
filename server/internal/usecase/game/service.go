package game

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

type InteractorI interface {
	Index() []*entity.Game
	Get(id int64) *entity.Game
	CreateGame(creatingPlayer string) *entity.Game
	Join(gameId int64, playerId string) (*entity.Game, error)
	Leave(gameId int64, playerId string) *entity.Game
	UpdateBoard(gameId, square, circle int64, playerId string) *entity.Game
	RemoveCircle(gameId, square, circle int64, playerId string) *entity.Game
	RollDice(dice1, dice2 int, gameId int64) *entity.Game
}

func NewService(repo repository.GameRepositoryI) InteractorI {
	return &Service{
		repo: repo,
	}
}

type Service struct {
	repo repository.GameRepositoryI
}
