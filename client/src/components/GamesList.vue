<template>
  <div class="w-screen h-screen absolute bg-black bg-opacity-50">
    <div class="bg-red-500">
      <div :key="game.id" v-for="game in games">
        <p>Game: {{ game.id }}</p>
        <router-link :to="'/game/' + game.id"
          ><button>Join</button></router-link
        >
      </div>
    </div>
  </div>
</template>


<script setup>
const props = defineProps({
  games: Array,
});

console.log("LIST: ", props.games);
</script>