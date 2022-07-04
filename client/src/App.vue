
<template>
  <div class="w-screen h-screen">
    <!-- <Board /> -->
    <router-view></router-view>
  </div>
</template>

<style>
</style>

<script setup>
import { RouterView } from "vue-router";
import Board from "./views/Board.vue";
</script>