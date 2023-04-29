<template>
  <div class="w-full md:w-7/8 lg:w-3/4 flex justify-center items-center">
    <button class="w-8 h-8 rounded-l border-none text-white font-bold bg-indigo-500 hover:bg-indigo-700 cursor-pointer" @click="decrement">-</button>
    <input type="number" :value="numToUpdate" class="w-full m-0 text-lg text-center border-gray-200 border-t border-b border-l-none border-r-none" />
    <button class="w-8 h-8 rounded-r border-none text-white font-bold bg-indigo-500 hover:bg-indigo-700 cursor-pointer" @click="increment">+</button>
  </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue';


const props = defineProps<{ value: number }>();

const { value: numToUpdate } = toRefs(props);

const emits = defineEmits<{(event: "onUpdate", value: number): void}>();

const increment = () => {
  const newVal = numToUpdate.value + 1;
  emits("onUpdate", newVal);
};

const decrement = () => {
  const newVal = numToUpdate.value - 1;
  emits("onUpdate", newVal);
};

</script>

<style scoped>
/* Chrome, Safari, Edge, Opera */
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
</style>
