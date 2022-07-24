package game

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/inconshreveable/log15"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewManager() *Manager {
	return &Manager{
		DB:      Init(),
		Games:   []*Game{},
		Players: make(map[int]bool),
	}
}

func Init() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("super-tic-tac-toe.db"), &gorm.Config{})
	db.AutoMigrate(Game{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

// Creates a new game in the manager and database
func (m *Manager) CreateGame(player int) (int, error) {
	game, err := m.createNewState()
	game.Players = fmt.Sprintf("%d,", player)
	if err != nil {
		log15.Error("Error creating new state CreateNewGame()::game.go", "err", err)
	}
	result := m.DB.Create(&game)
	if result.Error != nil {
		log15.Debug("Error saving new game CreateNewGame()::game.go", "err", result.Error)
		return -1, err
	}

	m.Games = append(m.Games, &game)
	return len(m.Games), nil
}

func (m *Manager) createNewState() (Game, error) {
	gb := CreateGameBoard()
	game := Game{
		PlayerTurn: 0,
		GameOver:   false,
		Winner:     2,
		GameBoard:  &gb,
	}

	return game, nil
}

func (m *Manager) JoinGame(gameId int, player int) error {
	for _, game := range m.Games {
		if game.ID == gameId {
			playerLength := strings.Split(game.Players, ",")
			log15.Debug("Player length", "playerLength", playerLength)
			if len(playerLength) >= 3 {
				return fmt.Errorf("game is full")
			}
			game.Players = game.Players + fmt.Sprintf("%d,", player)
			result := m.DB.Save(&game)
			if result.Error != nil {
				log15.Debug("Error saving game after join JoinGame()::game.go", "err", result.Error)
				return result.Error
			}
		}
	}
	return nil
}

// GetGame returns a single game from the manager by ID
func (m *Manager) GetGame(id int) (Game, error) {
	// log15.Debug("GAMES ARE", "games", *m.Games[0])
	return *m.Games[id-1], nil
}

// GetGames is used by the List games UI to show all games
func (m *Manager) GetGames() error {
	var games []Game
	result := m.DB.Find(&games)
	if result.Error != nil {
		log15.Error("Error retrieving database records game::GetGames()")
	}

	for _, game := range games {
		m.Games = append(m.Games, &game)
	}

	return nil
}

// ListGames returns a JSON representation of the manager's games
func (m *Manager) ListGames() ([]byte, error) {
	// err := m.GetGames()
	// if err != nil {
	// 	log15.Error("Error getting games list", "err", err)
	// 	return nil, err
	// }

	// var games []Game
	// for _, game := range m.Games {
	// 	games = append(games, *game)
	// }
	bb, err := json.Marshal(m.Games)
	if err != nil {
		log15.Error("Error marshalling games list", "err", err)
		return nil, err
	}
	return bb, nil
}

// CreatePlayer should return a unique identifier for the client to store in state
func (m *Manager) CreatePlayer(id int) int {
	if id == 0 {
		id = len(m.Players) + 1
		m.Players[id] = true
		return id
	} else {
		m.Players[id] = true
		return id
	}
}

func (m *Manager) RemovePlayer(id int) {
	m.Players[id] = false

	return
}

// ClearDB clears the database, run on start up for testing
func (m *Manager) ClearDB() {
	m.DB.Exec("DELETE FROM games")
}
