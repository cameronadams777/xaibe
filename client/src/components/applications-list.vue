<script lang="ts" setup>
import { IApplication } from "../types";

defineProps<{ applications: IApplication[] }>();
</script>

<template>
  <div class="w-full h-full flex">
    <div class="w-2/3 h-full">
      <h5 class="p-0 m-0 mb-1">Applications</h5>
      <div
        class="w-full h-full flex flex-col mr-4 border-1 border-gray-300 rounded-lg overflow-y-auto"
      >
        <router-link
          v-if="applications.length"
          v-for="application in applications"
          :key="application.ID"
          :to="`/applications/${application.ID}`"
          class="p-4 flex justify-between items-center border-b border-gray-300"
        >
          {{ application.Name }}
        </router-link>
        <div v-else class="w-full h-full flex justify-center items-center">
          <p>No applications</p>
        </div>
      </div>
    </div>
  </div>
</template>
