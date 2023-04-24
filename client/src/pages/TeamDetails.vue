<template>
  <the-main-layout>
    <div v-if="activeTeam != null" class="w-full h-full p-4">
      <div class="flex justify-between items-center">
        <h2 class="capitalize">{{ activeTeam?.name }}</h2>
        <base-fab-button
          v-if="managerControlsVisible"
          aria-label="Delete Team"
          @click="setIsDeleteTeamConfirmationModalShown(true)"
        >
          <trash-icon class="w-5 h-5 text-red-500" />
        </base-fab-button>
      </div>
      <div class="flex">
        <div class="w-full h-48 lg:h-96 lg:w-1/4 mr-4 mb-4">
          <applications-list :applications="activeTeam?.applications || []" />
        </div>
        <div class="w-full h-48 lg:h-96 lg:w-1/4 mr-4 mb-4">
          <team-member-list
            :team-id="activeTeam.id"
            :members="activeTeam?.users || []"
            :manager-controls-visible="managerControlsVisible"
          />
        </div>
      </div>
    </div>
  </the-main-layout>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { TrashIcon } from "@heroicons/vue/24/outline";
import TheMainLayout from "src/layouts/TheMainLayout.vue";
import ApplicationsList from "src/components/ApplicationsList.vue";
import TeamMemberList from "src/components/TeamMemberList.vue";
import { fetchTeamById } from "src/api/teams";
import { useModalStore } from "src/state/modals";
import { Team, ToastType } from "src/types";
import { useActiveUserStore, useToastStore } from "src/state";
import { storeToRefs } from "pinia";

const route = useRoute();
const router = useRouter();
const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);
const { setIsDeleteTeamConfirmationModalShown } = useModalStore();
const { setActiveToast } = useToastStore();

const managerControlsVisible = computed(
  () =>
    activeTeam.value?.managers.some(
      (manager) => manager.id === activeUser?.value?.id
    ) || false
);

const activeTeam = ref<Team | undefined>(undefined);

onMounted(async () => {
  try {
    const teamId = route.params.teamId as string;
    const team = await fetchTeamById({ teamId });
    if (!team) {
      setActiveToast({
        message: "The team you are trying to view does not exist.",
        type: ToastType.ERROR,
      });
      router.push("/404");
      return;
    }
    activeTeam.value = team;
  } catch (error) {
    setActiveToast({
      message: "An error occurred trying to fetch the team you wanted.",
      type: ToastType.ERROR,
    });
    router.push("/");
  }
});
</script>
