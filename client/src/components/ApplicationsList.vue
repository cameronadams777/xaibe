<template>
  <div class="w-full h-full flex flex-col">
    <h5 class="p-0 m-0 ml-1 mb-2">Applications</h5>
    <div
      class="w-full h-full flex flex-col mr-4 border-1 border-gray-300 rounded-lg overflow-y-auto"
    >
      <router-link
        v-if="applications.length"
        v-for="application in applications"
        :key="application.id"
        :to="`/applications/${application.id}`"
        class="p-4 flex justify-between items-center border-b border-gray-300 hover:bg-gray-600 hover:border-gray-600 hover:text-white font-bold text-black no-underline transition-all duration-300"
      >
        {{ application.name }}
        <chevron-right-icon class="h-full w-4" />
      </router-link>
      <div v-else class="w-full h-full flex justify-center items-center">
        <p>No applications</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ChevronRightIcon } from "@heroicons/vue/24/solid";
import { Application } from "src/types";

defineProps<{ applications: Application[] }>();
</script>
