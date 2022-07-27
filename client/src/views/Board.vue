<template>
  <div class="bg-gradient w-screen h-screen flex items-center justify-center">
    <div
      v-show="!playerStore.Player.value.game.full"
      class="
        bg-black bg-opacity-50
        p-12
        rounded-2xl
        shadow-2xl
        text-white
        font-bold
        text-center
      "
    >
      <p class="text-4xl py-4">
        Game {{ playerStore.Player.value.game.ID }} - Waiting for players...
      </p>
      <Loader />
    </div>

    <div
      v-show="playerStore.Player.value.game.full"
      class="flex flex-col space-y-10"
    >
      <div
        class="
          rounded-2xl
          shadow-2xl
          p-2
          text-white
          bg-black bg-opacity-60
          border-white border-4
        "
      >
        <p>Player ID: {{ playerStore.Player.value.id }}</p>
        <p>Game ID: {{ playerStore.Player.value.game.ID }}</p>
        <p>Your turn: {{ playerStore.Player.value.turn }}</p>
        <p>Game Full: {{ playerStore.Player.value.game.full }}</p>
      </div>
    </div>
    <Dice v-show="playerStore.Player.value.game.full" :clear="enableDice" />

    <div v-show="playerStore.Player.value.game.full" class="flex">
      <div
        class="
          z-40
          bg-black bg-opacity-80
          border-white border-2
          shadow-2xl
          p-4
          w-[40vw]
          h-[40vw]
          rounded-lg
          grid grid-cols-3
        "
      >
        <Square
          v-for="(square, idx) in playerStore.Player.value.game.game_board
            .squares"
          :key="idx"
          :square="square"
          :squareIdx="idx"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref, onBeforeUnmount } from "vue";
import { useGameStore } from "../stores/game.js";
import { useRouter } from "vue-router";
import { storeToRefs } from "pinia";
import APIClient from "../APIClient";

import Square from "../components/Square.vue";
import Dice from "../components/Dice.vue";
import Loader from "../components/Loader.vue";

const router = useRouter();
const store = useGameStore();

const { refreshGame } = store;

let playerStore = storeToRefs(store);
let enableDice = ref(false);
let waitForPlayerInterval = null;
// let waitForTurnInterval = null;

// const updateboard = async () => {
//   enableDice.value = true;
//   await refreshGame();
// };

// const fetchData = async () => {
//   storeFetched.value = false;
//   try {
//     // Send player id to verify if player is in game
//     const res = await APIClient.GetGameBoard(store.Player.game.ID);
//     game.value = res;
//     storeFetched.value = true;
//     if (game.value.player_turn === store.Player.id) {
//       store.Player.turn = true;
//     }
//   } catch (e) {
//     router.push("/");
//   }
// };

onMounted(() => {
  // game.value = store.Player.game;
  if (!playerStore.Player.value.game.full) {
    waitForPlayerInterval = setInterval(async () => {
      await refreshGame();
    }, 1000);
  } else {
    clearInterval(waitForPlayerInterval);
  }
});

onBeforeUnmount(async () => {
  // // TODO CALL API TO LEAVE GAME
  clearInterval(waitForPlayerInterval);
  // clearInterval(waitForTurnInterval);
});
</script>