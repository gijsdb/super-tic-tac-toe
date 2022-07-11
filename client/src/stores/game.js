import { defineStore } from 'pinia'
import APIClient from '../APIClient'

// useStore could be anything like useUser, useCart
// the first argument is a unique id of the store across your application
export const useStore = defineStore('game', {
  state: async () => {
    return {
      // all these properties will have their type inferred automatically
      ID: '',
      Players: 0,
      State: {},
    }
  },
  actions: {
    async fetchState(id) {
      try {
        console.log("ID IN STATE", id)
        const res = await APIClient.GetGameBoard(id)
        this.ID = res.id
        this.Players = res.players
        this.State = res.state
        console.log("RESPONSE", res)
      } catch (e) {
        console.log("Error fetching state", e)
      }
    },
    async updateState(player, square, circle, gameId) {
      try {
        const res = APIClient.UpdateGameBoard(player, square, circle, gameId)
        this.ID = res.id
        this.Players = res.players
        this.State = res.state
      } catch (e) {
        console.log("Error updating state", e)
      }
    }
  }
})

