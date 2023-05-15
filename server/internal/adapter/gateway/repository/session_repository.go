package repository

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type SessionRepositoryI interface {
	Create(session *entity.Session) string
	Get(token string) (*entity.Session, error)
	Delete(token string)
}

type SessionMemoryRepository struct {
	store *memorystore.SessionStore
}

func NewSessionRepository() SessionRepositoryI {
	return &SessionMemoryRepository{
		store: memorystore.NewSessionMemoryStore(),
	}
}

func (p *SessionMemoryRepository) Get(token string) (*entity.Session, error) {
	return p.store.Get(token)
}

func (p *SessionMemoryRepository) Create(session *entity.Session) string {
	return p.store.Create(session)
}

func (p *SessionMemoryRepository) Delete(token string) {
	p.store.Delete(token)
}
