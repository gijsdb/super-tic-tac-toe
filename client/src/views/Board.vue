<template>
  <div class="bg-primary w-screen h-screen flex items-center justify-center">
    <Dice />
    <div
      v-if="storeFetched"
      class="bg-secondary p-4 w-[50vw] h-[50vw] grid grid-cols-3 rounded-lg"
    >
      <Square
        v-for="(square, idx) in store.gameboard.squares"
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
import { useStore } from "../stores/game.js";
import Square from "../components/Square.vue";
import Dice from "../components/Dice.vue";

let storeFetched = ref(false);

const store = useStore();

const updateboard = async () => {
  storeFetched.value = false;
  await store.fetchState();
  storeFetched.value = true;
};

onMounted(async () => {
  storeFetched.value = false;
  await store.fetchState();
  storeFetched.value = true;
});
</script>