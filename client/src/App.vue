<template>
  <div class="w-screen h-screen">
    <player-info></player-info>
    <rules></rules>
    <router-view></router-view>
  </div>
</template>

<style></style>

<script setup>
import { RouterView } from "vue-router";
import PlayerInfo from "./components/PlayerInfo.vue";
import Rules from "./components/Rules.vue";
</script>
