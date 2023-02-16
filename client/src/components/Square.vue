<template>
  <div
    :class="{
      'border-red-500 border-2':
        playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== playerStore.Player.value.id &&
        playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== -1,
      'border-blue-500 border-2':
        playerStore.Player.value.game.game_board.squares[squareIdx].captured_by == playerStore.Player.value.id &&
        playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== -1,
      'border-white border-2': playerStore.Player.value.game.game_board.squares[squareIdx].captured_by == -1,
    }"
    class="lg:m-2 m-1 grid grid-cols-3 items-center"
  >
    <div
      :key="idx"
      v-for="(circle, idx) in playerStore.Player.value.game.game_board.squares[squareIdx].circles"
      class="flex flex-col mx-auto items-center p-1.5 rounded-xl"
      :class="{
        'border-red-500 border-2':
          idx + 3 == 7 &&
          playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== playerStore.Player.value.id &&
          playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== -1,
        'border-blue-500 border-2':
          idx + 3 == 7 &&
          playerStore.Player.value.game.game_board.squares[squareIdx].captured_by == playerStore.Player.value.id &&
          playerStore.Player.value.game.game_board.squares[squareIdx].captured_by !== -1,
        'border-white border-2':
          idx + 3 == 7 && playerStore.Player.value.game.game_board.squares[squareIdx].captured_by == -1,
      }"
      @click="updateboard(idx)"
    >
      <span v-if="idx + 3 != 7" class="text-sm text-white ml-1">{{ idx + 3 }}</span>
      <span v-if="idx + 3 == 7" class="text-sm text-white ml-1">{{ squareIdx + 3 }}</span>
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
    </div>
  </div>
</template>

<script setup>
import { useGameStore } from "../stores/game.js";
import { storeToRefs } from "pinia";
import { CheckRules } from "../game/rules.js";

const store = useGameStore();
let playerStore = storeToRefs(store);
const { updateGameBoard, rollDice, removeCircle } = store;

const emit = defineEmits(["ruleVerdict", "clearDice"]);

const props = defineProps({
  squareIdx: Number,
});

const updateboard = async (circleIdx) => {
  if (playerStore.Player.value.turn) {
    if (playerStore.Player.value.diceRolled) {
      try {
        let verdict = CheckRules(playerStore, props.squareIdx, circleIdx, store.diceTotal);
        // 2 = Remove opposition circle and roll again
        if (store.diceTotal === 2) {
          // make request to remove an opposition circle
          if (verdict.allowed) {
            await removeCircle(props.squareIdx, circleIdx);

            playerStore.Player.value.diceRolled = false;
            emit("clearDice");
            await rollDice(0, 0);
            emit("ruleVerdict", verdict);
            return;
          }
          emit("ruleVerdict", { allowed: false, reason: verdict.reason });
        } else {
          if (verdict.allowed) {
            await updateGameBoard(playerStore.Player.value.id, props.squareIdx, circleIdx, playerStore.Player.value.game.ID);
            playerStore.Player.value.diceRolled = false;
            emit("clearDice");
            await rollDice(0, 0);
            emit("ruleVerdict", verdict);
            return;
          }
          emit("ruleVerdict", { allowed: false, reason: verdict.reason });
        }
      } catch (e) {
        console.log("Error making request updateGameBoard", e);
      }
      return;
    } else {
      alert("Roll the dice first!");
    }
  } else {
    alert("It's not your turn!");
  }
};
</script>
