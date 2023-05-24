<template>
  <div class="w-screen my-auto flex flex-col lg:flex-row items-center justify-center gap-x-12 gap-y-4">
    <GameOver v-show="gameStoreRef.Player.value.game.game_over.over" />
    <div
      v-show="!gameStoreRef.Player.value.game.full"
      class="p-8 rounded-2xl shadow-2xl  font-bold text-center flex flex-col items-center gap-y-2 bg-main_color text-sub_alt_color"
    >
      <p class="text-4xl">Waiting for player...</p>
      <Loader />
      <button class="pt-2">Abandon</button>
    </div>

    <div
      v-show="gameStoreRef.Player.value.game.full && !gameStoreRef.Player.value.game.game_over.over"
      class="flex flex-col"
    >
      <div
        class="bg-sub_alt_color border-main_color px-8 rounded-lg rounded-b-none  border-2 border-b-0 text-center py-2 font-bold text-sm"
      >
        <p :class="{ 'text-error_color': !message.allowed, 'text-text_color': message.allowed }">
          {{ message.reason }}
        </p>
      </div>

      <div class="
        border-main_color
        bg-sub_alt_color
        border-2
        border-t-0
        shadow-2xl
        lg:px-4
        pb-4
        rounded-lg
        rounded-t-none
        grid
        grid-cols-3">
        <Square
          @clearDice="clearDiceHandler"
          v-for="(square, idx) in gameStoreRef.Player.value.game.game_board.squares"
          :key="idx"
          :squareIdx="idx"
          @ruleVerdict="updateMessage"
        />
      </div>
    </div>
    <div class="flex lg:flex-col gap-y-8 gap-x-8">
      <div
        v-show="gameStoreRef.Player.value.game.full"
        class="flex flex-col space-y-8"
      >
        <div
          class="bg-sub_alt_color border-main_color text-center rounded-2xl shadow-2xl p-2  font-bold py-6 flex flex-col"
        >
          <p
            v-show="gameStoreRef.Player.value.turn"
            class="text-caret_color"
          >Your turn</p>
          <p
            v-show="!gameStoreRef.Player.value.turn"
            class="text-error_color"
          >Not your turn</p>
          <p class="text-text_color">Game {{ gameStoreRef.Player.value.game.ID }}</p>
          <p
            class="text-text_color"
            v-for="(player, idx) in gameStoreRef.Player.value.game.players"
            :key="idx"
          >
            <span v-show="player != gameStoreRef.Player.value.id">VS {{ generateName(player) }}</span>
          </p>
        </div>
      </div>
      <Dice
        v-show="gameStoreRef.Player.value.game.full"
        @click="gameStoreRef.Player.value.diceRolled = true"
        :rolled="rolled"
      />
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref, onBeforeUnmount } from "vue";
import { useGameStore } from "../stores/game.js";
import { useRouter } from "vue-router";
import { storeToRefs } from "pinia";
import { generateName } from "../game/nameGenerator.js";

import Square from "../components/Square.vue";
import GameOver from "../components/GameOver.vue";
import Dice from "../components/Dice.vue";
import Loader from "../components/Loader.vue";

const router = useRouter();
const gameStore = useGameStore();

const { refreshGame, leaveGame } = gameStore;
let gameStoreRef = storeToRefs(gameStore);

let message = ref({ allowed: true, reason: "" });
let rolled = ref(false);
let feedbackLoop = null;
let waitForPlayerInterval = null;
let refreshInterval = null;

const beforeWindowUnload = (e) => {
  alert("Are you sure you want to abandon the game?")
  leaveGame();
  router.push("/");
};

const clearDiceHandler = () => {
  rolled.value = true;
  // Timeout to let dice clear
  setTimeout(() => {
    rolled.value = false;
  }, 500);
};

const checkMessage = () => {
  if (gameStoreRef.Player.value.game.game_over.over) {
    message.value = {
      allowed: true,
      reason: "Game over",
    };
    clearInterval(feedbackLoop);
  }
  if (!gameStoreRef.Player.value.turn) {
    message.value = {
      allowed: true,
      reason: "Waiting for other player",
    };
  }
  if (gameStoreRef.Player.value.turn && !gameStoreRef.Player.value.diceRolled) {
    message.value = {
      allowed: true,
      reason: "Your turn, roll the dice",
    };
  }
  if (gameStoreRef.Player.value.turn && gameStoreRef.Player.value.diceRolled) {
    message.value = {
      allowed: true,
      reason: "Make your move",
    };
  }
};

const updateMessage = async (verdict) => {
  message.value = verdict;

  if (feedbackLoop != null) {
    clearInterval(feedbackLoop);
  }
  feedbackLoop = setInterval(() => {
    checkMessage();
  }, 2500);
};

onMounted(() => {
  // prevents popup "Changes you made may not be saved" 
  window.onbeforeunload = null
  window.addEventListener("beforeunload", function (e) {
    console.log(e);
    beforeWindowUnload(e);
  });

  if (!gameStoreRef.Player.value.game.full) {
    waitForPlayerInterval = setInterval(async () => {
      await refreshGame();
    }, 1000);
  } else {
    clearInterval(waitForPlayerInterval);
  }
  refreshInterval = setInterval(async () => {
    let res = await refreshGame();
    if (res === "game abandoned") {
      router.push("/")
    }
  }, 1000);
  feedbackLoop = setInterval(() => {
    checkMessage();
  }, 1000);
});

onBeforeUnmount(() => {
  clearInterval(waitForPlayerInterval);
  clearInterval(refreshInterval);
  clearInterval(feedbackLoop);
  leaveGame();
});
</script>
