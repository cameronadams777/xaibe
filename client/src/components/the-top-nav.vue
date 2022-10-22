<script lang="ts" setup>
import { Bars3Icon } from "@heroicons/vue/24/solid";
import { UserIcon, PlusIcon } from "@heroicons/vue/24/outline";
import { useModalStore } from "../state/modals";

const emits = defineEmits<{ (event: "onMenuClick"): void }>();

const { setIsNewElementModalShown } = useModalStore();
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
    </div>
    <div class="flex justify-end">
      <button
        class="w-8 h-8 mr-4 bg-white text-gray-800 flex justify-center items-center text-white border-none rounded-full cursor-pointer hover:shadow-md transition-all duration-500"
        @click="setIsNewElementModalShown(true)"
      >
        <plus-icon class="w-8" />
      </button>
      <button
        class="w-8 h-8 bg-gray-800 flex justify-center items-center text-white border-none rounded-full cursor-pointer hover:shadow-md transition-all duration-500"
      >
        <user-icon class="w-4" />
      </button>
    </div>
  </div>
</template>
