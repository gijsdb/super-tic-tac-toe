package entity

// Manager is an overarching struct for accessing DB and Games.
type Manager struct {
	Games   []*Game      // List of games currently created (happening)
	Players map[int]bool // List of clients AKA players connected. int is id, bool is active
}
