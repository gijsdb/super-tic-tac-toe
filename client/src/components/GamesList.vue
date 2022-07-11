<template>
  <div class="w-screen h-screen absolute bg-black bg-opacity-50">
    <div class="bg-red-500">
      <div :key="game.id" v-for="game in games">
        <p>Game: {{ game.id }}</p>
        <router-link :to="'/game/' + game.id"
          ><button @click="JoinGame(game.id)">Join</button></router-link
        >
      </div>
    </div>
  </div>
</template>


<script setup>
import APIClient from "../APIClient";
import { useRouter } from "vue-router";

const router = useRouter();

const props = defineProps({
  games: Array,
});

const JoinGame = async (id) => {
  try {
    await APIClient.JoinGame(id);
    router.push("/game/" + res);
  } catch (e) {
    console.log("Error joining game", e);
  }
};

console.log("LIST: ", props.games);
</script>