package game

import "github.com/gijsdb/super-tic-tac-toe/internal/entity"

func (s *GameService) Get(id int64) *entity.Game {
	return s.g_mem_store.Get(id)
}

func (s *GameService) Index() []*entity.Game {
	return s.g_mem_store.Index()
}
