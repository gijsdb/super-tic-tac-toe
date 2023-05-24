package player

import (
	"testing"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/stretchr/testify/assert"
)

func SetUp() ServiceI {
	return NewPlayerService(repository.NewPlayerMemoryRepo(), repository.NewPlayerDatabaseRepo(":memory:"))
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

func TestUpdate(t *testing.T) {
	service := SetUp()

	player_id := service.Create()
	player, _ := service.Get(player_id)
	assert.Empty(t, player.Email)

	player = service.Update(&entity.Player{
		ID:    player_id,
		Email: "test",
	})

	assert.Equal(t, "test", player.Email)
}

func TestGetByEmail(t *testing.T) {
	service := SetUp()

	player_id := service.Create()
	player, err := service.Get(player_id)
	assert.NoError(t, err)

	// exisiting player
	expected, err := service.GetByEmail(player.Email)
	assert.NoError(t, err)
	assert.Equal(t, expected, player)

	// non-existant player
	_, err = service.GetByEmail("test")
	assert.Error(t, err)
}
