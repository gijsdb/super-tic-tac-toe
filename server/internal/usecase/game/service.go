package game

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

type InteractorI interface {
	Index() []*entity.Game
	CreateGame(creatingPlayer int64) *entity.Game
}

func NewService(repo repository.GameRepositoryI) InteractorI {
	return &Service{
		repo: repo,
	}
}

type Service struct {
	repo repository.GameRepositoryI
}
