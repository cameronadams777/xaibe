<script lang="ts" setup>
import { ChevronRightIcon } from "@heroicons/vue/24/solid";
import { storeToRefs } from "pinia";
import { useActiveUserStore } from "../state/active-user";

defineProps<{ isOpen: boolean }>();

const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);
</script>

<template>
  <div
    class="h-full hidden border-r-1 border-gray-300"
    :class="{ 'flex flex-col w-full md:w-2/5 xl:w-1/5': isOpen }"
  >
    <router-link
      class="px-6 py-4 flex items-center justify-between font-bold text-md capitalize no-underline bg-white hover:bg-gray-200 text-gray-800 text-left border-x-0 border-t-0 border-b-1 border-gray-300 cursor-pointer"
      to="/"
    >
      <span>All Alerts</span>
      <chevron-right-icon class="h-full w-4" />
    </router-link>
    <h3 class="text-gray-700 py-4 pl-6 m-0">My Applications</h3>
    <router-link
      v-if="activeUser?.Applications?.length"
      v-for="(application, index) in activeUser.Applications"
      class="px-6 py-4 flex items-center justify-between font-bold text-md capitalize no-underline bg-white hover:bg-gray-200 text-gray-800 text-left border-x-0 border-b-1 border-gray-300 cursor-pointer"
      :class="{ 'border-t-0': index !== 0, 'border-t-1': index === 0 }"
      :to="`/applications/${application.ID}`"
    >
      <span>{{ application.Name }}</span>
      <chevron-right-icon class="h-full w-4" />
    </router-link>
    <div
      v-else
      class="pb-4 flex justify-center items-center border-b-1 border-gray-300"
    >
      <span class="text-gray-700">No Applications</span>
    </div>
    <h3 class="text-gray-700 py-4 pl-6 m-0">Teams</h3>
    <router-link
      v-if="activeUser?.Teams?.length"
      v-for="(team, index) in activeUser.Teams"
      class="px-6 py-4 flex items-center justify-between font-bold text-md capitalize no-underline bg-white hover:bg-gray-200 text-gray-800 text-left border-x-0 border-b-1 border-gray-300 cursor-pointer"
      :class="{ 'border-t-0': index !== 0, 'border-t-1': index === 0 }"
      :to="`/teams/${team.ID}`"
    >
      <span>{{ team.Name }}</span>
      <chevron-right-icon class="h-full w-4" />
    </router-link>
    <div
      v-else
      class="pb-4 flex justify-center items-center border-b-1 border-gray-300"
    >
      <span class="text-gray-700">No Teams</span>
    </div>
  </div>
</template>
