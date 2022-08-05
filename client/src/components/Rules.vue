<template>
  <div>
    <div
      @click="toggleRules"
      class="absolute z-50 bg-black bg-opacity-90 border-2 border-white w-12 h-12 flex items-center justify-center rounded-full top-5 left-5 text-white font-bold"
    >
      <span v-show="!showRules">?</span>
      <span v-show="showRules">X</span>
    </div>
    <div
      v-show="showRules"
      class="w-screen h-screen absolute bg-black bg-opacity-60 flex flex-col items-center justify-center text-white"
    >
      <h1 class="text-4xl font-bold">Super Tic Tac Toe Rules</h1>
      <ul class="text-2xl text-center space-y-2">
        <li>Roll dice</li>
        <li>Click stuff</li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

let showRules = ref(false);

const toggleRules = () => {
  showRules.value = !showRules.value;
};
</script>
