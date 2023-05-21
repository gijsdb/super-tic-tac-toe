<template>
    <div
        v-if="colorStoreRef.Show.value"
        class="fixed w-screen h-screen flex items-center justify-center "
    >
        <div
            class="px-28 py-8 rounded-xl w-[80vw]"
            :style="{ backgroundColor: colorStoreRef.ActiveTheme.value.Secondary }"
        >
            <button
                @click="handleToggleThemePicker"
                :style="{ backgroundColor: colorStoreRef.ActiveTheme.value.HighlightTwo, color: colorStoreRef.ActiveTheme.value.Primary }"
                class="flex mx-auto mb-4 px-2 py-1 rounded-md font-bold"
            >Close</button>

            <ul class="flex flex-col gap-y-8">
                <li v-for="color in  colorStoreRef.Themes.value ">
                    <button
                        @click="handleChangeTheme(color.Name)"
                        class="flex w-full gap-y-2 justify-between items-center p-2"
                    >
                        <div class="flex gap-x-2">
                            <svg
                                v-if="colorStoreRef.ActiveTheme.value.Name === color.Name"
                                :style="{ stroke: colorStoreRef.ActiveTheme.value.Highlight }"
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
                            <p :style="{ color: colorStoreRef.ActiveTheme.value.Primary }">{{ color.Name }}</p>
                        </div>
                        <div
                            class="flex gap-x-4 p-2 rounded-3xl"
                            :style="{ backgroundColor: color.Primary }"
                        >
                            <div
                                class="w-6 h-6 rounded-full"
                                :style="{ backgroundColor: color.Secondary }"
                            ></div>
                            <div
                                class="w-6 h-6 rounded-full"
                                :style="{ backgroundColor: color.Highlight }"
                            ></div>
                            <div
                                class="w-6 h-6 rounded-full"
                                :style="{ backgroundColor: color.HighlightTwo }"
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