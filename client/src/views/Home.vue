<template>
  <div class="flex justify-center my-auto">
    <div class="flex flex-col gap-y-4 text-white items-center bg-black bg-opacity-50 rounded-md p-8 shadow-2xl">
      <h1 class="text-4xl font-white font-bold">Super Tic Tac Toe</h1>
      <div
        v-show="gamesAvailable"
        class="p-8 my-4 text-white overflow-y-scroll space-y-4 max-h-40vh border-y custom-scroll-bar"
      >
        <div
          class="flex border-2 rounded-md"
          :key="game.ID"
          v-for="game in games"
          v-show="!game.game_over.over"
        >
          <div class="mx-2 flex flex-col justify-center text-xl px-4">
            <p>Game waiting for player</p>
          </div>
          <button
            class="p-4 border-l-2 border-white font-bold text-xl px-6"
            :class="{ 'bg-red-600': game.full, 'bg-green-500': !game.full }"
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
        class="bg-[#1fddff] p-4 rounded-md text-white font-bold"
      >Create New
        Game</button>
    </div>
  </div>
</template>

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

<script setup>
import { ref, onMounted, onUnmounted, computed } from "vue";
import { storeToRefs } from "pinia";

import { useRouter } from "vue-router";
import APIClient from "../APIClient";
import { useGameStore } from "../stores/game.js";

const router = useRouter();
const store = useGameStore();
const { registerClient, createGame, joinGame, getPlayer } = store;
let playerStore = storeToRefs(store);

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
