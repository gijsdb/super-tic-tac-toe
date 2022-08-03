<template>
  <div>
    <button
      @click="closeGamesList"
      class="absolute bg-red-600 rounded-full w-12 h-12 flex justify-center items-center text-white right-5 top-5 font-bold"
    >
      X
    </button>

    <div class="flex flex-col items-center justify-center h-screen w-screen bg-gradient">
      <h1 class="font-bold p-8 text-white text-4xl">Games list</h1>
      <div
        class="bg-black bg-opacity-60 w-[50vw] h-[70vh] border-4 border-white rounded-2xl text-white items-center justify-center overflow-y-scroll space-y-4"
      >
        <div class="flex border-2 rounded-md mx-2 my-2" :key="game.ID" v-for="game in games">
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
    </div>
  </div>
</template>

<script setup>
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
