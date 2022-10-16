<script lang="ts" setup>
import { ref } from "vue";
import { PlusIcon } from "@heroicons/vue/24/solid";

import TheTopNav from "../components/the-top-nav.vue";
import TheLeftNav from "../components/the-left-nav.vue";

const isLeftNavOpen = ref(true);

const toggleLeftNav = () => (isLeftNavOpen.value = !isLeftNavOpen.value);
</script>

<template>
  <the-top-nav @onMenuClick="toggleLeftNav" />
  <div class="relative h-full flex">
    <the-left-nav :isOpen="isLeftNavOpen" />
    <slot />
    <button
      class="absolute bottom-20 right-5 flex justify-center items-center w-10 h-10 bg-indigo-600 hover:bg-indigo-800 rounded-full border-none text-white cursor-pointer"
    >
      <plus-icon class="w-8" />
    </button>
  </div>
</template>
