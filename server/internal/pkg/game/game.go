package game

func (g *Game) createNewState() (State, error) {
	gb := CreateGameBoard()
	state := State{
		PlayerTurn: 0,
		GameOver:   false,
		Winner:     2,
		GameID:     g.ID,
		GameBoard:  &gb,
	}

	return state, nil
}
