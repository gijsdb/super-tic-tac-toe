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
    async identifyPlayer() {
      await fetch('https://api.ipify.org?format=json')
        .then(x => x.json())
        .then(({ ip }) => {

          this.Player.id = ip;
        });
      let device = ""
      const ua = navigator.userAgent;
      if (/(tablet|ipad|playbook|silk)|(android(?!.*mobi))/i.test(ua)) {
        device = "tablet";
      }
      if (
        /Mobile|iP(hone|od)|Android|BlackBerry|IEMobile|Kindle|Silk-Accelerated|(hpw|web)OS|Opera M(obi|ini)/.test(
          ua
        )
      ) {
        device = "mobile";
      } else {
        device = "desktop";
      }
      this.Player.id = this.Player.id + device
    },
  }
})

