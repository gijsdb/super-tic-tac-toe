package game

import (
	"testing"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/stretchr/testify/assert"
)

var player_id_one = "1234"
var player_id_two = "12345"

func SetUp() InteractorI {
	return NewService(repository.NewGameRepository())
}

func TestCreate(t *testing.T) {
	service := SetUp()

	game_board := newGameboard(player_id_one)
	expected := &entity.Game{
		ID:         1,
		GameBoard:  &game_board,
		PlayerTurn: player_id_one,
		GameOver: &entity.GameOver{
			Over:   false,
			Reason: "",
		},
		Winner:   "",
		Players:  []string{player_id_one},
		Full:     false,
		LastRoll: []int{0, 0},
	}

	actual := service.CreateGame(player_id_one)

	assert.Equal(t, expected, actual)
}

func TestGet(t *testing.T) {
	service := SetUp()

	// Get a game that exists
	expected := service.CreateGame(player_id_one)
	actual := service.Get(expected.ID)
	assert.Equal(t, expected, actual)

	// Get a game that does not exist
}

func TestIndex(t *testing.T) {
	service := SetUp()

	game_one := service.CreateGame(player_id_one)
	game_two := service.CreateGame(player_id_two)

	all_games := service.Index()

	assert.Contains(t, all_games, game_one)
	assert.Contains(t, all_games, game_two)
}

func TestJoin(t *testing.T) {
	service := SetUp()

	game := service.CreateGame(player_id_one)

	// a new player joins a game with one player
	updated_game, err := service.Join(game.ID, player_id_two)
	assert.NoError(t, err)

	assert.Equal(t, updated_game.Players, []string{player_id_one, player_id_two})
	assert.Equal(t, updated_game.Full, true)
	assert.Equal(t, updated_game.GameBoard.Player2, player_id_two)

	// a player joins a game that is full

	// a player tries to join a game that does not exist
}

func TestLeave(t *testing.T) {
	service := SetUp()

	// Two players in game, one leaves
	game := service.CreateGame(player_id_one)
	updated_game, err := service.Join(game.ID, player_id_two)
	assert.NoError(t, err)
	err = service.Leave(updated_game.ID, player_id_one)
	assert.NoError(t, err)

	expected := service.Get(game.ID)
	assert.Equal(t, &entity.GameOver{
		Over:    true,
		Reason:  "Player left the game",
		EndTime: expected.GameOver.EndTime,
	}, expected.GameOver)
	assert.Equal(t, []string{player_id_two}, expected.Players)

	// one player creates a game and leaves before another player joins
}

func TestRollDice(t *testing.T) {
	service := SetUp()

	game := service.CreateGame(player_id_one)
	service.Join(game.ID, player_id_two)
	updated_game := service.RollDice(4, 4, game.ID)

	assert.Equal(t, []int{4, 4}, updated_game.LastRoll)
}

func TestRemoveCircle(t *testing.T) {
	service := SetUp()

	game := service.CreateGame(player_id_one)
	service.Join(game.ID, player_id_two)

	//updated_game := service.RemoveCircle(game.ID, 3, 4, player_id_one)
}

func TestUpdateBoard(t *testing.T) {
	//service := SetUp()

	// Check for win if three squares in a row
	// Ensure player turn has changed
	//
}
