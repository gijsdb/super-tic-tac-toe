package manager

import (
	"encoding/json"
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/game"
	"github.com/inconshreveable/log15"
)

func NewManager() *Manager {
	return &Manager{
		Games:   []*game.Game{},
		Players: make(map[int]bool),
	}
}

// Creates a new game in the manager and database, takes the player creating the game as param
func (m *Manager) CreateGame(creatingPlayer int) (*game.Game, error) {
	game := m.createNewGame(creatingPlayer)
	game.Players = append(game.Players, creatingPlayer)
	game.PlayerTurn = creatingPlayer

	m.Games = append(m.Games, &game)
	log15.Debug("Created new game for player", "player", creatingPlayer, "game", game)
	return &game, nil
}

func (m *Manager) createNewGame(creatingPlayer int) game.Game {
	gb := game.CreateGameBoard(creatingPlayer)
	game := game.Game{
		ID:         len(m.Games) + 1,
		PlayerTurn: -1,
		GameOver: &game.GameOver{
			Over:   false,
			Reason: "",
		},
		Winner:    -1,
		GameBoard: &gb,
		LastRoll:  []int{0, 0},
	}
	return game
}

func (m *Manager) JoinGame(gameId int, joiningPlayer int) (*game.Game, error) {
	for _, game := range m.Games {
		if game.ID == gameId {
			playerLength := len(game.Players)
			if playerLength >= 3 {
				return nil, fmt.Errorf("game is full")
			}
			game.Players = append(game.Players, joiningPlayer)
			game.GameBoard.Player2 = joiningPlayer
			game.Full = true

			return game, nil
		}
	}
	return nil, fmt.Errorf("game not found")
}

// GetGame returns a single game from the manager by ID
func (m *Manager) GetGame(idx int) (game.Game, error) {
	if idx >= 0 && idx < len(m.Games) {
		return *m.Games[idx], nil
	} else {
		return game.Game{}, fmt.Errorf("no games")
	}
}

// ListGames returns a JSON representation of the manager's games
func (m *Manager) ListGames() ([]byte, error) {
	gamesList := []game.Game{}

	for _, game := range m.Games {
		if game.Full || game.GameOver.Over {
			continue
		}
		gamesList = append(gamesList, *game)
	}

	bb, err := json.Marshal(gamesList)
	if err != nil {
		log15.Error("Error marshalling games list", "err", err)
		return nil, err
	}
	return bb, nil
}

// CreatePlayer returns a unique identifier for the client to store in state
func (m *Manager) CreatePlayer(id int) int {
	// 0 means new ID
	if id == 0 {
		id = len(m.Players) + 1
		m.Players[id] = true
		log15.Debug("Created new player", "player id", id, "players", m.Players)
		return id
	} else {
		// returning player, set to active
		m.Players[id] = true
		log15.Debug("Returning player", "player id", id, "players", m.Players)
		return id
	}
}

// RemovePlayer sets the player to inactive
func (m *Manager) SetPlayerInactive(playerId int) {
	m.Players[playerId] = false
	log15.Debug("Player set to inactive", "playerId", playerId, "all players", m.Players)
	return
}
