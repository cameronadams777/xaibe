<template>
  <the-main-layout>
    <div class="w-full h-full p-4">
      <div class="flex justify-between items-center">
        <h2 class="capitalize">{{ activeTeam?.Name }}</h2>
        <base-fab-button
          aria-label="Delete Team"
          @click="setIsDeleteTeamConfirmationModalShown(true)"
        >
          <trash-icon class="w-5 h-5 text-red-500" />
        </base-fab-button>
      </div>
      <div class="flex">
        <div class="w-full h-48 lg:h-96 lg:w-1/4 mr-4 mb-4">
          <applications-list :applications="activeTeam?.Applications || []" />
        </div>
        <div class="w-full h-48 lg:h-96 lg:w-1/4 mr-4 mb-4">
          <team-member-list :members="activeTeam?.Users || []" />
        </div>
      </div>
    </div>
  </the-main-layout>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { TrashIcon } from "@heroicons/vue/24/outline";
import TheMainLayout from "src/layouts/the-main-layout.vue";
import ApplicationsList from "src/components/applications-list.vue";
import TeamMemberList from "src/components/team-member-list.vue";
import { fetchTeamById } from "src/api/teams";
import { useModalStore } from "src/state/modals";
import { ITeam, ToastType } from "src/types";
import { useToastStore } from "src/state";

const route = useRoute();
const router = useRouter();
const { setIsDeleteTeamConfirmationModalShown } = useModalStore();
const { setActiveToast } = useToastStore();

const activeTeam = ref<ITeam | undefined>(undefined);

onMounted(async () => {
  try {
    const teamId = parseInt(route.params.teamId as string);
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
    console.log(error);
    setActiveToast({
      message: "An error occurred trying to fetch the team you wanted.",
      type: ToastType.ERROR,
    });
    router.push("/");
  }
});
</script>
