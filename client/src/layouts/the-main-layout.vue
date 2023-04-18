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
import { useRouter } from "vue-router";
import TheTopNav from "src/components/the-top-nav.vue";
import TheLeftNav from "src/components/the-left-nav.vue";
import TheSpinner from "src/components/the-spinner.vue";
import { useActiveUserStore } from "src/state";

defineProps<{ isLoading?: boolean }>();

const router = useRouter();
const { getActiveUser } = useActiveUserStore();

const isLeftNavOpen = ref(true);

const toggleLeftNav = () => (isLeftNavOpen.value = !isLeftNavOpen.value);

onMounted(async () => {
  try {
    await getActiveUser();
  } catch (error) {
    console.error(error);
    router.push("/500");
  }
});
</script>
