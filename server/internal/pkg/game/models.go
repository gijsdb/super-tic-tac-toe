package game

// Manager is an overarching struct for accessing DB and Games.
type Manager struct {
	Games   []*Game      // List of games currently created (happening)
	Players map[int]bool // List of clients AKA players connected. int is id, bool is active
}

// State holds the game details including state of the board
type Game struct {
	ID         int        `json:"ID"`
	GameBoard  *GameBoard `json:"game_board" ` // JSON for gameboard
	PlayerTurn int        `json:"player_turn"`
	GameOver   bool       `json:"game_over"`
	Winner     int        `json:"winner"`
	Players    []int      `json:"players"` // List of string for the moment, needs changing
	Full       bool       `json:"full"`
	LastRoll   []int      `json:"last_roll"` // Last roll of the game, e.g. 5,5
}
