<template>
  <div class="w-screen h-screen bg-gradient flex items-center justify-center">
    <div class="w-10/12 flex flex-col text-white items-center">
      <h1 class="text-4xl font-white font-bold py-4">Super Tic Tac Toe</h1>
      <div
        v-show="games.length > 0"
        class="bg-black bg-opacity-60 border-4 p-8 my-4 border-white rounded-2xl text-white overflow-y-scroll space-y-4 w-6/12 max-h-40vh"
      >
        <div
          class="flex border-2 rounded-md my-2 w-8/12 mx-auto"
          :key="game.ID"
          v-for="game in games"
          v-show="!game.game_over.over"
        >
          <div class="w-8/12 mx-2 flex flex-col justify-center text-xl">
            <p>Game: {{ game.ID }}</p>
          </div>
          <button
            class="flex-grow p-4 border-l-2 border-white font-bold text-xl"
            :class="{ 'bg-red-600': game.full, 'bg-green-500': !game.full }"
            @click="JoinGame(game.ID)"
            :disabled="game.full"
          >
            <p v-if="!game.full">Join</p>
            <p v-if="game.full">Full</p>
          </button>
        </div>
      </div>
      <button @click="createGameHandler()" class="bg-[#1fddff] p-4 rounded-md text-white font-bold w-2/12">
        Create Game
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from "vue";
import { storeToRefs } from "pinia";

import { useRouter } from "vue-router";
import APIClient from "../APIClient";
import { useGameStore } from "../stores/game.js";

const router = useRouter();
const store = useGameStore();
const { registerClient, removeClient, createGame, joinGame } = store;
let playerStore = storeToRefs(store);

let getGamesLoop;
let games = ref([]);

const listGames = async () => {
  try {
    const res = await APIClient.ListGames();
    games.value = res;
    showGames.value = true;
  } catch (e) {
    console.log("Error listing games", e);
  }
};

const JoinGame = async (gameId) => {
  try {
    let res = await store.joinGame(gameId);
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
    router.push("/game/" + playerStore.Player.value.game.ID);
  }
};

onMounted(async () => {
  await registerClient();
  await listGames();

  getGamesLoop = setInterval(async () => {
    await listGames();
  }, 3000);
});

onUnmounted(() => {
  clearInterval(getGamesLoop);
});

window.onbeforeunload = async () => {
  removeClient();
};
</script>
