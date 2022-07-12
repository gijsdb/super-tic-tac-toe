import { defineStore } from 'pinia'
import APIClient from '../APIClient'

// useStore could be anything like useUser, useCart
// the first argument is a unique id of the store across your application
export const useGameStore = defineStore('game', {
  state: () => {
    return {
      // all these properties will have their type inferred automatically
      Players: [{ id: 0, gameId: 0 }],

    }
  },
  actions: {
    newPlayer(gameId) {
      this.Players.push({ id: this.Players.length, gameId: gameId })
      console.log("PLAYERS", this.Players)
    },
    async updateState(player, square, circle, gameId) {
      // try {
      //   const res = APIClient.UpdateGameBoard(player, square, circle, gameId)
      //   this.ID = res.id
      //   this.Players = res.players
      //   this.State = res.state
      // } catch (e) {
      //   console.log("Error updating state", e)
      // }
    }
  }
})

