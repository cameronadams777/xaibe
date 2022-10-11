<script lang="ts" setup>
import { ChevronDownIcon } from "@heroicons/vue/24/solid";

defineProps<{ isOpen: boolean }>();
</script>

<template>
  <div
    class="h-full hidden border-r-1 border-gray-300"
    :class="{ 'flex flex-col w-1/8': isOpen }"
  >
    <h3 class="text-gray-700 py-4 pl-6 m-0">Teams</h3>
    <button
      class="px-6 py-4 flex items-center bg-white hover:bg-gray-200 text-left border-x-0 border-t-1 border-b-1 border-gray-300 cursor-pointer"
    >
      <ChevronDownIcon class="h-full w-4 mr-2" />
      <span class="font-bold">Marketing</span>
    </button>
    <button
      class="px-6 py-4 flex items-center bg-white hover:bg-gray-200 text-left border-x-0 border-t-0 border-b-1 border-gray-300 cursor-pointer"
    >
      <ChevronDownIcon class="h-full w-4 mr-2" />
      <span class="font-bold">Web</span>
    </button>
    <button
      class="px-6 py-4 flex items-center bg-white hover:bg-gray-200 text-left border-x-0 border-t-0 border-b-1 border-gray-300 cursor-pointer"
    >
      <ChevronDownIcon class="h-full w-4 mr-2" />
      <span class="font-bold">Mobile</span>
    </button>
  </div>
</template>
