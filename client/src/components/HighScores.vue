<template>
    <div
        v-show="gameStoreRef.ShowHighscores.value"
        class="w-screen h-screen bg-red-500"
    >
        <button
            @click="handleCloseHighscores()"
            class="absolute text-white top-5 left-5"
        >X</button>
        highscore
    </div>
</template>

<script setup>
import { storeToRefs } from "pinia";
import { useGameStore } from "../stores/game.js";

const gameStore = useGameStore();
let gameStoreRef = storeToRefs(gameStore);

const { toggleHighscores } = gameStore;

const handleCloseHighscores = () => {
    toggleHighscores()
}
</script>