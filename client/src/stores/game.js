import { defineStore } from 'pinia'
import APIClient from '../APIClient'

// useStore could be anything like useUser, useCart
// the first argument is a unique id of the store across your application
export const useStore = defineStore('game', {
  state: async () => {
    return {
      // all these properties will have their type inferred automatically
      gameboard: {},
      gameover: false,
      playerTurn: -1,
      winner: -1
    }
  },
  actions: {
    async fetchState() {
      try {
        const res = await APIClient.GetGameBoard()
        this.gameboard = res.game_board
        this.gameover = res.game_over
        this.playerTurn = res.player_turn
        this.winner = res.winner
      } catch (e) {
        console.log("Error fetching state", e)
      }
      
    },
    async updateState(player, square, circle) {
      try {
        const res = APIClient.UpdateGameBoard(player, square, circle)
        this.gameboard = res.game_board
        this.gameover = res.game_over
        this.playerTurn = res.player_turn
        this.winner = res.winner
      } catch(e) {
        console.log("Error updating state", e)
      }
    }
  }
})

