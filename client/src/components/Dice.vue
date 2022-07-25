<template>
  <div
    class="bg-black bg-opacity-60 border-white border-4 p-8 mx-8 rounded-2xl"
  >
    <button
      :disabled="rolled"
      @click="rollDice"
      class="
        bg-green-500
        rounded-md
        p-4
        text-white
        font-bold
        disabled:bg-gray-500
      "
    >
      Roll dice
    </button>
    <ul class="flex text-white text-6xl" id="rolls"></ul>
  </div>
</template>

<script setup>
import { onMounted, ref, watch } from "vue";

let rolled = ref(false);

const props = defineProps({
  clear: {
    type: Boolean,
    default: false,
  },
});

const rollDice = () => {
  rolled.value = true;
  createDice();
  createDice();
};

const getRandomValue = () => {
  return Math.floor(Math.random() * 6);
};

const createDice = () => {
  const roll = document.createElement("li");
  roll.className = "roll";
  roll.innerHTML = `&#x268${getRandomValue()};`;
  roll.innerHTML = `&#x268${getRandomValue()};`;
  document.getElementById("rolls").appendChild(roll);
};

const clearDice = () => {
  const dice = document.getElementsByClassName("roll");
  dice[0].remove();
  dice[0].remove();
};

watch(
  () => props.clear,
  () => {
    if (props.clear) {
      rolled.value = true;
      clearDice();
    } else {
      rolled.value = false;
    }
  },
  { immediate: true }
);
</script>