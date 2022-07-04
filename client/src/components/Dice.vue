<template>
  <div class="bg-secondary mr-8 p-2 rounded-md">
    <button
      v-if="!rolled"
      @click="rollDice"
      class="bg-green-500 rounded-md p-2 border-2 border-white"
    >
      Roll dice
    </button>
    <ul class="flex text-white text-6xl" id="rolls"></ul>
  </div>
</template>

<script setup>
import { ref } from "vue";

let rolled = ref(false);

const rollDice = () => {
  rolled.value = true;
  createDice();
  createDice();
};

const getRandomValue = () => {
  return Math.floor(Math.random() * 6);
};

const createDice = () => {
  const roll = document.createElement("li");
  roll.innerHTML = `&#x268${getRandomValue()};`;
  roll.innerHTML = `&#x268${getRandomValue()};`;
  document.getElementById("rolls").appendChild(roll);
};
</script>