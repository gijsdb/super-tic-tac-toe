package player

import (
	"testing"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"github.com/stretchr/testify/assert"
)

func SetUp() InteractorI {
	return NewService(repository.NewPlayerRepository())
}

func TestGet(t *testing.T) {
	service := SetUp()

	expected := service.Create()
	actual, err := service.Get(expected)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual.ID)
}

func TestCreate(t *testing.T) {
	service := SetUp()

	player_id := service.Create()
	assert.NotEmpty(t, player_id)
}
