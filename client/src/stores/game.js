import { defineStore } from 'pinia'
import APIClient from '../APIClient'

export const useGameStore = defineStore('game', {
  state: () => {
    return {
      Player: {
        id: 0,
        name: 'Player',
        inGame: false,
        turn: false,
        diceRolled: false,
        game: {
          ID: -1,
          game_board: {},
          player_turn: -1,
          game_over: {
            over: false,
            reason: "",
          },
          winner: -1,
          players: "",
          full: false,
          last_roll: [0, 0]
        },
      },
      // TODO ADD ERRORS
      Errors: {

      },
    }
  },
  getters: {
    diceTotal(state) {
      let total = 0
      state.Player.game.last_roll.forEach((item, idx) => {

        total += item
      })
      return total
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
    removeClient() {
      APIClient.RemovePlayer(this.Player.id);
    },
    async checkClient() {
      if (this.Player.id === 0) {
        return { allowed: false, reason: "no client id" }
      }
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
    async leaveGame() {
      try {
        const res = await APIClient.LeaveGame(this.Player.game.ID, this.Player.id)
        this.Player.inGame = false
        this.Player.game = {}
        console.log("RES FROM LEAVING", res)

      } catch (e) {
        console.log("Error leaving game", e)
      }
    },
    async refreshGame() {
      try {
        const res = await APIClient.GetGame(this.Player.game.ID);
        this.Player.game = res;
        if (this.Player.game.player_turn == this.Player.id) {
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
    },
    async removeCircle(square, circle) {
      try {
        let res = await APIClient.RemoveCircle(
          this.Player.id,
          square,
          circle,
          this.Player.game.ID
        );
        if (res == "success") {
          this.Player.turn = false
          await this.refreshGame()
        } else {
          return res
        }
      } catch (e) {
        console.log("Erroring updating game in store", e)
      }
    },
    async rollDice(dice1, dice2) {
      try {
        let res = await APIClient.RollDice(dice1, dice2, this.Player.game.ID);
        this.Player.game = res
      } catch (e) {
        console.log("Erroring rolling dice in store", e)
      }
    }
  }
})

