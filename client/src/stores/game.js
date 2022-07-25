import { defineStore } from 'pinia'
import APIClient from '../APIClient'

export const useGameStore = defineStore('game', {
  state: () => {
    return {
      // all these properties will have their type inferred automatically
      Player: {
        id: 0,
        name: 'Player',
        inGame: -1,
        turn: false,
      }
    }
  },
  actions: {
    async registerClient() {
      var cookieArr = document.cookie.split(";");
      for (var i = 0; i < cookieArr.length; i++) {
        var cookiePair = cookieArr[i].split("=");
        if (cookiePair[0].trim() === "client_id") {
          this.Player.id = cookiePair[1]
          await APIClient.CreatePlayer(cookiePair[1])

          return
        }
      }
      console.log("NEW PLAYER")
      const id = await APIClient.CreatePlayer(0)

      document.cookie = "client_id=" + id + ";SameSite=none;Secure=false";

      this.Player.id = id
    },
    async createGame() {
      try {
        const res = await APIClient.CreateGame(this.Player.id);
        if (res === -1) {
          console.log("Erroring creating game in store", e)
          return -1
        }
        this.Player.inGame = res
        return res
      } catch {
        console.log("Erroring creating game in store", e)
        return -1
      }
    },
    async joinGame(gameId) {
      try {
        let res = await APIClient.JoinGame(gameId, this.Player.id)
        if (res == "true") {
          this.Player.inGame = gameId
          return true
        } else {
          return false
        }
      } catch (e) {
        console.log("Erroring joining game in store", e)
        return false
      }
    }
  }
})

