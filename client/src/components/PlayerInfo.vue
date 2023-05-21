<template>
  <div class="bg-opacity-50 rounded-xl bg-black p-4">
    <div class="flex justify-end items-center">
      <div class="flex-col mr-4 text-sm text-white">
        <p>
          <span class="font-bold mr-2">{{ generateName(playerStore.Player.value.id) }} </span>
        </p>
        <p class="font-bold" :class="{
          'text-red-500': !playerStore.Player.value.inGame,
          'text-green-500': playerStore.Player.value.inGame != false,
        }">
          In Game
        </p>
        <a v-show="playerStore.Player.value.temporary" href="http://localhost:1323/login">login</a>
        <button v-show="!playerStore.Player.value.temporary" @click="handleLogout()">logout</button>
      </div>
      <img v-if="playerStore.Player.value.picture == null" class="rounded-full w-14 h-14" src="../assets/avatar.png" />
      <img v-if="playerStore.Player.value.picture != null" class="rounded-full w-14 h-14"
        :src="playerStore.Player.value.picture" />
    </div>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import { useGameStore } from "../stores/game.js";
import { storeToRefs } from "pinia";
import { generateName } from "../game/nameGenerator.js";

const router = useRouter();

const store = useGameStore();
const { logOut } = store;
let playerStore = storeToRefs(store);

const handleLogout = async () => {
  await logOut()
  window.location.reload()
}
</script>
