<template>
  <div class="bg-black bg-opacity-60 border-white border-4 p-8 mx-8 rounded-2xl">
    <div v-if="!playerStore.Player.value.turn">
      <p v-if="playerStore.Player.value.game.last_roll == 0">Waiting for other player to roll dice</p>
      <p>{{ store.diceAsArray }}</p>
      <p>{{ `&#x268${store.diceAsArray[0]};` }};</p>
    </div>
    <button
      :disabled="!playerStore.Player.value.turn"
      @click="rollDiceHandler"
      class="bg-green-500 rounded-md p-4 text-white font-bold disabled:bg-gray-500"
    >
      Roll dice
    </button>
    <ul class="flex text-white text-6xl" id="rolls"></ul>
  </div>
</template>

<script setup>
import { onMounted, ref, watch } from "vue";
import { useGameStore } from "../stores/game.js";
import { storeToRefs } from "pinia";

const store = useGameStore();
let playerStore = storeToRefs(store);
const { rollDice } = store;

const rollDiceHandler = async () => {
  let values = createDice();
  console.log("DIE ", values[0], values[1]);
  await rollDice(values[0], values[1], playerStore.Player.value.game.ID);
};

const getRandomValue = () => {
  return Math.floor(Math.random() * 6);
};

const createDice = () => {
  clearDice();
  let dice1 = getRandomValue();
  let dice2 = getRandomValue();
  const roll1 = document.createElement("li");
  const roll2 = document.createElement("li");
  roll1.className = "roll";
  roll2.className = "roll";
  roll1.innerHTML = `&#x268${dice1};`;
  roll2.innerHTML = `&#x268${dice2};`;
  document.getElementById("rolls").appendChild(roll1);
  document.getElementById("rolls").appendChild(roll2);
  return [dice1 + 1, dice2 + 1];
};

const clearDice = () => {
  const dice = document.getElementsByClassName("roll");
  if (dice[0] != undefined) {
    dice[0].remove();
    dice[0].remove();
  }
};
</script>
