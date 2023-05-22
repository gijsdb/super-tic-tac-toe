<template>
  <div class="flex justify-center my-auto">
    <div class="flex flex-col gap-y-4 items-center bg-opacity-50 rounded-md p-8 mt-[-3em]">
      <h1 class="text-4xl font-white font-bold text-main_color">Super Tic Tac Toe</h1>
      <div
        v-show="gamesAvailable"
        class="p-8 my-4 overflow-y-scroll space-y-4 max-h-40vh border-y border-caret_color custom-scroll-bar"
      >
        <div
          class="flex border-2 border-text_color rounded-md"
          :key="game.ID"
          v-for="game in games"
          v-show="!game.game_over.over"
        >
          <div class="mx-2 flex flex-col justify-center text-xl px-4 text-main_color">
            <p>Game waiting for player</p>
          </div>
          <button
            class="p-4 border-l-2 border border-text_color font-bold text-xl px-6 bg-sub_color text-text_color"
            @click="JoinGame(game.ID)"
            :disabled="game.full"
          >
            <p v-if="!game.full">Join</p>
            <p v-if="game.full">Full</p>
          </button>
        </div>
      </div>
      <button
        @click="createGameHandler()"
        class="p-4 rounded-md font-bold text-text_color bg-caret_color"
      >Create New
        Game</button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed } from "vue";
import { storeToRefs } from "pinia";

import { useRouter } from "vue-router";
import APIClient from "../APIClient";
import { useGameStore } from "../stores/game.js";


const router = useRouter();
const gameStore = useGameStore();

const { registerClient, createGame, joinGame, getPlayer } = gameStore;
let gameStoreRef = storeToRefs(gameStore);

let getGamesLoop;
let games = ref([]);

const gamesAvailable = computed({
  get() {
    let res = false;
    games.value.forEach((game) => {
      if (!game.full && !game.game_over.over) {
        res = true;
      }
    });

    return res;
  },
});

const listGames = async () => {
  try {
    const res = await APIClient.ListGames();
    games.value = res;
  } catch (e) {
    console.log("Error listing games", e);
  }
};

const JoinGame = async (gameId) => {
  try {
    let res = await gameStore.joinGame(gameId);
    if (res === true) {
      router.push("/game/" + gameId);
    }
  } catch (e) {
    console.log("Error joining game", e);
    return;
  }
};

const createGameHandler = async () => {
  let res = await createGame();
  if (res) {
    router.push("/game/" + gameStoreRef.Player.value.game.ID);
  }
};

onMounted(async () => {
  await registerClient();
  await getPlayer();
  await listGames();

  getGamesLoop = setInterval(async () => {
    await listGames();
  }, 3000);

  // setTimeout(() => {
  //   resetFlashMessage()
  // }, 3000)
});

onUnmounted(() => {
  clearInterval(getGamesLoop);
});

window.onbeforeunload = async () => {
  // removeClient();
};
</script>

<style scoped>
/* Firefox */
.custom-scroll-bar {
  scrollbar-width: thin;
  scrollbar-color: #ffffff #000000;
}

/* Chrome, Edge, and Safari */
.custom-scroll-bar::-webkit-scrollbar {
  width: 10px;
}

.custom-scroll-bar::-webkit-scrollbar-track {
  background: transparent;
}

.custom-scroll-bar::-webkit-scrollbar-thumb {
  background-color: #ffffff;
  border-radius: 4px;
  border: 3px solid #ffffff;
}
</style>
