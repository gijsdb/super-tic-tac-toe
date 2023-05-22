package sqlite

import (
	"os"
	"testing"

	"github.com/inconshreveable/log15"
	"github.com/stretchr/testify/assert"
)

func SetUp() *SQLiteStore {
	return NewSQLiteStore()
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

	player_id := ss.Create(&Player{Email: "test", Picture: "test", GoogleID: "test", TempID: "1234"})
	assert.Equal(t, uint(1), player_id)
}
