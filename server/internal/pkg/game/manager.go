package game

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/inconshreveable/log15"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewManager() *Manager {
	return &Manager{
		DB:      initDB(),
		Games:   []*Game{},
		Players: make(map[int]bool),
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("super-tic-tac-toe.db"), &gorm.Config{})
	db.AutoMigrate(Game{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// Creates a new game in the manager and database, takes the player creating the game as param
func (m *Manager) CreateGame(creatingPlayer int) (*Game, error) {
	game := m.createNewGame(creatingPlayer)
	game.Players = fmt.Sprintf("%d,", creatingPlayer)
	game.PlayerTurn = fmt.Sprintf("%d,", creatingPlayer)

	result := m.DB.Create(&game)
	if result.Error != nil {
		log15.Debug("Error saving new game CreateNewGame()::game.go", "err", result.Error)
		return nil, result.Error
	}

	m.Games = append(m.Games, &game)
	log15.Debug("Created new game for player", "player", creatingPlayer, "game", game)
	return &game, nil
}

func (m *Manager) createNewGame(creatingPlayer int) Game {
	gb := CreateGameBoard(creatingPlayer)
	game := Game{
		PlayerTurn: "",
		GameOver:   false,
		Winner:     -1,
		GameBoard:  &gb,
		LastRoll:   "0",
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
			result := m.DB.Save(&game)
			if result.Error != nil {
				log15.Debug("Error saving game after join JoinGame()::game.go", "err", result.Error)
				return nil, result.Error
			}
			return game, nil
		}
	}
	return nil, fmt.Errorf("game not found")
}

func (m *Manager) LeaveGame(gameId int, leavingPlayer int) (*Game, error) {
	for _, game := range m.Games {
		if game.ID == gameId {
			// Remove player from Players on game
			playersSplit := strings.Split(game.Players, ",")
			for i, player := range playersSplit {
				playerInt, err := strconv.Atoi(player)
				if err != nil {
					log15.Crit("Error converting player to int while leaving game", "err", err)
				}
				if playerInt == leavingPlayer {
					playersSplit = append(playersSplit[:i], playersSplit[i+1:]...)
					game.Players = strings.Join(playersSplit, ",")
					break
				}
			}
			// If game is empty, delete it
			if len(playersSplit) == 0 {
				m.DB.Delete(&game)
				m.Games = append(m.Games[:gameId], m.Games[gameId+1:]...)
				return nil, fmt.Errorf("game is empty")
			}
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

// ClearDB clears the database, run on start up for testing
func (m *Manager) ClearDB() {
	m.DB.Exec("DELETE FROM games")
}
