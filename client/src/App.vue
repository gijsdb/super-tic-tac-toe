<template>
  <div class="w-screen min-h-screen bg-gradient flex flex-col ">
    <div class="flex justify-between p-2 ">
      <rules></rules>
      <player-info></player-info>
    </div>
    <router-view></router-view>
  </div>
</template>

<style></style>

<script setup>
import { RouterView } from "vue-router";
import PlayerInfo from "./components/PlayerInfo.vue";
import Rules from "./components/Rules.vue";
</script>
