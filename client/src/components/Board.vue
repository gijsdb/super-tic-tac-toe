<template>
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
</template>

<script setup>
import { onMounted, ref } from "vue";
import { useStore } from "../stores/game.js";
import Square from "./Square.vue";

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