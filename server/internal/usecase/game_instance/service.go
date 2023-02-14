package game_instance

import "github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"

type InteractorI interface {
}

func NewService(repo repository.GameInstanceRepositoryI) InteractorI {
	return &Service{
		repo: repo,
	}
}

type Service struct {
	repo repository.GameInstanceRepositoryI
}
