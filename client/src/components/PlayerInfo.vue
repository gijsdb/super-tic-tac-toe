<template>
  <div
    :style="{ backgroundColor: colorStoreRef.ActiveTheme.value.Secondary }"
    class="rounded-xl p-2"
  >
    <div class="flex justify-end items-center">
      <div class="flex flex-col mr-4 text-sm gap-y-2">
        <p>
          <span
            :style="{ color: colorStoreRef.ActiveTheme.value.Highlight }"
            class="font-bold mr-2 text-lg"
          >{{
            generateName(gameStoreRef.Player.value.id) }} </span>
        </p>
        <div class="flex items-center gap-x-2">
          <rules />
          <a
            v-show="gameStoreRef.Player.value.temporary"
            href="http://localhost:1323/login"
          >
            <svg
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              class="lucide lucide-log-in"
              :style="{ stroke: colorStoreRef.ActiveTheme.value.Highlight }"
            >
              <path d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4"></path>
              <polyline points="10 17 15 12 10 7"></polyline>
              <line
                x1="15"
                x2="3"
                y1="12"
                y2="12"
              ></line>
            </svg></a>
          <button
            v-show="!gameStoreRef.Player.value.temporary"
            @click="handleLogout()"
          >
            <svg
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              class="lucide lucide-log-out"
              :style="{ stroke: colorStoreRef.ActiveTheme.value.HighlightTwo }"
            >
              <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path>
              <polyline points="16 17 21 12 16 7"></polyline>
              <line
                x1="21"
                x2="9"
                y1="12"
                y2="12"
              ></line>
            </svg></button>
          <button @click="handleToggleThemePicker()">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              :style="{ stroke: colorStoreRef.ActiveTheme.value.Highlight }"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              class="lucide lucide-palette"
            >
              <circle
                cx="13.5"
                cy="6.5"
                r=".5"
              ></circle>
              <circle
                cx="17.5"
                cy="10.5"
                r=".5"
              ></circle>
              <circle
                cx="8.5"
                cy="7.5"
                r=".5"
              ></circle>
              <circle
                cx="6.5"
                cy="12.5"
                r=".5"
              ></circle>
              <path
                d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10c.926 0 1.648-.746 1.648-1.688 0-.437-.18-.835-.437-1.125-.29-.289-.438-.652-.438-1.125a1.64 1.64 0 0 1 1.668-1.668h1.996c3.051 0 5.555-2.503 5.555-5.554C21.965 6.012 17.461 2 12 2z"
              >
              </path>
            </svg>
          </button>
        </div>

      </div>
      <div class="flex flex-col items-center gap-y-1">
        <img
          v-if="gameStoreRef.Player.value.picture == null"
          class="rounded-full w-14 h-14 border-4"
          :style="{ borderColor: colorStoreRef.ActiveTheme.value.Highlight }"
          src="../assets/avatar.png"
        />
        <img
          v-if="gameStoreRef.Player.value.picture != null"
          class="rounded-full w-14 h-14 border-4"
          :style="{ borderColor: colorStoreRef.ActiveTheme.value.Highlight }"
          :src="gameStoreRef.Player.value.picture"
        />
        <p
          :style="{ color: colorStoreRef.ActiveTheme.value.Highlight }"
          class=" font-bold text-sm"
        >
          <span v-if="gameStoreRef.Player.value.inGame">In Game</span>
          <span v-if="!gameStoreRef.Player.value.inGame">Idle</span>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import { useGameStore } from "../stores/game.js";
import { useColorStore } from "../stores/color.js";
import { storeToRefs } from "pinia";
import { generateName } from "../game/nameGenerator.js";
import Rules from "../components/Rules.vue";
import ThemePicker from "./ThemePicker.vue";

const router = useRouter();

const gameStore = useGameStore();
const colorStore = useColorStore();

const { logOut } = gameStore;
const { toggleThemePicker } = colorStore;

let gameStoreRef = storeToRefs(gameStore);
let colorStoreRef = storeToRefs(colorStore);

const handleToggleThemePicker = async () => {
  toggleThemePicker()
}

const handleLogout = async () => {
  await logOut()
  window.location.reload()
}
</script>
