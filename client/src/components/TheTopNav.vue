<script lang="ts" setup>
import { Bars3Icon } from "@heroicons/vue/24/solid";
import { PlusIcon } from "@heroicons/vue/24/outline";
import TheTopNavMenu from "./the-top-nav-menu.vue";
import { useModalStore } from "../state/modals";

const emits = defineEmits<{ (event: "onMenuClick"): void }>();

const { setIsNewElementModalShown } = useModalStore();
</script>

<template>
  <div class="p-4 bg-white border-b border-gray-300 flex justify-between">
    <div class="flex justify-center">
      <base-fab-button @click="emits('onMenuClick')">
        <bars-3-icon class="h-full w-6" />
      </base-fab-button>
    </div>
    <div class="flex justify-end">
      <base-fab-button @click="setIsNewElementModalShown(true)">
        <plus-icon class="w-8" />
      </base-fab-button>
      <the-top-nav-menu />
    </div>
  </div>
</template>
