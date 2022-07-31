import { defineStore } from 'pinia'
import APIClient from '../APIClient'

export const useGameStore = defineStore('game', {
  state: () => {
    return {
      Player: {
        id: 0,
        name: 'Player',
        inGame: true,
        turn: false,
        game: {},
      },
      // TODO ADD ERRORS
      Errors: {

      },
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
    async checkClient() {
      var cookieArr = document.cookie.split(";");
      for (var i = 0; i < cookieArr.length; i++) {
        var cookiePair = cookieArr[i].split("=");
        if (cookiePair[0].trim() === "client_id") {
          console.log("Player found")

          // Check cookie and store match
          if (!cookiePair[1] === this.Player.id) {
            return { allowed: false, reason: "Cookie does not match store" }
          }
          // Make sure player is active
          await APIClient.CreatePlayer(cookiePair[1])

          return { allowed: true, reason: "player active, cookie matches store" }
        }
      }
    },
    async createGame() {
      try {
        const res = await APIClient.CreateGame(this.Player.id);
        this.Player.inGame = res.ID
        this.Player.turn = true
        this.Player.game = res
        return true
      } catch (e) {
        console.log("Erroring creating game in store", e)
        return false
      }
    },
    async joinGame(gameId) {
      try {
        let res = await APIClient.JoinGame(gameId, this.Player.id)
        this.Player.inGame = res.ID
        this.Player.turn = false
        this.Player.game = res
        return true
      } catch (e) {
        console.log("Erroring joining game in store", e)
        return false
      }
    },
    async leaveGame(gameId, playerId) {
      this.Player.inGame = false
      this.Player.game = {}
      try {
        const res = await APIClient.LeaveGame(gameId, playerId)
        console.log("RES FROM LEAVING", res)

      } catch (e) {
        console.log("Error leaving game", e)
      }
    },
    async refreshGame() {
      try {
        const res = await APIClient.GetGame(this.Player.game.ID);
        this.Player.game = res;
        if (this.Player.game.player_turn.replace(/,/g, "") == this.Player.id) {
          this.Player.turn = true
        } else {
          this.Player.turn = false
        }
      } catch (e) {
        console.log("Error refreshing game in store", e)

        return e
      }
    },
    async updateGameBoard(player, square, circle, game) {
      try {
        let res = await APIClient.UpdateGameBoard(
          player,
          square,
          circle,
          game
        );
        this.Player.turn = false
        this.Player.game = res
      } catch (e) {
        console.log("Erroring updating game in store", e)
      }
    }
  }
})

