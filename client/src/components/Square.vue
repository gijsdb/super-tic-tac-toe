<template>
  <div
    :class="{
      'border-red-500 border-4':
        square.captured_by !== store.Player.id && square.captured_by !== -1,
      'border-blue-500 border-4':
        square.captured_by == store.Player.id && square.captured_by !== -1,
      'border-white border-4': square.captured_by == -1,
    }"
    class="m-2 grid grid-cols-3 items-center"
  >
    <div
      :key="idx"
      v-for="(circle, idx) in square.circles"
      class="flex mx-auto items-center p-1.5 rounded-xl"
      :class="{
        'border-red-500 border-4':
          idx + 3 == 7 &&
          square.captured_by !== store.Player.id &&
          square.captured_by !== -1,
        'border-blue-500 border-4':
          idx + 3 == 7 &&
          square.captured_by == store.Player.id &&
          square.captured_by !== -1,
        'border-white border-4': idx + 3 == 7 && square.captured_by == -1,
      }"
      @click="updateboard(idx)"
    >
      <div
        class="w-8 h-8 rounded-full"
        :class="{
          'bg-red-500':
            circle.selected_by !== store.Player.id && circle.selected_by !== -1,
          'bg-blue-500':
            circle.selected_by == store.Player.id && circle.selected_by !== -1,
          'bg-black': circle.selected_by == -1,
        }"
      ></div>
      <span v-if="idx + 3 != 7" class="text-sm text-white ml-1">{{
        idx + 3
      }}</span>
      <span v-if="idx + 3 == 7" class="text-sm text-white ml-1">{{
        squareIdx + 3
      }}</span>
    </div>
  </div>
</template>

<script setup>
import { useGameStore } from "../stores/game.js";
import APIClient from "../APIClient.js";
import { useRouter } from "vue-router";

const store = useGameStore();
const emit = defineEmits(["updateboard"]);
const router = useRouter();

const props = defineProps({
  square: Object,
  squareIdx: Number,
});

const updateboard = (circleIdx) => {
  if (store.Player.turn) {
    try {
      store.updateGameBoard(
        store.Player.id,
        props.squareIdx,
        circleIdx,
        router.currentRoute.value.params.id
      );
      emit("updateboard");
    } catch (e) {
      console.log("Error making request updateGameBoard");
    }

    return;
  }
  alert("It's not your turn!");
};
</script>