<template>
  <div class="bg-black bg-opacity-60 border-white border-4 p-8 mx-8 rounded-2xl">
    <div v-if="!playerStore.Player.value.turn">
      <ul v-if="!playerStore.Player.value.game.last_roll != '0,0'" class="flex text-red-500 text-6xl" id="lastroll"></ul>
    </div>
    <button
      :disabled="!playerStore.Player.value.turn || diceRolled"
      @click="rollDiceHandler"
      class="bg-green-500 rounded-md p-4 text-white font-bold disabled:bg-gray-500"
    >
      Roll dice
    </button>
    <ul class="flex text-white text-6xl" id="rolls"></ul>
  </div>
</template>

<script setup>
import { ref, watch } from "vue";
import { useGameStore } from "../stores/game.js";
import { storeToRefs } from "pinia";

const props = defineProps({
  rolled: Boolean,
});

const store = useGameStore();
let playerStore = storeToRefs(store);
const { rollDice } = store;
let diceRolled = ref(false);

const rollDiceHandler = async () => {
  diceRolled.value = true;
  let values = createDice();
  await rollDice(values[0], values[1]);
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

const showLastRoll = () => {
  const dice1 = document.createElement("li");
  const dice2 = document.createElement("li");
  dice1.className = "roll";
  dice2.className = "roll";
  dice1.innerHTML = `&#x268${playerStore.Player.value.game.last_roll[0] - 1};`;
  dice2.innerHTML = `&#x268${playerStore.Player.value.game.last_roll[1] - 1};`;
  document.getElementById("lastroll").appendChild(dice1);
  document.getElementById("lastroll").appendChild(dice2);
};

const clearDice = () => {
  const dice = document.getElementsByClassName("roll");

  if (dice[0] != undefined) {
    dice[0].remove();
    dice[0].remove();
  }
};

watch(
  () => store.Player.game.last_roll,
  (newRoll, oldRoll) => {
    if (playerStore.Player.value.turn) {
      return;
    }
    if (newRoll[0] == 0 && newRoll[1] == 0) {
      console.log("AHHHHHH");
      return;
    }
    if (newRoll[0] != oldRoll[0] && newRoll[1] != oldRoll[1]) {
      showLastRoll();
    }
  }
);

watch(
  () => store.Player.turn,
  () => {
    diceRolled.value = false;
  }
);

watch(
  () => props.rolled,
  () => {
    clearDice();
  }
);
</script>
