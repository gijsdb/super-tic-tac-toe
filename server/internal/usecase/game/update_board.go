package game

import (
	"fmt"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"github.com/inconshreveable/log15"
)

func (s *GameService) UpdateBoard(gameId, square, circle int64, playerId string) *entity.Game {
	game := s.g_mem_store.Get(gameId)

	game.GameBoard.Squares[square].Circles[circle].SelectedBy = playerId
	updatedSquare := game.GameBoard.CheckCirclesCondition(game.GameBoard.Squares[square])
	game.GameBoard.Squares[square] = updatedSquare
	game.GameBoard.CheckSquareCondition()

	// Check for winner
	if game.GameBoard.Winner != "" {
		game.Winner = game.GameBoard.Winner
		game.GameOver.Reason = fmt.Sprintf("Player %s wins!", game.Winner)
		game.GameOver.Over = true
		game.GameOver.EndTime = time.Now().Unix()

		player, err := s.p_db.GetWhere(game.Winner, "id")
		if err != nil {
			log15.Error("Could not find winner in db")
			// player not found in database, must be temporary player who's stats we don't track
			return s.g_mem_store.Update(game)
		}
		// Increment player wins count
		player.Wins++
		s.p_mem_store.Update(player)
		s.p_db.Update(player)

		// Find loser and increment losses if they're not a temporary account
		for _, player_id := range game.Players {
			if player_id != game.Winner {
				losing_player, err := s.p_db.GetWhere(player_id, "id")
				if err != nil {
					log15.Error("Could not find loser in db")

					// player not found in database, must be temporary player who's stats we don't track
					return s.g_mem_store.Update(game)
				}
				losing_player.Losses++
				s.p_mem_store.Update(losing_player)
				s.p_db.Update(losing_player)
			}
		}
	}

	// Change turn
	if playerId == game.Players[0] {
		game.PlayerTurn = game.Players[1]
	} else {
		game.PlayerTurn = game.Players[0]
	}

	return s.g_mem_store.Update(game)
}
