<template>
  <div class="w-full h-full flex">
    <div class="w-2/3 h-full">
      <h5 class="p-0 m-0 mb-1">Alerts</h5>
      <div
        class="w-full h-full flex flex-col mr-4 border-1 border-gray-300 rounded-lg overflow-y-auto"
      >
        <div
          v-if="applicationAlertsMappedToSchema?.length"
          v-for="(alert, index) of applicationAlertsMappedToSchema"
          :key="index"
          class="p-4 flex justify-between items-center border-b border-gray-300"
        >
          <div>
            <span class="font-bold">{{ alert.title }}</span>
            <p class="m-0">{{ alert.description }}</p>
          </div>
          <a
            :href="alert.link"
            target="_blank"
            class="text-gray-400 hover:text-gray-800"
          >
            <ArrowTopRightOnSquareIcon class="w-6 h-6" />
          </a>
        </div>
        <div v-else class="w-full h-full flex justify-center items-center">
          <p class="w-9/10 text-center">
            There are no recent alerts to review. Remember, Xaibe will keep
            track of an applications most recent alerts for up to two weeks
            after it is received. Afterwards, the alert will be lost.
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, toRaw } from "vue";
import { ArrowTopRightOnSquareIcon } from "@heroicons/vue/24/outline";
import { IAlert, AlertSchema } from "src/types";
import { getElByKey } from "src/helpers";

const props = defineProps<{ alertSchema: AlertSchema; alerts: IAlert[] }>();

const applicationAlertsMappedToSchema = computed(() =>
  props.alerts
    .map((alert) => {
      const titleKeys = props.alertSchema.title.split(".");
      const title = getElByKey(toRaw(alert), titleKeys);
      const descriptionKeys = props.alertSchema.description.split(".");
      const description = getElByKey(toRaw(alert), descriptionKeys);
      const linkKeys = props.alertSchema.link.split(".");
      const link = getElByKey(toRaw(alert), linkKeys);

      return { title, description, link } as IAlert;
    })
    .filter((alert) => Object.values(alert).every((value) => !!value))
);
</script>
