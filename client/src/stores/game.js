import { defineStore } from 'pinia'
import APIClient from '../APIClient'

export const useGameStore = defineStore('game', {
  state: () => {
    return {
      Player: {
        id: '',
        name: 'Player',
        inGame: false,
        turn: false,
        diceRolled: false,
        game: {
          ID: -1,
          game_board: {},
          player_turn: '',
          game_over: {
            over: false,
            reason: '',
          },
          winner: '',
          players: '',
          full: false,
          last_roll: [0, 0]
        },
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
    // Check for cookie temp_account, if it exists ask for a session for the player id in the cookie
    // If it does not exist, create a new player
    async registerClient() {
      var cookieArr = document.cookie.split(";");
      for (var i = 0; i < cookieArr.length; i++) {
        var cookiePair = cookieArr[i].split("=");
        if (cookiePair[0].trim() === "temp_account") {
          try {
            await APIClient.CreateSession(cookiePair[1])
            this.Player.id = cookiePair[1]
            return
          } catch (e) {
            console.log("could not create session for returning player, creating new temp user")
          }
        }
      }
      try {
        const id = await APIClient.CreatePlayer()
        this.Player.id = id
      } catch (e) {
        console.log("Error creating player in store", e)
      }
    },
    // Replace with deleting session for user AKA logout
    // removeClient() {
    //   APIClient.RemovePlayer(this.Player.id);
    // },
    // REPLACE WITH REFRESHING SESSION
    async checkClient() {
      if (this.Player.id === '') {
        return { allowed: false, reason: "no client id" }
      }
      var cookieArr = document.cookie.split(";");
      for (var i = 0; i < cookieArr.length; i++) {
        var cookiePair = cookieArr[i].split("=");
        if (cookiePair[0].trim() === "client_id") {
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
        console.log("Error creating game in store", e)
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
        console.log("Error joining game in store", e)
        return false
      }
    },
    async leaveGame() {
      try {
        const res = await APIClient.LeaveGame(this.Player.game.ID, this.Player.id)
        this.Player.inGame = false
        this.Player.game = {}
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

        this.Player.turn = false
        this.Player.game = res

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
    },
    async OAuthLogin() {
      try {
        await APIClient.OAuthLogin()
        for (var i = 0; i < cookieArr.length; i++) {
          var cookiePair = cookieArr[i].split("=");
          if (cookiePair[0].trim() === "client_id") {
            this.Player.id = cookiePair[1]
          }
        }
      } catch (e) {
        console.log("Erroring performing OAuth login", e)
      }
    }
  }
})

