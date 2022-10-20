<script lang="ts" setup>
import TheMainLayout from "../layouts/the-main-layout.vue";
import NoAlerts from "../components/no-alerts.vue";
import { useAlertsStore } from "../state/alerts";
import { useActiveUserStore } from "../state/active-user";
import { onMounted } from "vue";
import { storeToRefs } from "pinia";

const { getActiveUser, ...activeUserStore } = useActiveUserStore();
const { getAlerts, ...alertsStore } = useAlertsStore();

const { activeUser } = storeToRefs(activeUserStore);
const { alerts } = storeToRefs(alertsStore);

// Get all existing alerts in the redis queue
// Connect via websocket to the api and await new alerts

onMounted(async () => {
  await getActiveUser();
  await getAlerts({ applicationId: 1, serviceToken: "" });
});
</script>

<template>
  <the-main-layout>
    <!-- {{ alerts }} -->
  </the-main-layout>
</template>
