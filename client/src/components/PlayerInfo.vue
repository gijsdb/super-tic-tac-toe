<template>
  <div class="absolute right-5 top-5 bg-black bg-opacity-50 rounded-xl shadow-2xl p-4">
    <div class="flex justify-end items-center">
      <div class="flex-col mr-4 text-sm text-white">
        <p>
          <span class="font-bold mr-2">Player {{ playerStore.Player.value.id }}</span>
        </p>
        <p
          class="font-bold"
          :class="{
            'text-red-500': !playerStore.Player.value.inGame,
            'text-green-500': playerStore.Player.value.inGame != false,
          }"
        >
          In Game
        </p>
      </div>
      <img class="rounded-full w-14 h-14" src="../assets/avatar.png" />
    </div>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import { useGameStore } from "../stores/game.js";
import { storeToRefs } from "pinia";

const router = useRouter();
const store = useGameStore();
let playerStore = storeToRefs(store);
</script>
