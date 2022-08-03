<template>
  <div class="w-screen h-screen">
    <player-info></player-info>
    <router-view></router-view>
  </div>
</template>

<style></style>

<script setup>
import { RouterView } from "vue-router";
import Board from "./views/Board.vue";
import PlayerInfo from "./components/PlayerInfo.vue";
</script>
