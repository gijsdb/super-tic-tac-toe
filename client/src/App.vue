<template>
  <div
    :style="{ backgroundColor: colorStoreRef.ActiveTheme.value.Primary }"
    class="w-screen min-h-screen flex flex-col"
  >
    <theme-picker />
    <div class="p-5 flex self-end">
      <player-info></player-info>
    </div>
    <router-view></router-view>
  </div>
</template>

<style></style>

<script setup>
import { storeToRefs } from "pinia";
import { RouterView } from "vue-router";
import PlayerInfo from "./components/PlayerInfo.vue";
import ThemePicker from "./components/ThemePicker.vue";

import { useColorStore } from "./stores/color.js";
const colorStore = useColorStore();
let colorStoreRef = storeToRefs(colorStore);
</script>
