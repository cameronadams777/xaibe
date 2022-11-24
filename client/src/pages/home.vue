<script lang="ts" setup>
import TheMainLayout from "src/layouts/the-main-layout.vue";
import AllApplicationAlertsList from "src/components/all-application-alerts-list.vue";
import { useActiveUserStore, useToastStore } from "src/state";
import { storeToRefs } from "pinia";
import { onMounted, ref } from "vue";
import { IAlert, IAlertSchema, IApplication, ToastType } from "src/types";
import { fetchAllCachedAlerts, ICachedAlerts } from "src/api/alerts";

const { setActiveToast } = useToastStore();
const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);

const alerts = ref<ICachedAlerts[]>([]);

onMounted(async () => {
  try {
    const groupedAlerts = await fetchAllCachedAlerts();
    alerts.value = Object.values(groupedAlerts);
  } catch (error) {
    console.error(error);
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred while fetching your alerts.",
    });
  }
});
</script>

<template>
  <the-main-layout>
    <div class="p-4 w-full">
      <h2>Hello {{ activeUser?.FirstName }}!</h2>
      <all-application-alerts-list :alerts="alerts" />
    </div>
  </the-main-layout>
</template>
