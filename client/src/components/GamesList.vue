<template>
  <div class="absolute flex items-center justify-center h-screen w-screen">
    <button
      @click="closeGamesList"
      class="absolute bg-red-600 rounded-full w-12 h-12 flex justify-center items-center text-white right-5 top-5 font-bold"
    >
      X
    </button>
    <div
      class="bg-black bg-opacity-60 w-[50vw] h-[50vh] border-4 border-white rounded-2xl flex flex-col text-white items-center justify-center"
    >
      <h1 class="text-2xl font-bold">Games list</h1>
      <div :key="game.ID" v-for="game in games">
        <p>Game: {{ game.ID }}</p>
        <p>Full: {{ game.full }}</p>
        <button v-if="!game.full" @click="JoinGame(game.ID)">Join</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import APIClient from "../APIClient";
import { useRouter } from "vue-router";
import { useGameStore } from "../stores/game.js";

const emit = defineEmits(["closeGamesList"]);

const router = useRouter();
const store = useGameStore();

const props = defineProps({
  games: Array,
});

const closeGamesList = () => {
  emit("closeGamesList");
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
</script>
