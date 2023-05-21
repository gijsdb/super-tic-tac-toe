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
                class="bg-red-500 flex mx-auto mb-4 px-2 py-1"
            >Close</button>

            <ul class="flex flex-col gap-y-8">
                <li v-for="color in  colorStoreRef.Themes.value ">
                    <button
                        :class="{ 'border-2 rounded-3xl': color.Name === colorStoreRef.ActiveTheme.value.Name }"
                        :style="{ borderColor: colorStoreRef.ActiveTheme.HighlightTwo }"
                        @click="handleChangeTheme(color.Name)"
                        class="flex w-full gap-y-2 justify-between items-center p-2"
                    >

                        <p :style="{ color: colorStoreRef.ActiveTheme.value.Primary }">{{ color.Name }}</p>
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