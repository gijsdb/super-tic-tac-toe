<template>
  <div>
    <GamesList
      v-if="showGames"
      :games="games"
      @closeGamesList="showGames = false"
    />
    <div class="w-screen h-screen bg-gradient flex items-center justify-center">
      <div
        class="
          bg-black
          w-4/12
          py-48
          rounded-2xl
          shadow-2xl
          bg-opacity-70
          border-white border-4
          text-white
          flex
        "
      >
        <div
          class="w-full flex flex-col items-center justify-center text-center"
        >
          <h1 class="text-4xl font-white font-bold">SUPER TIC TAC TOE</h1>
          <div>
            <ul>
              <li>
                <button
                  @click="listGames()"
                  class="bg-[#1fddff] p-4 rounded-md text-white font-bold m-4"
                >
                  List Games
                </button>
              </li>
              <li>
                <button
                  @click="createGame()"
                  class="bg-[#1fddff] p-4 rounded-md text-white font-bold"
                >
                  Create Game
                </button>
              </li>
            </ul>
          </div>
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
    games.value = res;
    showGames.value = true;
  } catch (e) {
    console.log("Error listing games", e);
  }
};

const createGame = async () => {
  try {
    const res = await store.createGame();
    if (res != -1) {
      router.push("/game/" + res);
    }
  } catch (e) {
    console.log("Error creating game", e);
  }
};

onMounted(() => {
  store.registerClient();
});

window.onbeforeunload = async () => {
  await APIClient.RemovePlayer(store.Player.id);
};
</script>