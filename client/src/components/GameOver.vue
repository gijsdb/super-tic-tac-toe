<template>
  <div class="w-screen h-screen absolute bg-black bg-opacity-90 flex items-center justify-center">
    <div class="border-white rounded-md p-12 text-white flex flex-col text-center gap-y-2">
      <p class="text-4xl font-bold">Game over</p>

      <div v-show="playerStore.Player.value.game.game_over.reason !== 'Player left the game'">
        <p v-show="playerStore.Player.value.game.winner == playerStore.Player.value.id">You win</p>
        <p
          v-show="
            playerStore.Player.value.game.winner !== playerStore.Player.value.id &&
            playerStore.Player.value.game.game_over.reason !== 'Player left the game'
          "
        >
          You lose
        </p>
      </div>

      <p class="text-xl" v-show="playerStore.Player.value.game.game_over.reason === 'Player left the game'">
        Other player left!
      </p>
      <button class="bg-[#1fddff] font-bold rounded-md text-2xl py-2 my-2" @click="returnToMenu">Menu</button>
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
