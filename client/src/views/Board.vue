<template>
  <div class="bg-primary w-screen h-screen flex items-center justify-center">
    <div class="bg-red-600">
      <span>Players : {{ game.players }}</span>
      <ul>
        <li :key="idx" v-for="(field, idx) in game">
          <span v-show="idx !== 'game_board'">{{ idx }}:{{ field }}</span>
        </li>
      </ul>
    </div>
    <Dice />
    <div
      v-if="storeFetched"
      class="bg-secondary p-4 w-[50vw] h-[50vw] grid grid-cols-3 rounded-lg"
    >
      <Square
        v-for="(square, idx) in game.game_board.squares"
        :key="idx"
        :square="square"
        :squareIdx="idx"
        @updateboard="updateboard"
      />
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { useGameStore } from "../stores/game.js";
import { useRouter } from "vue-router";
import APIClient from "../APIClient";

import Square from "../components/Square.vue";
import Dice from "../components/Dice.vue";

const router = useRouter();
const store = useGameStore();

let storeFetched = ref(false);
let game = ref({});

const updateboard = async () => {
  storeFetched.value = false;
  const res = await APIClient.GetGameBoard(router.currentRoute.value.params.id);
  game.value = res;
  storeFetched.value = true;
};

onMounted(async () => {
  // store.newPlayer(router.currentRoute.value.params.id);
  updateboard();
  // const res = await APIClient.JoinGame(router.currentRoute.value.params.id);
  // console.log("JOINED", res);
});

// onBeforeUnmount(async () => {
//   // CALL API TO LEAVE GAME
// });
</script>