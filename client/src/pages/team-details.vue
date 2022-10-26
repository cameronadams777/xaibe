<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { TrashIcon } from "@heroicons/vue/24/outline";
import TheMainLayout from "../layouts/the-main-layout.vue";
import ApplicationsList from "../components/applications-list.vue";
import { fetchTeamById } from "../api/teams";
import { useModalStore } from "../state/modals";
import { IApplication, ITeam } from "../types";

const route = useRoute();
const router = useRouter();
const { setIsDeleteTeamConfirmationModalShown } = useModalStore();

const activeTeam = ref<ITeam | undefined>(undefined);

onMounted(async () => {
  try {
    const teamId = parseInt(route.params.teamId as string);
    const team = await fetchTeamById({ teamId });
    if (!team) {
      router.push("/404");
      return;
    }
    activeTeam.value = team;
  } catch (error) {
    console.error(
      "Galata Error: An error occurred trying to fetch the team you wanted:",
      error
    );
    router.push("/");
  }
});
</script>

<template>
  <the-main-layout>
    <div class="w-full h-full p-4">
      <div class="flex justify-between items-center">
        <h2 class="capitalize">{{ activeTeam?.Name }}</h2>
        <button
          role="button"
          class="w-8 h-8 p-0 m-0 bg-white text-red-500 hover:shadow-md rounded-full border-none transition-all duration-500 cursor-pointer"
          aria-label="Delete Team"
          @click="setIsDeleteTeamConfirmationModalShown(true)"
        >
          <trash-icon class="w-5 h-5" />
        </button>
      </div>
      <div class="flex">
        <applications-list
          :applications="(activeTeam?.Applications as IApplication[])"
        />
      </div>
    </div>
  </the-main-layout>
</template>
