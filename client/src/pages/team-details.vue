<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { fetchTeamById } from "../api/teams";
import TheMainLayout from "../layouts/the-main-layout.vue";
import { ITeam } from "../types";

const route = useRoute();
const router = useRouter();
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
      </div>
    </div>
  </the-main-layout>
</template>
