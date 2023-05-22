<template>
    <div
        v-if="colorStoreRef.Show.value"
        class="fixed w-screen h-screen flex items-center justify-center shadow-2xl"
    >
        <div class="p-2 rounded-xl w-screen sm:w-[80vw] flex flex-col bg-sub_alt_color">
            <button
                @click="handleToggleThemePicker"
                class="mb-2 rounded-full h-8 w-8 font-bold self-start text-sm bg-bg_color text-text_color"
            >X</button>

            <ul class="flex flex-col gap-y-8">
                <li v-for="color in colorStoreRef.Themes.value ">
                    <button
                        @click="handleChangeTheme(color.name)"
                        class="flex w-full gap-y-2 justify-between items-center p-2 "
                    >
                        <div class="flex gap-x-2 text-main_color">
                            <svg
                                v-if="colorStoreRef.ActiveTheme.value.name === color.name"
                                width="24"
                                height="24"
                                viewBox="0 0 24 24"
                                fill="none"
                                stroke="currentColor"
                                stroke-width="2"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                class="lucide lucide-check"
                            >
                                <polyline points="20 6 9 17 4 12"></polyline>
                            </svg>
                            <p class="text-sm text-text_color">{{ color.name }}</p>
                        </div>
                        <div
                            class="flex gap-x-4 px-2 py-1 rounded-3xl"
                            :style="{ backgroundColor: color.bg_color }"
                        >
                            <div
                                class="w-4 h-4 rounded-full"
                                :style="{ backgroundColor: color.main_color }"
                            ></div>
                            <div
                                class="w-4 h-4 rounded-full"
                                :style="{ backgroundColor: color.caret_color }"
                            ></div>
                            <div
                                class="w-4 h-4 rounded-full"
                                :style="{ backgroundColor: color.sub_color }"
                            ></div>
                        </div>
                    </button>
                </li>
            </ul>
        </div>

    </div>
</template>

<script setup>
import { useColorStore } from "../stores/color.js";
import { storeToRefs } from "pinia";

const colorStore = useColorStore();
let colorStoreRef = storeToRefs(colorStore);
const { changeTheme, toggleThemePicker } = colorStore;

const handleChangeTheme = (theme) => {
    changeTheme(theme)
}

const handleToggleThemePicker = () => {
    toggleThemePicker()
}
</script>