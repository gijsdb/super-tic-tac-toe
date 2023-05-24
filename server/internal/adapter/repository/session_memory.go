package repository

import (
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/gijsdb/super-tic-tac-toe/pkg/memorystore"
)

type SessionMemoryRepoI interface {
	Create(session *entity.Session) string
	Get(token string) (*entity.Session, error)
	Delete(token string)
}

type SessionMemoryRepo struct {
	store *memorystore.SessionStore
}

func NewSessionRepo() SessionMemoryRepoI {
	return &SessionMemoryRepo{
		store: memorystore.NewSessionMemoryStore(),
	}
}

func (p *SessionMemoryRepo) Get(token string) (*entity.Session, error) {
	return p.store.Get(token)
}

func (p *SessionMemoryRepo) Create(session *entity.Session) string {
	return p.store.Create(session)
}

func (p *SessionMemoryRepo) Delete(token string) {
	p.store.Delete(token)
}
