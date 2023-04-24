<template>
  <div class="p-4 bg-white border-b border-gray-300 flex justify-between">
    <div class="flex justify-center">
      <BaseFabButton @click="emits('onMenuClick')">
        <Bars3Icon class="h-full w-6" />
      </BaseFabButton>
    </div>
    <div class="flex justify-end">
      <BaseFabButton @click="setIsNewElementModalShown(true)">
        <PlusIcon class="w-8" />
      </BaseFabButton>
      <TheTopNavMenu />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Bars3Icon } from "@heroicons/vue/24/solid";
import { PlusIcon } from "@heroicons/vue/24/outline";
import TheTopNavMenu from "./TheTopNavMenu.vue";
import { useModalStore } from "../state/modals";

const emits = defineEmits<{ (event: "onMenuClick"): void }>();

const { setIsNewElementModalShown } = useModalStore();
</script>

