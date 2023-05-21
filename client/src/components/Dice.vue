<template>
  <div class="bg-black bg-opacity-60 p-8 rounded-2xl flex flex-col items-center justify-center text-center">
    <p
      class="text-sm text-white font-bold py-2 w-10/12"
      v-if="message != null"
    >{{ message }}</p>
    <div v-if="!playerStore.Player.value.turn">
      <ul
        v-if="!playerStore.Player.value.game.last_roll != '0,0'"
        class="flex text-red-500 text-6xl"
        id="lastroll"
      ></ul>
    </div>
    <button
      :disabled="!playerStore.Player.value.turn || diceRolled"
      @click="rollDiceHandler"
      class="bg-green-500 rounded-md p-4 text-white font-bold disabled:bg-gray-500"
    >
      Roll dice
    </button>
    <ul
      class="flex text-white text-6xl"
      id="rolls"
    ></ul>
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
let message = ref(null);

const rollDiceHandler = async () => {
  let values = createDice();
  let totalRoll = values[0] + values[1];
  diceRolled.value = true;

  // check if there are any circles left uncaptured for that roll
  let availableCircles = [];
  let availableSquare = 0;
  let totalAvailableCircles = 0;

  // Loop over each square on the board
  store.Player.game.game_board.squares.forEach((square) => {
    // If a square is captured by a player, it is locked and no moves can be made on it so we ignore it from the count of circles left
    if (square.captured_by == "") {
      let circleCount = 0;
      let availableCircle = 0;
      // For each square, loop over the 9 circles
      square.circles.forEach((circle) => {
        // if the circle index is equal to the dice total or the dicetotal is 12 (any circle) AND the circle is not captured by a player
        if ((circle.Index + 3 == totalRoll || totalRoll == 12) && circle.selected_by == "") {
          // Record the circle
          circleCount = circleCount + 1;
          availableSquare++;
        }
        if (circle.selected_by == "") {
          availableCircle = availableCircle + 1;
        }
      });
      totalAvailableCircles = totalAvailableCircles + availableCircle;
      // After ALL of the squares circles are counted add the circleCount to an arrray
      availableCircles.push(circleCount);
    }
  });

  // If the roll is a 2 and there are no captured circles yet reroll
  if (totalRoll == 2 && totalAvailableCircles == 81) {
    message.value = "Please reroll, a 2 doesn't work at this time";
    diceRolled.value = false;
  } else if (totalRoll == 2 && totalAvailableCircles < 81) {
    message.value = null;
    await rollDice(values[0], values[1]);
    return;
  } else {
    let allowRoll = false
    // Loop over the available circles for each square that we recorded earlier
    availableCircles.forEach(async (val) => {
      // if a circle had been recorded the value would be greater than 1
      if (val > 0) {
        allowRoll = true
      } else {
        if (availableSquare == 0) {
          diceRolled.value = false;
          message.value = `Please reroll, a ${totalRoll} doesn't work at this time`;
        }
      }
    });
    if (allowRoll) {
      message.value = null;
      await rollDice(values[0], values[1]);
    }
  }
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
