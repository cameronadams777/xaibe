<template>
  <div class="w-full flex flex-col">
    <div class="mb-4 w-full flex flex-wrap justify-center items-center">
      <button
        v-for="application of applicationTileData"
        :key="application.id"
        class="flex flex-col justify-center items-center lg:w-48 h-48 mr-2 mb-2 rounded-lg cursor-pointer transition-all"
        :class="{
          'bg-indigo-600 text-white border-1 border-indigo-500':
            selectedApplication === application.id,
          'bg-white border-1 border-gray-400 hover:border-indigo-600':
            selectedApplication !== application.id,
        }"
        @click="selectedApplication = application.id"
      >
        <img :src="application.image" class="w-12 h-12" />
        <span class="text-lg font-bold">{{ application.name }}</span>
      </button>
    </div>
    <base-button
      text="Continue"
      :text-size="ButtonTextSize.LARGE"
      aria-label="proceed to next form step"
      :disabled="!selectedApplication"
      :aria-disabled="!selectedApplication"
      class="w-1/2 lg:w-1/4 self-center"
      @click="onSubmit"
    />
  </div>
</template>

<script lang="ts" setup>
import { getSrc } from "src/helpers";
import { ApplicationType, ButtonTextSize } from "src/types";
import { ref } from "vue";

const emits = defineEmits<{
  (event: "onContinue", id: ApplicationType): void;
}>();

const selectedApplication = ref<ApplicationType | undefined>(undefined);

const applicationTileData = [
  {
    id: ApplicationType.AIRBRAKE,
    name: "Airbrake",
    image: getSrc("/src/assets/images/airbrake-logo.png"),
  },
  {
    id: ApplicationType.NEWRELIC,
    name: "NewRelic",
    image: getSrc("/src/assets/images/newrelic-logo.svg"),
  },
  {
    id: ApplicationType.SENTRY,
    name: "Sentry",
    image: getSrc("/src/assets/images/sentry-logo.svg"),
  },
  {
    id: ApplicationType.OTHER,
    name: "Other",
    image: getSrc("/src/assets/images/other-app-icon.svg"),
  },
];

const onSubmit = () => {
  if (!selectedApplication.value) return;
  emits("onContinue", selectedApplication.value);
};
</script>
