<script lang="ts" setup>
import { ref } from "vue";
import TheTopNav from "../components/the-top-nav.vue";
import TheLeftNav from "../components/the-left-nav.vue";

const isLeftNavOpen = ref(true);

const toggleLeftNav = () => (isLeftNavOpen.value = !isLeftNavOpen.value);
</script>

<template>
  <the-top-nav @on-menu-click="toggleLeftNav" />
  <div class="relative h-full flex">
    <the-left-nav :isOpen="isLeftNavOpen" />
    <slot />
  </div>
</template>
