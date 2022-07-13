<template>
  <div>
    <GamesList v-if="showGames" :games="games" />
    <div class="w-screen h-screen bg-primary flex items-center justify-center">
      <div
        class="
          bg-secondary
          w-6/12
          py-48
          rounded-md
          border-red-500 border-4
          flex
        "
      >
        <div class="w-6/12 flex justify-center items-center">
          <h1 class="text-4xl font-white">SUPER TIC TAC TOE</h1>
          {{ store.Player.id }}
        </div>
        <div class="w-6/12 flex justify-center text-center">
          <ul>
            <li>
              <button
                @click="listGames()"
                class="bg-green-500 p-4 rounded-md text-white font-bold m-4"
              >
                List Games
              </button>
            </li>
            <li>
              <button
                @click="createGame()"
                class="bg-green-500 p-4 rounded-md text-white font-bold"
              >
                Create Game
              </button>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import GamesList from "../components/GamesList.vue";
import APIClient from "../APIClient";
import { useGameStore } from "../stores/game.js";

const router = useRouter();
const store = useGameStore();

let showGames = ref(false);
let games = ref([]);

const listGames = async () => {
  try {
    const res = await APIClient.ListGames();
    console.log("LISTED GAMES: ", res);
    games.value = res;
    showGames.value = true;
  } catch (e) {
    console.log("Error listing games", e);
  }
};

const createGame = async () => {
  try {
    const res = await APIClient.CreateGame(store.Player.id);
    const res2 = await APIClient.JoinGame(res, store.Player.id);
    router.push("/game/" + res);
    console.log("CREATED GAME: ", res, res2);
  } catch (e) {
    console.log("Error creating game", e);
  }
};

onMounted(() => {
  store.identifyPlayer();
  console.log(store.Player);
});
</script>