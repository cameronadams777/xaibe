<script lang="ts" setup>
import { Bars3Icon } from "@heroicons/vue/24/solid";
import { UserIcon } from "@heroicons/vue/24/outline";

const emits = defineEmits<{ (event: "onMenuClick"): void }>();
</script>

<template>
  <div class="p-4 border-b border-gray-300 flex justify-between">
    <div class="flex justify-center">
      <button
        class="bg-white border-none mr-2 cursor-pointer hover:bg-gray-200 rounded-xl"
        @click="emits('onMenuClick')"
      >
        <bars-3-icon class="h-full w-6" />
      </button>
      <h2 class="m-0 p-0">Galata</h2>
    </div>
    <div class="flex justify-end">
      <button
        class="w-8 h-8 bg-gray-800 flex justify-center items-center text-white border-none rounded-full cursor-pointer hover:shadow-md"
      >
        <user-icon class="w-4" />
      </button>
    </div>
  </div>
</template>
