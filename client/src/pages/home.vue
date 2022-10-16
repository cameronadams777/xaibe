<script lang="ts" setup>
import TheMainLayout from "../layouts/the-main-layout.vue";
import NoAlerts from "../components/no-alerts.vue";
import { useActiveUserStore } from "../state/active-user";
import { onMounted } from "vue";

const { activeUser, getActiveUser } = useActiveUserStore();

// Get all existing alerts in the redis queue
// Connect via websocket to the api and await new alerts

onMounted(async () => {
  await getActiveUser();
});
</script>

<template>
  <the-main-layout>
    <!-- <NoAlerts /> -->
  </the-main-layout>
</template>
