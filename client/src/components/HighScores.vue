<template>
    <div
        v-show="gameStoreRef.ShowHighscores.value"
        class="fixed w-screen h-screen flex items-center justify-center p-8 "
    >
        <div class="bg-sub_alt_color text-text_color p-4 w-screen rounded-xl">
            <button
                @click="handleCloseHighscores()"
                class="mb-2 rounded-full h-8 w-8 font-bold self-start text-sm bg-bg_color text-text_color"
            >X</button>
            <p class="text-text_color text-center pb-4 font-bold">Highscores only contain players who've authenticated with
                Google
            </p>
            <table class="w-full text-center">
                <tr class="text-caret_color">
                    <th>Rank</th>
                    <th>Player</th>
                    <th>Wins</th>
                    <th>Losses</th>
                </tr>
                <tr
                    class="text-text_color"
                    v-for="score, idx in scores"
                    :class="idx + 1 == 1 ? 'font-bold' : ''"
                >
                    <td class="stroke-main_color flex justify-center items-center gap-x-2">
                        <svg
                            v-show="idx + 1 == 1"
                            width="20"
                            height="20"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="current"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="lucide lucide-crown"
                        >
                            <path d="m2 4 3 12h14l3-12-6 7-4-7-4 7-6-7zm3 16h14"></path>
                        </svg>
                        {{ idx + 1 }}
                    </td>
                    <td>{{ generateName(score.ID) }}</td>
                    <td>{{ score.wins }}</td>
                    <td>{{ score.losses }}</td>
                </tr>
            </table>
        </div>
    </div>
</template>

<script setup>
import { storeToRefs } from "pinia";
import { onMounted, ref } from "vue";
import { useGameStore } from "../stores/game.js";
import { generateName } from "../game/nameGenerator";

const gameStore = useGameStore();
let gameStoreRef = storeToRefs(gameStore);
const { toggleHighscores, getHighscores } = gameStore;

let scores = ref([])

const handleCloseHighscores = () => {
    toggleHighscores()
}

onMounted(async () => {
    scores = await getHighscores()
})
</script>