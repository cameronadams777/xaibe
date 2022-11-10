<script lang="ts" setup>
import { onMounted, ref } from "vue";
import TheTopNav from "../components/the-top-nav.vue";
import TheLeftNav from "../components/the-left-nav.vue";
import { useActiveUserStore } from "src/state";

const { getActiveUser } = useActiveUserStore();

const isLeftNavOpen = ref(true);

const toggleLeftNav = () => (isLeftNavOpen.value = !isLeftNavOpen.value);

onMounted(async () => {
  await getActiveUser();
});
</script>

<template>
  <the-top-nav @on-menu-click="toggleLeftNav" />
  <div class="relative h-full flex">
    <the-left-nav :isOpen="isLeftNavOpen" />
    <slot />
  </div>
</template>
