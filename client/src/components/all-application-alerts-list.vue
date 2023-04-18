<template>
  <div class="w-full h-full">
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
          <arrow-top-right-on-square-icon class="w-6 h-6" />
        </a>
      </div>
      <div v-else class="w-full h-full flex justify-center items-center">
        <p class="w-9/10 text-center">
          There are no recent alerts to review. Remember, alata will keep track
          of an applications most recent alerts for up to two weeks after it is
          received. Afterwards, the alert will be lost.
        </p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import flatten from "lodash/flatten";
import compact from "lodash/compact";
import { ArrowTopRightOnSquareIcon } from "@heroicons/vue/24/outline";
import { IAlert, AlertSchema } from "src/types";
import { getElByKey } from "src/helpers";

const props = defineProps<{
  alerts: { AlertSchema: AlertSchema; Alerts: IAlert[] }[];
}>();

// TODO: Refactor to be more performant
let applicationAlertsMappedToSchema = computed(() => {
  return flatten(
    props.alerts.map((alert) => {
      if (!alert.Alerts) return [];
      const alertsByApp = compact(alert.Alerts).map((a) => {
        const titleKeys = alert.AlertSchema.title.split(".");
        const title = getElByKey(a, titleKeys);
        const descriptionKeys = alert.AlertSchema.description.split(".");
        const description = getElByKey(a, descriptionKeys);
        const linkKeys = alert.AlertSchema.link.split(".");
        const link = getElByKey(a, linkKeys);

        return { title, description, link } as IAlert;
      });
      return alertsByApp;
    })
  ).filter((alert) => Object.values(alert).every((value) => !!value));
});
</script>
