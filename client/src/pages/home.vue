<script lang="ts" setup>
import TheMainLayout from "../layouts/the-main-layout.vue";
import { useActiveUserStore } from "../state/active-user";
import { onMounted } from "vue";

const { getActiveUser } = useActiveUserStore();

onMounted(async () => {
  await getActiveUser();
});
</script>

<template>
  <the-main-layout>
    <!-- {{ alerts }} -->
  </the-main-layout>
</template>
