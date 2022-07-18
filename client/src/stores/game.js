import { defineStore } from 'pinia'
import APIClient from '../APIClient'

export const useGameStore = defineStore('game', {
  state: () => {
    return {
      // all these properties will have their type inferred automatically
      Player: {
        id: 0,
        name: 'Player',
        turn: false,
        score: 0
      }
    }
  },
  actions: {
    async registerClient() {
      let currentId = localStorage.getItem('playerId')
      if (currentId === null) {
        console.log("NEW PLAYER")
        const id = await APIClient.CreatePlayer()
        this.Player.id = id
        localStorage.setItem('playerId', id)
      }

      console.log("currentId: ", currentId)
    },
  }
})

