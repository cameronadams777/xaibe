<script lang="ts" setup>
import { ChevronRightIcon, ChevronDownIcon } from "@heroicons/vue/24/solid";
import { storeToRefs } from "pinia";
import { useActiveUserStore } from "../state/active-user";
import { ITeam } from "../types";

defineProps<{ isOpen: boolean }>();

const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);
</script>

<template>
  <div
    class="h-full hidden border-r-1 border-gray-300"
    :class="{ 'flex flex-col w-full md:w-1/4 lg:w-1/5 xl:w-1/8': isOpen }"
  >
    <h3 class="text-gray-700 py-4 pl-6 m-0">My Applications</h3>
    <router-link
      v-if="activeUser?.applications?.length"
      v-for="(application, index) in activeUser.applications"
      class="px-6 py-4 flex items-center justify-between font-bold text-md capitalize no-underline bg-white hover:bg-gray-200 text-gray-800 text-left border-x-0 border-t-1 border-b-1 border-gray-300 cursor-pointer"
      :class="{ 'border-t-0': index === 0 }"
      to="/applications/:applicationId"
    >
      <span>{{ application.name }}</span>
      <chevron-right-icon class="h-full w-4" />
    </router-link>
    <div
      v-else
      class="pb-4 flex justify-center items-center border-b-1 border-gray-300"
    >
      <span class="text-gray-700">No Applications</span>
    </div>
    <h3 class="text-gray-700 py-4 pl-6 m-0">Teams</h3>
    <button
      v-if="[].length"
      v-for="(team, index) in ([] as ITeam[])"
      class="px-6 py-4 flex items-center bg-white hover:bg-gray-200 text-left border-x-0 border-t-1 border-b-1 border-gray-300 cursor-pointer"
      :class="{ 'border-t-0': index === 0 }"
    >
      <ChevronDownIcon class="h-full w-4 mr-2" />
      <span class="font-bold">{{ team.name }}</span>
    </button>
    <div
      v-else
      class="pb-4 flex justify-center items-center border-b-1 border-gray-300"
    >
      <span class="text-gray-700">No Teams</span>
    </div>
  </div>
</template>
