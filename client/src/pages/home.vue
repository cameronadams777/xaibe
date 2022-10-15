<script lang="ts" setup>
import TheMainLayout from "../layouts/the-main-layout.vue";
import NoAlerts from "../components/no-alerts.vue";
import { useAlertsStore } from "../state/alerts";
import { onMounted } from "vue";

const { getAlerts } = useAlertsStore();

// Get all existing alerts in the redis queue
// Connect via websocket to the api and await new alerts

onMounted(async () => {
  await getAlerts({ applicationId: 0, serviceToken: "1234" });
});
</script>

<template>
  <the-main-layout>
    <!-- <NoAlerts /> -->
  </the-main-layout>
</template>
