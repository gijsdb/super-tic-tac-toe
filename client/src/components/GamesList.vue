<template>
  <div class="w-screen h-screen absolute bg-black bg-opacity-50">
    <button
      @click="closeGamesList"
      class="
        absolute
        bg-red-600
        rounded-full
        w-12
        h-12
        flex
        justify-center
        items-center
        text-white
        right-5
        top-5
        font-bold
      "
    >
      X
    </button>
    <div class="bg-red-500">
      <div :key="game.ID" v-for="game in games">
        <p>Game: {{ game.ID }}</p>

        <button @click="JoinGame(game.ID)">Join</button>
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

const JoinGame = async (id) => {
  try {
    let res = await APIClient.JoinGame(id, store.Player.id);
    if (res == "true") {
      router.push("/game/" + id);
    } else {
      router.push("/");
    }
  } catch (e) {
    console.log("Error joining game", e);
    return;
  }
};

console.log("LIST: ", props.games);
</script>