<template>
  <div class="bg-gradient w-screen h-screen flex items-center justify-center">
    <div
      v-show="!game.full"
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
      <p class="text-4xl py-4">Game {{ game.ID }} - Waiting for players...</p>
      <Loader />
    </div>

    <div v-show="game.full" class="flex flex-col space-y-10">
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
        <ul>
          <li :key="idx" v-for="(field, idx) in game">
            <span v-show="idx !== 'game_board'"
              ><span class="font-bold">{{ idx }}: </span>{{ field }}</span
            >
          </li>
        </ul>
      </div>
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
        <ul>
          <li :key="idx" v-for="(field, idx) in store.Player">
            <span v-show="idx !== 'game_board'"
              ><span class="font-bold">{{ idx }}: </span>{{ field }}</span
            >
          </li>
        </ul>
      </div>
    </div>
    <Dice v-show="game.full" :clear="enableDice" />

    <div v-show="game.full" class="flex">
      <div
        v-if="storeFetched"
        class="
          z-40
          bg-black bg-opacity-80
          border-white border-2
          shadow-2xl
          p-4
          w-[50vw]
          h-[50vw]
          rounded-lg
          grid grid-cols-3
        "
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
  </div>
</template>

<script setup>
import { onMounted, ref, onBeforeUnmount } from "vue";
import { useGameStore } from "../stores/game.js";
import { useRouter } from "vue-router";
import APIClient from "../APIClient";

import Square from "../components/Square.vue";
import Dice from "../components/Dice.vue";
import Loader from "../components/Loader.vue";

const router = useRouter();
const store = useGameStore();

let storeFetched = ref(false);
let game = ref({});
let enableDice = ref(false);
let waitForPlayerInterval = null;

const updateboard = async () => {
  enableDice.value = true;
  fetchData();
};

const fetchData = async () => {
  storeFetched.value = false;
  try {
    // Send player id to verify if player is in game
    const res = await APIClient.GetGameBoard(
      router.currentRoute.value.params.id
    );
    game.value = res;
    storeFetched.value = true;
  } catch (e) {
    router.push("/");
  }
};

const waitingForPlayer = () => {};

onMounted(async () => {
  fetchData();
  waitForPlayerInterval = setInterval(() => {
    if (game.value.full) {
      clearInterval(waitForPlayerInterval);
      return;
    }
    fetchData();
  }, 1000);
});

onBeforeUnmount(async () => {
  // TODO CALL API TO LEAVE GAME
  clearInterval(waitForPlayerInterval);
});
</script>