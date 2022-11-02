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
          v-if="alerts.length"
          v-for="(alert, index) of alerts"
          :key="index"
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
          >
            <arrow-top-right-on-square-icon class="w-6 h-6" />
          </a>
        </div>
        <div class="w-full h-full flex justify-center items-center">
          <p>
            There are no recent alerts to review. Remember, Galata will keep
            track of an applications most recent alerts for up to two weeks
            after it is received. Afterwards, the alert will be lost.
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
