<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { TrashIcon } from "@heroicons/vue/24/outline";
import { useApplicationsStore } from "../state/applications";
import TheMainLayout from "../layouts/the-main-layout.vue";
import { IApplication } from "../types";
import { fetchApplicationById } from "../api/applications";

const { getCachedApplication, cacheApplication } = useApplicationsStore();

const activeApplication = ref<IApplication | undefined>(undefined);

const router = useRouter();
const route = useRoute();

onMounted(async () => {
  const applicationId = parseInt(route.params.applicationId as string);
  console.log(applicationId);
  const cachedApplication = getCachedApplication(applicationId);
  if (cachedApplication != null) {
    activeApplication.value = cachedApplication;
    return;
  }
  try {
    const application = await fetchApplicationById({ applicationId });
    if (!application) {
      router.push("/404");
      return;
    }
    cacheApplication(application);
    activeApplication.value = application;
  } catch (error) {
    console.error(
      "Galata Error: An error occurred trying to fetch the application you wanted:",
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
        <h2 class="capitalize">{{ activeApplication?.Name }}</h2>
        <button
          role="button"
          class="w-8 h-8 p-0 m-0 bg-white text-red-500 hover:shadow-md rounded-full border-none transition-all duration-500"
          aria-label="Delete Application"
        >
          <trash-icon class="w-5 h-5" />
        </button>
      </div>
      <div class="w-full h-full flex">
        <div class="w-2/3 h-full">
          <h5 class="p-0 m-0 mb-1">Alerts</h5>
          <div
            class="w-full h-full mr-4 border-1 border-gray-300 rounded-lg overflow-y-auto"
          ></div>
        </div>
      </div>
    </div>
  </the-main-layout>
</template>
