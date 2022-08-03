package game

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/inconshreveable/log15"
)

func NewManager() *Manager {
	return &Manager{
		Games:   []*Game{},
		Players: make(map[int]bool),
	}
}

// Creates a new game in the manager and database, takes the player creating the game as param
func (m *Manager) CreateGame(creatingPlayer int) (*Game, error) {
	game := m.createNewGame(creatingPlayer)
	game.Players = fmt.Sprintf("%d,", creatingPlayer)
	game.PlayerTurn = fmt.Sprintf("%d,", creatingPlayer)

	// db := db.Get()
	// result := db.Ins.Create(&game)
	// if result.Error != nil {
	// 	log15.Debug("Error saving new game CreateNewGame()::game.go", "err", result.Error)
	// 	return nil, result.Error
	// }

	m.Games = append(m.Games, &game)
	log15.Debug("Created new game for player", "player", creatingPlayer, "game", game)
	return &game, nil
}

func (m *Manager) createNewGame(creatingPlayer int) Game {
	gb := CreateGameBoard(creatingPlayer)
	game := Game{
		ID:         len(m.Games) + 1,
		PlayerTurn: "",
		GameOver:   false,
		Winner:     -1,
		GameBoard:  &gb,
		LastRoll:   "0,0",
	}
	return game
}

func (m *Manager) JoinGame(gameId int, joiningPlayer int) (*Game, error) {
	for _, game := range m.Games {
		if game.ID == gameId {
			playerLength := strings.Split(game.Players, ",")
			if len(playerLength) >= 3 {
				return nil, fmt.Errorf("game is full")
			}
			game.Players = game.Players + fmt.Sprintf("%d,", joiningPlayer)
			game.GameBoard.Player2 = joiningPlayer
			game.Full = true

			// db := db.Get()
			// result := db.Ins.Save(&game)
			// if result.Error != nil {
			// 	log15.Debug("Error saving game after join JoinGame()::game.go", "err", result.Error)
			// 	return nil, result.Error
			// }
			return game, nil
		}
	}
	return nil, fmt.Errorf("game not found")
}

// GetGame returns a single game from the manager by ID
func (m *Manager) GetGame(idx int) (Game, error) {
	if idx >= 0 && idx < len(m.Games) {
		return *m.Games[idx], nil
	} else {
		return Game{}, fmt.Errorf("no games")
	}
}

// // GetGames is used by the List games UI to show all games
// func (m *Manager) GetGames() error {
// 	var games []Game
// 	db := db.Get()
// 	result := db.Ins.Find(&games)
// 	if result.Error != nil {
// 		log15.Error("Error retrieving database records game::GetGames()")
// 	}

// 	for _, game := range games {
// 		m.Games = append(m.Games, &game)
// 	}

// 	return nil
// }

// ListGames returns a JSON representation of the manager's games
func (m *Manager) ListGames() ([]byte, error) {
	bb, err := json.Marshal(m.Games)
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
