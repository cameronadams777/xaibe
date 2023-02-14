<template>
  <the-main-layout :is-loading="isLoading">
    <div class="p-4 w-full">
      <h2>Hello {{ activeUser?.firstName }}!</h2>
      <div class="flex flex-col lg:flex-row">
        <div class="lg:w-1/2 mr-2">
          <all-application-alerts-list :alerts="alerts" />
        </div>
        <div class="lg:w-1/2">
          <pending-team-invites-list />
        </div>
      </div>
    </div>
  </the-main-layout>
</template>

<script lang="ts" setup>
import { storeToRefs } from "pinia";
import { onMounted, ref } from "vue";
import { ToastType } from "src/types";
import { fetchAllCachedAlerts, ICachedAlerts } from "src/api/alerts";
import { useActiveUserStore, useTeamsStore, useToastStore } from "src/state";
import TheMainLayout from "src/layouts/the-main-layout.vue";
import PendingTeamInvitesList from "src/components/pending-team-invites-list.vue";
import AllApplicationAlertsList from "src/components/all-application-alerts-list.vue";

const { getPendingTeamInvites } = useTeamsStore();
const { setActiveToast } = useToastStore();
const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);

const isLoading = ref(true);
const alerts = ref<ICachedAlerts[]>([]);

onMounted(async () => {
  try {
    // TODO: Refactor this into global state
    const [groupedAlerts, _] = await Promise.all([fetchAllCachedAlerts(), getPendingTeamInvites()]);
    alerts.value = Object.values(groupedAlerts);
    isLoading.value = false;
  } catch (error) {
    console.error(error);
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred while fetching your alerts.",
    });
    isLoading.value = false;
  }
});
</script>
