<template>
  <div
    :class="{
      'border-red-500 border-4':
        playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== playerStore.Player.value.id &&
        playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== -1,
      'border-blue-500 border-4':
        playerStore.Player.value.game.game_board.squares[squareIdx].captured_by == playerStore.Player.value.id &&
        playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== -1,
      'border-white border-4': playerStore.Player.value.game.game_board.squares[squareIdx].captured_by == -1,
    }"
    class="m-2 grid grid-cols-3 items-center"
  >
    <div
      :key="idx"
      v-for="(circle, idx) in playerStore.Player.value.game.game_board.squares[squareIdx].circles"
      class="flex mx-auto items-center p-1.5 rounded-xl"
      :class="{
        'border-red-500 border-4':
          idx + 3 == 7 &&
          playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== playerStore.Player.value.id &&
          playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== -1,
        'border-blue-500 border-4':
          idx + 3 == 7 &&
          playerStore.Player.value.game.game_board.squares[squareIdx].captured_by == playerStore.Player.value.id &&
          playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== -1,
        'border-white border-4':
          idx + 3 == 7 && playerStore.Player.value.game.game_board.squares[squareIdx].captured_by == -1,
      }"
      @click="updateboard(idx)"
    >
      <div
        class="w-8 h-8 rounded-full"
        :class="{
          'bg-red-500':
            playerStore.Player.value.game.game_board.squares[squareIdx].circles[idx].selected_by !==
              playerStore.Player.value.id &&
            playerStore.Player.value.game.game_board.squares[squareIdx].circles[idx].selected_by !== -1,
          'bg-blue-500':
            playerStore.Player.value.game.game_board.squares[squareIdx].circles[idx].selected_by ==
              playerStore.Player.value.id &&
            playerStore.Player.value.game.game_board.squares[squareIdx].circles[idx].selected_by !== -1,
          'bg-black': playerStore.Player.value.game.game_board.squares[squareIdx].circles[idx].selected_by == -1,
        }"
      ></div>
      <span v-if="idx + 3 != 7" class="text-sm text-white ml-1">{{ idx + 3 }}</span>
      <span v-if="idx + 3 == 7" class="text-sm text-white ml-1">{{ squareIdx + 3 }}</span>
    </div>
  </div>
</template>

<script setup>
import { useGameStore } from "../stores/game.js";
import { storeToRefs } from "pinia";
import { CheckRules } from "../game/rules.js";

const store = useGameStore();
let playerStore = storeToRefs(store);
const { updateGameBoard } = store;

const emit = defineEmits(["ruleVerdict"]);

const props = defineProps({
  squareIdx: Number,
});

const updateboard = async (circleIdx) => {
  if (playerStore.Player.value.turn) {
    try {
      let verdict = CheckRules(playerStore, props.squareIdx, circleIdx);
      if (verdict.allowed) {
        await updateGameBoard(playerStore.Player.value.id, props.squareIdx, circleIdx, playerStore.Player.value.game.ID);
      }
      emit("ruleVerdict", verdict);
    } catch (e) {
      console.log("Error making request updateGameBoard");
    }

    return;
  }
  alert("It's not your turn!");
};
</script>
