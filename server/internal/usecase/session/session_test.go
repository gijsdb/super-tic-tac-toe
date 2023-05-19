package session

import (
	"testing"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/stretchr/testify/assert"
)

var player_id string = "1234"

func SetUp() InteractorI {
	return NewService(repository.NewSessionRepository())
}

func TestIsSessionExpired(t *testing.T) {
	service := SetUp()
	token := service.Create(player_id, time.Now())
	assert.Equal(t, true, service.IsSessionExpired(token))

	token = service.Create(player_id, time.Now().Add(time.Minute*2))
	assert.Equal(t, false, service.IsSessionExpired(token))
}

func TestCreate(t *testing.T) {
	service := SetUp()

	expiry := time.Now()
	token := service.Create(player_id, expiry)

	actual, err := service.Get(token)
	assert.NoError(t, err)

	assert.Equal(t, &entity.Session{
		PlayerID: player_id,
		Expiry:   expiry,
	}, actual)
}

func TestGet(t *testing.T) {
	service := SetUp()

	//session does not exist
	_, err := service.Get("token")
	assert.Error(t, err)

	// session exists
	expiry := time.Now()
	token := service.Create(player_id, expiry)
	session, err := service.Get(token)
	assert.NoError(t, err)

	assert.Equal(t, &entity.Session{
		PlayerID: player_id,
		Expiry:   expiry,
	}, session)
}

func TestDelete(t *testing.T) {
	service := SetUp()

	token := service.Create(player_id, time.Now())
	service.Delete(token)

	_, err := service.Get(token)
	assert.Error(t, err)
}

func TestRefresh(t *testing.T) {
	service := SetUp()

	// Refresh token that does not exist
	_, _, err := service.Refresh("token")
	assert.Error(t, err)

	// Refresh token that exists
	expiry := time.Now()
	token := service.Create(player_id, expiry)
	session, new_token, err := service.Refresh(token)
	assert.NoError(t, err)
	assert.NotEqual(t, token, new_token)
	assert.Equal(t, true, expiry.Before(session.Expiry))
}

func TestGetTempSessionExpiry(t *testing.T) {
	//service := SetUp()

	//expiry := service.GetTempSessionExpiry()

	// TODO: need to pass the expiry to the session service
}
