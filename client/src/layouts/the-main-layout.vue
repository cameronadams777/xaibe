<template>
  <the-top-nav @on-menu-click="toggleLeftNav" />
  <div class="relative h-full flex">
    <the-left-nav :isOpen="isLeftNavOpen" />
    <slot v-if="!isLoading" />
    <the-spinner v-else />
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import TheTopNav from "src/components/the-top-nav.vue";
import TheLeftNav from "src/components/the-left-nav.vue";
import TheSpinner from "src/components/the-spinner.vue";
import { useActiveUserStore } from "src/state";

defineProps<{ isLoading?: boolean }>();

const { getActiveUser } = useActiveUserStore();

const isLeftNavOpen = ref(true);

const toggleLeftNav = () => (isLeftNavOpen.value = !isLeftNavOpen.value);

onMounted(async () => {
  await getActiveUser();
});
</script>
