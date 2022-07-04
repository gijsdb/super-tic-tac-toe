<template>
  <div class="w-screen h-screen bg-primary flex items-center justify-center">
    <div
      class="bg-secondary w-6/12 py-48 rounded-md border-red-500 border-4 flex"
    >
      <div class="w-6/12 flex justify-center items-center">
        <h1 class="text-4xl font-white">SUPER TIC TAC TOE</h1>
      </div>
      <div class="w-6/12 flex justify-center text-center">
        <ul>
          <li>
            <button
              @click="joinGame(4)"
              class="bg-green-500 p-4 rounded-md text-white font-bold m-4"
            >
              Join Game
            </button>
          </li>
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
</template>

<script setup>
import { useRouter } from "vue-router";
import APIClient from "../APIClient";
const router = useRouter();
const listGames = async () => {
  try {
    const res = await APIClient.ListGames();
    console.log("LISTED GAMES: ", res);
  } catch (e) {
    console.log("Error listing games", e);
  }
};

const createGame = async () => {
  try {
    const res = await APIClient.CreateGame();
    console.log("CREATED GAME: ", res);
  } catch (e) {
    console.log("Error creating game", e);
  }
};

const joinGame = async (id) => {
  try {
    router.push({ path: "/game/" + id });
  } catch (e) {
    console.log("Error creating game", e);
  }
};
</script>