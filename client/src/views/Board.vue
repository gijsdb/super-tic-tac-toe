<template>
  <div class=" w-screen my-auto flex flex-col lg:flex-row items-center justify-center gap-x-12 gap-y-4">
    <GameOver v-show="gameStoreRef.Player.value.game.game_over.over" />
    <div
      :style="{ backgroundColor: colorStoreRef.ActiveTheme.value.Secondary, color: colorStoreRef.ActiveTheme.value.Primary }"
      v-show="!gameStoreRef.Player.value.game.full"
      class="p-12 rounded-2xl shadow-2xl  font-bold text-center"
    >
      <p class="text-4xl py-4">Waiting for player...</p>
      <Loader />
    </div>

    <div
      v-show="gameStoreRef.Player.value.game.full && !gameStoreRef.Player.value.game.game_over.over"
      class="flex flex-col"
    >
      <div
        :style="{ backgroundColor: colorStoreRef.ActiveTheme.value.Secondary, borderColor: colorStoreRef.ActiveTheme.value.Highlight }"
        class="bg-opacity-80 px-8 rounded-lg rounded-b-none  border-2 border-b-0 text-center py-2 font-bold text-sm"
      >
        <p :class="{ 'text-red-600': !message.allowed, 'text-white': message.allowed }">
          {{ message.reason }}
        </p>
      </div>

      <div
        :style="{ backgroundColor: colorStoreRef.ActiveTheme.value.Secondary, borderColor: colorStoreRef.ActiveTheme.value.Highlight }"
        class="
        border-2
        border-t-0
        shadow-2xl
        lg:px-4
        pb-4
        rounded-lg
        rounded-t-none
        grid
        grid-cols-3"
      >
        <Square
          @clearDice="clearDiceHandler"
          v-for="( square, idx ) in  gameStoreRef.Player.value.game.game_board.squares "
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
          :style="{ backgroundColor: colorStoreRef.ActiveTheme.value.Secondary, borderColor: colorStoreRef.ActiveTheme.value.Highlight }"
          class="text-center rounded-2xl shadow-2xl p-2  font-bold py-6 flex flex-col"
        >
          <p
            v-show="gameStoreRef.Player.value.turn"
            class="text-green-500"
          >Your turn</p>
          <p
            v-show="!gameStoreRef.Player.value.turn"
            class="text-red-500"
          >Not your turn</p>
          <p>Game {{ gameStoreRef.Player.value.game.ID }}</p>
          <p
            v-for="( player, idx ) in  gameStoreRef.Player.value.game.players "
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
import { useColorStore } from "../stores/color.js";
import { useRouter } from "vue-router";
import { storeToRefs } from "pinia";
import { generateName } from "../game/nameGenerator.js";

import Square from "../components/Square.vue";
import GameOver from "../components/GameOver.vue";
import Dice from "../components/Dice.vue";
import Loader from "../components/Loader.vue";

const router = useRouter();
const gameStore = useGameStore();
const colorStore = useColorStore();

const { refreshGame, leaveGame } = gameStore;
let gameStoreRef = storeToRefs(gameStore);
let colorStoreRef = storeToRefs(colorStore);

let message = ref({ allowed: true, reason: "" });
let rolled = ref(false);
let feedbackLoop = null;
let waitForPlayerInterval = null;
let refreshInterval = null;

// const beforeWindowUnload = (e) => {
//   leaveGame();
//   router.push("/");
// };

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
  if (gameStoreRef.Player.value.turn) {
    message.value = {
      allowed: true,
      reason: "Your turn, roll the dice",
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
  //window.onbeforeunload = null
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
