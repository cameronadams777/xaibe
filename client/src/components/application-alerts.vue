<script lang="ts" setup>
import { ArrowTopRightOnSquareIcon } from "@heroicons/vue/24/outline";
import { IAlert } from "../types";

defineProps<{ alerts: IAlert[] }>();
</script>

<template>
  <div class="w-full h-full flex">
    <div class="w-2/3 h-full">
      <h5 class="p-0 m-0 mb-1">Alerts</h5>
      <div
        class="w-full h-full flex flex-col mr-4 border-1 border-gray-300 rounded-lg overflow-y-auto"
      >
        <div
          v-for="alert in alerts"
          :key="alert.ID"
          class="p-4 flex justify-between items-center border-b border-gray-300"
        >
          <div>
            <span class="font-bold">{{ alert.Title }}</span>
            <p class="m-0">{{ alert.Description }}</p>
          </div>
          <a
            :href="alert.Link"
            target="_blank"
            class="text-gray-400 hover:text-gray-800"
            ><arrow-top-right-on-square-icon class="w-6 h-6"
          /></a>
        </div>
      </div>
    </div>
  </div>
</template>
