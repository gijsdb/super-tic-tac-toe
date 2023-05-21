<template>
  <div
    :class="{
      'capturedByOppositionBorder border-2':
        gameStoreRef.Player.value.game.game_board.squares[squareIdx].captured_by !== gameStoreRef.Player.value.id &&
        gameStoreRef.Player.value.game.game_board.squares[squareIdx].captured_by !== '',
      'capturedByYouBorder border-2':
        gameStoreRef.Player.value.game.game_board.squares[squareIdx].captured_by == gameStoreRef.Player.value.id &&
        gameStoreRef.Player.value.game.game_board.squares[squareIdx].captured_by !== '',
      'unCapturedBorder border-2': gameStoreRef.Player.value.game.game_board.squares[squareIdx].captured_by == '',
    }"
    class="lg:m-2 m-1 grid grid-cols-3 items-center"
  >
    <div
      :key="idx"
      v-for="(circle, idx) in gameStoreRef.Player.value.game.game_board.squares[squareIdx].circles"
      class="flex flex-col mx-auto items-center p-1.5 rounded-xl"
      :class="{
          'capturedByOppositionBorder border-2':
            idx + 3 == 7 &&
            gameStoreRef.Player.value.game.game_board.squares[squareIdx].captured_by !== gameStoreRef.Player.value.id &&
            gameStoreRef.Player.value.game.game_board.squares[squareIdx].captured_by !== '',
          'capturedByYouBorder border-2':
            idx + 3 == 7 &&
            gameStoreRef.Player.value.game.game_board.squares[squareIdx].captured_by == gameStoreRef.Player.value.id &&
            gameStoreRef.Player.value.game.game_board.squares[squareIdx].captured_by !== '',
          'unCapturedBorder border-2':
            idx + 3 == 7 && gameStoreRef.Player.value.game.game_board.squares[squareIdx].captured_by == '',
        }"
      @click="updateboard(idx)"
    >
      <span
        v-if="idx + 3 != 7"
        :style="{ color: colorStoreRef.ActiveTheme.value.Primary }"
        class="text-sm ml-1"
      >{{ idx + 3 }}</span>
      <span
        v-if="idx + 3 == 7"
        :style="{ color: colorStoreRef.ActiveTheme.value.Primary }"
        class="text-sm ml-1"
      >{{ squareIdx + 3 }}</span>
      <div
        class="w-8 h-8 rounded-full"
        :class="{
          'capturedByOpposition':
            gameStoreRef.Player.value.game.game_board.squares[squareIdx].circles[idx].selected_by !==
            gameStoreRef.Player.value.id &&
            gameStoreRef.Player.value.game.game_board.squares[squareIdx].circles[idx].selected_by !== '',
          'capturedByYou':
            gameStoreRef.Player.value.game.game_board.squares[squareIdx].circles[idx].selected_by ==
            gameStoreRef.Player.value.id &&
            gameStoreRef.Player.value.game.game_board.squares[squareIdx].circles[idx].selected_by !== '',
          'unCaptured': gameStoreRef.Player.value.game.game_board.squares[squareIdx].circles[idx].selected_by == '',
        }"
      ></div>
    </div>
  </div>
</template>

<style scoped>
.capturedByOpposition {
  background-color: v-bind(highlight);
}

.capturedByOppositionBorder {
  border-color: v-bind(highlight);
}

.capturedByYou {
  background-color: v-bind(highlightTwo);
}

.capturedByYouBorder {
  border-color: v-bind(highlightTwo);
}

.unCaptured {
  background-color: v-bind(primary);
}

.unCapturedBorder {
  border-color: v-bind(primary);
}
</style>

<script setup>
import { ref } from "vue"
import { storeToRefs } from "pinia";
import { useGameStore } from "../stores/game.js";
import { useColorStore } from "../stores/color.js";
import { CheckRules } from "../game/rules.js";

const gameStore = useGameStore();
const colorStore = useColorStore();

let gameStoreRef = storeToRefs(gameStore);
let colorStoreRef = storeToRefs(colorStore);

const { updateGameBoard, rollDice, removeCircle } = gameStore;

const emit = defineEmits(["ruleVerdict", "clearDice"]);

const props = defineProps({
  squareIdx: Number,
});

let highlightTwo = ref(colorStoreRef.ActiveTheme.value.HighlightTwo)
let highlight = ref(colorStoreRef.ActiveTheme.value.Highlight)
let primary = ref(colorStoreRef.ActiveTheme.value.Primary)

const updateboard = async (circleIdx) => {
  if (gameStoreRef.Player.value.turn) {
    if (gameStoreRef.Player.value.diceRolled) {
      try {
        let verdict = CheckRules(gameStoreRef, props.squareIdx, circleIdx, gameStore.diceTotal);
        // 2 = Remove opposition circle and roll again
        if (gameStore.diceTotal === 2) {
          // make request to remove an opposition circle
          if (verdict.allowed) {
            await removeCircle(props.squareIdx, circleIdx);

            gameStoreRef.Player.value.diceRolled = false;
            emit("clearDice");
            await rollDice(0, 0);
            emit("ruleVerdict", verdict);
            return;
          }
          emit("ruleVerdict", { allowed: false, reason: verdict.reason });
        } else {
          if (verdict.allowed) {
            await updateGameBoard(gameStoreRef.Player.value.id, props.squareIdx, circleIdx, gameStoreRef.Player.value.game.ID);
            gameStoreRef.Player.value.diceRolled = false;
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
