package game

import "github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"

type InteractorI interface {
}

func NewService(repo repository.GameRepositoryI) InteractorI {
	return &Service{
		repo: repo,
	}
}

type Service struct {
	repo repository.GameRepositoryI
}
