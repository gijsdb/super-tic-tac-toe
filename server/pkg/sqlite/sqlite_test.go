package sqlite

import (
	"os"
	"testing"

	"github.com/inconshreveable/log15"
	"github.com/stretchr/testify/assert"
)

func SetUp() *SQLiteStore {
	return NewSQLiteStore("super_tic_tac_toe.db")
}

func CleanUp() {
	err := os.Remove("super_tic_tac_toe.db")
	if err != nil {
		log15.Error("error cleaning up after test")
	}
}

func TestCreate(t *testing.T) {
	ss := SetUp()
	defer CleanUp()

	player_id := ss.Create(&Player{Email: "test", Picture: "test", GoogleID: "test", OriginalID: "1234"})
	assert.Equal(t, uint(1), player_id)
}

func TestGetWhere(t *testing.T) {
	ss := SetUp()
	defer CleanUp()

	expected_id := ss.Create(&Player{Email: "test", Picture: "test", GoogleID: "test", OriginalID: "1234"})
	actual_player, err := ss.GetWhere("1234", "original_id")
	assert.NoError(t, err)
	assert.Equal(t, expected_id, actual_player.ID)

	_, err = ss.GetWhere("abc", "original_id")
	assert.Error(t, err)
}

func TestGetAllPlayers(t *testing.T) {
	ss := SetUp()
	defer CleanUp()

	players_to_insert := []Player{{
		Email:      "test",
		Picture:    "test",
		GoogleID:   "1",
		OriginalID: "1234",
	}, {
		Email:      "test",
		Picture:    "test",
		GoogleID:   "2",
		OriginalID: "12345",
	},
	}

	ss.Create(&players_to_insert[0])
	ss.Create(&players_to_insert[1])

	actual, err := ss.GetAllPlayers()
	assert.NoError(t, err)

	assert.Equal(t, len(players_to_insert), len(actual))
}
