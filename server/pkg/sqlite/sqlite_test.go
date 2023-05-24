package sqlite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetUp() *SQLiteStore {
	return NewSQLiteStore(":memory:")
}

func TestCreate(t *testing.T) {
	ss := SetUp()

	player_id := ss.Create(&Player{ID: "1234", Email: "test", Picture: "test", GoogleID: "test"})
	assert.Equal(t, "1234", player_id)
}

func TestGetWhere(t *testing.T) {
	ss := SetUp()

	expected_id := ss.Create(&Player{ID: "1234", Email: "test", Picture: "test", GoogleID: "test"})
	actual_player, err := ss.GetWhere("1234", "id")
	assert.NoError(t, err)
	assert.Equal(t, expected_id, actual_player.ID)

	_, err = ss.GetWhere("abc", "id")
	assert.Error(t, err)
}

func TestGetAllPlayers(t *testing.T) {
	ss := SetUp()

	players_to_insert := []Player{{
		ID:       "1234",
		Email:    "test",
		Picture:  "test",
		GoogleID: "1",
	}, {
		ID:       "12345",
		Email:    "test",
		Picture:  "test",
		GoogleID: "2",
	},
	}

	ss.Create(&players_to_insert[0])
	ss.Create(&players_to_insert[1])

	actual, err := ss.GetAllPlayers()
	assert.NoError(t, err)

	assert.Equal(t, len(players_to_insert), len(actual))
}
