package game

import (
	"encoding/json"

	"github.com/inconshreveable/log15"
	"gorm.io/datatypes"
)

func (m *Manager) CreateNewGame() (int, error) {
	log15.Debug("Creating new game")
	game := Game{
		Players: 1,
	}
	state, err := game.createNewState()
	if err != nil {

	}
	game.State = state
	result := m.DB.Create(&game)
	m.Games = append(m.Games, game)
	if result.Error != nil {
		return -1, err
	}
	return game.ID, nil
}

func (m *Manager) GetGame(id int) (Game, error) {
	var game Game
	result := m.DB.Where("id = ?", id).First(&game)
	game.State.JsonToGameboard(game.State.GameBoard)
	if result.Error != nil {
		log15.Error("Error retrieving database record game::GetGame()")
		return Game{}, result.Error
	}
	return game, nil
}

func (m *Manager) JSONGames() ([]byte, error) {
	err := m.GetGames()
	if err != nil {
		log15.Error("Error getting games list", "err", err)
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	bb, err := json.Marshal(m.Games)
	if err != nil {
		log15.Error("Error marshalling games list", "err", err)
		return nil, err
	}
	return bb, nil
}

func (m *Manager) GetGames() error {
	var games []Game
	result := m.DB.Find(&games)
	if result.Error != nil {
		log15.Error("Error retrieving database records game::GetGames()")
	}

	m.Games = games
	return nil
}

func (g *Game) createNewState() (State, error) {
	state := State{
		PlayerTurn: 0,
		GameOver:   false,
		Winner:     2,
	}

	gb := CreateGameBoard()

	json, err := g.gameboardToJson(gb)
	if err != nil {
		return State{}, err
	}
	state.GameBoard = json

	return state, nil
}

func (g *Game) gameboardToJson(gameboard GameBoard) ([]byte, error) {
	bb, err := json.Marshal(gameboard)
	if err != nil {
		log15.Debug("Error marshalling JSON game.go::gameToJson()")
		return nil, err
	}

	return bb, nil
}

func (s *State) JsonToGameboard(bb datatypes.JSON) (*GameBoard, error) {
	gb := &GameBoard{}
	err := json.Unmarshal(bb, gb)
	if err != nil {
		log15.Error("Error unmarshalling JSON game.go::jsonToGameboard()")
		return &GameBoard{}, err
	}
	return gb, nil
}
