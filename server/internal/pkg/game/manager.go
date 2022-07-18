package game

import (
	"encoding/json"
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/pkg/db"
	"github.com/inconshreveable/log15"
)

func NewManager() *Manager {
	return &Manager{
		DB:    db.Init(),
		Games: map[int]*Game{},
	}
}

// Creates a new game in the manager and database
func (m *Manager) CreateGame(player string) (int, error) {
	var players = map[int]string{}
	players[0] = player
	game := Game{
		ID:      len(m.Games) + 1,
		Players: players,
	}
	state, err := game.createNewState()
	if err != nil {
		log15.Error("Error creating new state CreateNewGame()::game.go", "err", err)
	}
	game.State = &state
	result := m.DB.Create(&game)
	if result.Error != nil {
		log15.Debug("Error saving new game CreateNewGame()::game.go", "err", result.Error)
		return -1, err
	}

	m.Games[game.ID] = &game
	log15.Debug("game manager games after DB call CreateNewGame()", "game", m.Games[game.ID])
	return game.ID, nil
}

func (m *Manager) JoinGame(gameId int, player string) error {
	if m.Games[gameId].PlayerCnt >= 2 {
		return fmt.Errorf("game is full")
	}
	m.Games[gameId].PlayerCnt++
	m.Games[gameId].Players[m.Games[gameId].PlayerCnt] = player

	return nil
}

// GetGame returns a single game from the manager by ID
func (m *Manager) GetGame(id int) (Game, error) {
	log15.Debug("game manager games GetGame()", "game", m.Games[id])

	return *m.Games[id], nil
}

// GetGames is used by the List games UI to show all games
func (m *Manager) GetGames() error {
	var games []Game
	result := m.DB.Find(&games)
	if result.Error != nil {
		log15.Error("Error retrieving database records game::GetGames()")
	}

	for _, game := range games {
		m.Games[game.ID] = &game
	}

	return nil
}

// ListGames returns a JSON representation of the manager's games
func (m *Manager) ListGames() ([]byte, error) {
	err := m.GetGames()
	if err != nil {
		log15.Error("Error getting games list", "err", err)
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	var games []Game
	for _, game := range m.Games {
		games = append(games, *game)
	}
	bb, err := json.Marshal(games)
	if err != nil {
		log15.Error("Error marshalling games list", "err", err)
		return nil, err
	}
	return bb, nil
}

// CreateClient should return a unique identifier for the client to store in state
func (m *Manager) CreateClient() int {

	return 1
}

// ClearDB clears the database, run on start up for testing
func (m *Manager) ClearDB() {
	m.DB.Exec("DELETE FROM games")
	m.DB.Exec("DELETE FROM states")
}
