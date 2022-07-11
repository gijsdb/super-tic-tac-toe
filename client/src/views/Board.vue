<template>
  <div class="bg-primary w-screen h-screen flex items-center justify-center">
    <Dice />
    <div
      v-if="storeFetched"
      class="bg-secondary p-4 w-[50vw] h-[50vw] grid grid-cols-3 rounded-lg"
    >
      <Square
        v-for="(square, idx) in store.State.game_board.squares"
        :key="idx"
        :square="square"
        :squareIdx="idx"
        @updateboard="updateboard"
      />
    </div>
  </div>
</template>

<script setup>
import { onMounted, onBeforeUnmount, ref } from "vue";
import { useStore } from "../stores/game.js";
import { useRouter } from "vue-router";

import Square from "../components/Square.vue";
import Dice from "../components/Dice.vue";

const router = useRouter();
let storeFetched = ref(false);

const store = useStore();

const updateboard = async () => {
  storeFetched.value = false;
  await store.fetchState(router.currentRoute.value.params.id);
  storeFetched.value = true;
};

onMounted(async () => {
  storeFetched.value = false;
  await store.fetchState(router.currentRoute.value.params.id);
  storeFetched.value = true;
});

// onBeforeUnmount(async () => {
//   // CALL API TO LEAVE GAME
// });
</script>