package memorystore

import (
	"fmt"
	"sync"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

type SessionStore struct {
	mutex    sync.Mutex
	sessions map[string]*entity.Session
}

func NewSessionMemoryStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[string]*entity.Session),
	}
}

func (ss *SessionStore) Create(session *entity.Session) string {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	token := NewUUID()
	ss.sessions[token] = session
	return token
}

func (ss *SessionStore) Get(token string) (*entity.Session, error) {
	session, exists := ss.sessions[token]
	if !exists {
		return nil, fmt.Errorf("session does not exist for token: %s", token)
	}
	return session, nil
}

func (ss *SessionStore) Delete(token string) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	delete(ss.sessions, token)
}
