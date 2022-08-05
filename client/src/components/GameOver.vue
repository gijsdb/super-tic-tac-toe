<template>
  <div class="w-screen h-screen absolute bg-black bg-opacity-60 flex items-center justify-center">
    <div class="bg-black border-white rounded-md p-12 text-white">
      Game over
      <div v-show="playerStore.Player.value.game.game_over.reason !== 'Player left the game'">
        <p v-show="playerStore.Player.value.game.winner == playerStore.Player.value.id">You win</p>
        <p v-show="playerStore.Player.value.game.winner !== playerStore.Player.value.id">You lose</p>
      </div>

      <p v-show="playerStore.Player.value.game.game_over.reason === 'Player left the game'">Other player left!</p>
      <button @click="returnToMenu">Menu</button>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import { useGameStore } from "../stores/game.js";
import { storeToRefs } from "pinia";

const router = useRouter();
const store = useGameStore();

let playerStore = storeToRefs(store);

const returnToMenu = () => {
  router.push("/");
};
</script>
