<template>
  <div class="bg-gradient w-screen h-screen flex items-center justify-center">
    <div class="flex flex-col space-y-10">
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
    <Dice :clear="enableDice" />

    <div class="flex">
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
let enableDice = ref(false);

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

onMounted(async () => {
  fetchData();
});

// onBeforeUnmount(async () => {
//   // CALL API TO LEAVE GAME
// });
</script>