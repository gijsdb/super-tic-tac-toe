<template>
  <div>
    <button
      @click="toggleRules"
      :class="showRules ? 'border-2 border-white' : 'border-0'"
      class="absolute z-50 bg-black bg-opacity-60 w-12 h-12 flex items-center justify-center rounded-full top-5 left-5 text-white font-bold shadow-2xl"
    >
      <span v-show="!showRules">?</span>
      <span v-show="showRules">X</span>
    </button>
    <div
      v-show="showRules"
      class="z-40 w-screen h-screen absolute bg-black bg-opacity-95 flex flex-col items-center justify-center text-gray-200"
    >
      <div class="flex flex-col items-center text-center mx-24 gap-y-4">
        <h1 class="text-4xl font-bold underline">Super Tic Tac Toe Rules</h1>
        <p class="text-2xl leading-10">Super Tic Tac Toe is a variation of regular tic tac toe or 'three in a row'.</p>

        <ul class="text-xl text-center space-y-4 list-disc">
          <li>The gameboard contains a grid of 9 squares, each square has a grid of 9 circles inside.</li>
          <li>
            Players take turns rolling dice and capturing circles to form 3 in a row horizontally, diagonally or vertically.
          </li>
          <li>The dice decides what circles a player can select.</li>
          <li>
            Once a player gets 3 in a row or has a total of 5 circles in a square, the square is captured by the player.
          </li>
          <li>The game is won by capturing 3 squares vertically, horizontally or diagonally.</li>
          <li>Rolling a 7 lets the player select any center circle of a square.</li>
          <li>Rolling a 12 lets the player select any circle in any square.</li>
          <li>Rolling a 2 lets the player remove a circle of their opponent in an uncaptured square.</li>
        </ul>
      </div>
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
