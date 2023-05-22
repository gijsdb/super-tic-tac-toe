<template>
  <div class="w-screen min-h-screen flex flex-col bg-bg_color">
    <theme-picker />
    <rules />
    <div class="p-5 flex self-end">
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
import ThemePicker from "./components/ThemePicker.vue";


</script>
