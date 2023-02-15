package player

import "github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"

type InteractorI interface {
	CreatePlayer(isNewPlayer int64) int64
	SetInactive(playerId int64)
}

func NewService(repo repository.PlayerRepositoryI) InteractorI {
	return &Service{
		repo: repo,
	}
}

type Service struct {
	repo repository.PlayerRepositoryI
}
