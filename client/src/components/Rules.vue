<template>
  <div>
    <div
      v-show="gameStoreRef.ShowRules.value"
      class="w-screen h-screen fixed flex flex-col items-center justify-center text-text_color"
    >

      <div class="flex flex-col items-center text-center gap-y-4 bg-sub_alt_color p-4 rounded-xl">
        <button
          @click="handleCloseRules()"
          class="mb-2 rounded-full h-8 w-8 font-bold self-start text-sm bg-bg_color text-text_color"
        >X</button>
        <h1 class="text-4xl font-bold text-main_color">Super Tic Tac Toe Rules</h1>
        <p class="text-2xl leading-10">Super Tic Tac Toe is a variation of regular tic tac toe or 'three in a row'.</p>

        <ul class="text-xl space-y-4 list-disc list-inside text-left">
          <li>The gameboard contains a grid of 9 squares, each square has a grid of 9 circles inside.</li>
          <li>
            Players take turns rolling dice and capturing circles to form 3 in a row horizontally, diagonally or
            vertically.
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
import { storeToRefs } from "pinia";
import { useColorStore } from "../stores/color.js";
import { useGameStore } from "../stores/game.js";

const colorStore = useColorStore();
const gameStore = useGameStore();
let colorStoreRef = storeToRefs(colorStore);
let gameStoreRef = storeToRefs(gameStore);

const { toggleShowRules } = gameStore;

const handleCloseRules = () => {
  toggleShowRules()
}
</script>
